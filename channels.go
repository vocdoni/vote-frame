package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vocdoni/vote-frame/farcasterapi"
	"go.vocdoni.io/dvote/httprouter"
	"go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
)

func (v *vocdoniHandler) channelHandler(msg *apirest.APIdata, ctx *httprouter.HTTPContext) error {
	channelID := ctx.URLParam("channelID")
	if channelID == "" {
		return ctx.Send([]byte("no channel id provided"), http.StatusBadRequest)
	}
	ch, err := v.fcapi.Channel(ctx.Request.Context(), channelID)
	if err != nil {
		if err == farcasterapi.ErrChannelNotFound {
			return ctx.Send([]byte("channel not found"), http.StatusNotFound)
		}
		return ctx.Send([]byte(err.Error()), apirest.HTTPstatusInternalErr)
	}
	res, err := json.Marshal(Channel{
		ID:          ch.ID,
		Name:        ch.Name,
		Description: ch.Description,
		Followers:   ch.Followers,
		ImageURL:    ch.Image,
		URL:         ch.URL,
	})
	if err != nil {
		return ctx.Send([]byte("error encoding channel details"), apirest.HTTPstatusInternalErr)
	}
	return ctx.Send(res, http.StatusOK)
}

func (v *vocdoniHandler) findChannelHandler(msg *apirest.APIdata, ctx *httprouter.HTTPContext) error {
	adminFid := uint64(0)
	if _, ok := ctx.Request.URL.Query()["imAdmin"]; ok {
		token := msg.AuthToken
		if token == "" {
			return fmt.Errorf("missing auth token header")
		}
		auth, err := v.db.UpdateActivityAndGetData(token)
		if err != nil {
			return ctx.Send([]byte(err.Error()), apirest.HTTPstatusNotFound)
		}
		adminFid = auth.UserID
	}
	query := ctx.Request.URL.Query().Get("q")
	if query == "" {
		return ctx.Send([]byte("query parameter not provided"), http.StatusBadRequest)
	}
	channels, err := v.fcapi.FindChannel(ctx.Request.Context(), query, adminFid)
	if err != nil {
		log.Errorw(err, "failed to list channels")
		return ctx.Send([]byte("error getting list of channels"), http.StatusInternalServerError)
	}
	res := ChannelList{
		Channels: []*Channel{},
	}
	for _, ch := range channels {
		res.Channels = append(res.Channels, &Channel{
			ID:          ch.ID,
			Name:        ch.Name,
			Description: ch.Description,
			Followers:   ch.Followers,
			ImageURL:    ch.Image,
			URL:         ch.URL,
		})
	}
	bRes, err := json.Marshal(res)
	if err != nil {
		log.Errorw(err, "failed to marshal channels")
		return ctx.Send([]byte("error marshaling channels"), http.StatusInternalServerError)
	}
	return ctx.Send(bRes, http.StatusOK)
}
