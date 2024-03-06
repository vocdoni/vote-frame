package main

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/vocdoni/vote-frame/imageframe"
	"go.vocdoni.io/dvote/api"
	"go.vocdoni.io/dvote/httprouter"
	"go.vocdoni.io/dvote/httprouter/apirest"
	"go.vocdoni.io/dvote/log"
	"go.vocdoni.io/dvote/types"
)

// checkIfElectionFinishedAndHandle checks if the election is finished and if so, sends the final results.
// Returns true if the election is finished and the response was sent, false otherwise.
// The caller should return immediately after this function returns true.
func (v *vocdoniHandler) checkIfElectionFinishedAndHandle(electionID types.HexBytes, ctx *httprouter.HTTPContext) bool {
	pngResults := v.db.FinalResultsPNG(electionID)
	if pngResults == nil {
		return false
	}
	response := strings.ReplaceAll(frame(frameFinalResults), "{image}", v.addImageToCache(pngResults, electionID))
	response = strings.ReplaceAll(response, "{processID}", electionID.String())
	response = strings.ReplaceAll(response, "{title}", "Final results")

	ctx.SetResponseContentType("text/html; charset=utf-8")
	if err := ctx.Send([]byte(response), http.StatusOK); err != nil {
		log.Warnw("failed to send response", "error", err)
		return true
	}
	return true
}

func (v *vocdoniHandler) results(msg *apirest.APIdata, ctx *httprouter.HTTPContext) error {
	electionID := ctx.URLParam("electionID")
	if len(electionID) == 0 {
		return errorImageResponse(ctx, fmt.Errorf("invalid electionID"))
	}
	log.Infow("received results request", "electionID", electionID)
	electionIDbytes, err := hex.DecodeString(electionID)
	if err != nil {
		return errorImageResponse(ctx, fmt.Errorf("failed to decode electionID: %w", err))
	}
	// check if the election is finished and if so, send the final results as a static PNG
	if v.checkIfElectionFinishedAndHandle(electionIDbytes, ctx) {
		return nil
	}

	// get the election from the vochain and create a PNG image with the results
	election, err := v.cli.Election(electionIDbytes)
	if err != nil {
		return errorImageResponse(ctx, fmt.Errorf("failed to fetch election: %w", err))
	}
	if election.Results == nil || len(election.Results) == 0 {
		return errorImageResponse(ctx, fmt.Errorf("election results not ready"))
	}
	// Update LRU cached election
	evicted := v.electionLRU.Add(electionID, election)
	log.Debugw("updated election cache", "electionID", electionID, "evicted", evicted)

	// if final results, create the static PNG image with the results
	if election.FinalResults {
		png, err := buildResultsPNG(election)
		if err != nil {
			return fmt.Errorf("failed to create image: %w", err)
		}
		if err := v.db.AddFinalResults(electionIDbytes, png); err != nil {
			return fmt.Errorf("failed to add final results to database: %w", err)
		}
		log.Infow("final results image built ondemand", "electionID", electionID)
		if v.checkIfElectionFinishedAndHandle(electionIDbytes, ctx) {
			return nil
		}
	}
	// if not final results, create the dynamic PNG image with the results
	png, err := buildResultsPNG(election)
	if err != nil {
		return fmt.Errorf("failed to create image: %w", err)
	}
	response := strings.ReplaceAll(frame(frameResults), "{image}", v.addImageToCache(png, electionIDbytes))
	response = strings.ReplaceAll(response, "{title}", election.Metadata.Title["default"])
	response = strings.ReplaceAll(response, "{processID}", electionID)
	ctx.SetResponseContentType("text/html; charset=utf-8")
	return ctx.Send([]byte(response), http.StatusOK)
}

func buildResultsPNG(election *api.Election) ([]byte, error) {
	return imageframe.ResultsImage(election)
}
