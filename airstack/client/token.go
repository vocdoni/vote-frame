package client

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	gql "github.com/vocdoni/vote-frame/airstack/graphql"
	"go.vocdoni.io/dvote/log"
)

// TokenDetails wraps useful information about a token
type TokenDetails struct {
	Decimals    int
	Name        string
	Symbol      string
	TotalSupply *big.Int
}

// TokenDetails gets a token details given its address and blockchain
func (c *Client) TokenDetails(
	tokenAddress common.Address,
	blockchain string,
) (*TokenDetails, error) {
	cctx, cancel := context.WithTimeout(c.ctx, apiTimeout)
	defer cancel()
	r := 0
	resp := &gql.GetTokenDetailsResponse{}
	var err error
	chain, ok := c.blockchainToTokenBlockchain(blockchain)
	if !ok {
		return nil, fmt.Errorf("invalid blockchain provided")
	}
	for r < maxAPIRetries {
		resp, err = gql.GetTokenDetails(cctx, c.Client, tokenAddress, chain)
		if err != nil {
			r += 1
			time.Sleep(time.Second * 3)
			continue
		}
		if len(resp.Tokens.GetToken()) != 1 {
			return nil, fmt.Errorf("invalid token details response from Airstack")
		}
		td := resp.GetTokens()
		td.GetToken()
		totalSupply := new(big.Int)
		totalSupply.SetString(resp.Tokens.GetToken()[0].TotalSupply, 10)
		return &TokenDetails{
			Decimals:    resp.Tokens.GetToken()[0].Decimals,
			Name:        resp.Tokens.GetToken()[0].Name,
			Symbol:      resp.Tokens.GetToken()[0].Symbol,
			TotalSupply: totalSupply,
		}, nil
	}
	return nil, fmt.Errorf("max GraphQL retries reached, error: %w", err)
}

// TokenHolder wraps a token holder with its address and balance of a certain token
type TokenHolder struct {
	Address common.Address
	Balance *big.Int
}

// tokenBalances is a wrapper around the generated function for GraphQL query TokenBalances.
func (c *Client) tokenBalances(
	tokenAddress common.Address,
	blockchain gql.TokenBlockchain,
	limit int,
	cursor string,
) (*gql.GetTokenBalancesResponse, error) {
	cctx, cancel := context.WithTimeout(c.ctx, apiTimeout)
	defer cancel()
	r := 0
	var err error
	resp := &gql.GetTokenBalancesResponse{}
	for r < maxAPIRetries {
		resp, err = gql.GetTokenBalances(cctx, c.Client, tokenAddress, blockchain, limit, cursor)
		if err != nil {
			r += 1
			time.Sleep(time.Second * 3)
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("max GraphQL retries reached, error: %w", err)
}

// TokenBalances gets all the token holders of a given token in a given chain
// calling the Airstack API. This function also takes care of Airstack API pagination.
func (c *Client) TokenBalances(tokenAddress common.Address, blockchain string) ([]*TokenHolder, error) {
	hasNextPage := true
	cursor := ""
	th := make([]*TokenHolder, 0)
	chain, ok := c.blockchainToTokenBlockchain(blockchain)
	if !ok {
		return nil, fmt.Errorf("invalid blockchain provided")
	}
	totalHolders := 0
	totalPages := 0
	for hasNextPage {
		resp, err := c.tokenBalances(tokenAddress, chain, airstackAPIlimit, cursor)
		if err != nil {
			return nil, fmt.Errorf("cannot get token balances from Airstack: %w", err)
		}
		for _, holder := range resp.TokenBalances.TokenBalance {
			balance := new(big.Int)
			balance.SetString(holder.Amount, 10)
			th = append(th, &TokenHolder{
				Address: holder.Owner.Addresses[0],
				Balance: balance,
			})
		}
		cursor = resp.TokenBalances.PageInfo.NextCursor
		hasNextPage = cursor != ""
		totalHolders += len(resp.TokenBalances.TokenBalance)
		totalPages++
	}
	log.Debugf("fetched %d items in %d pages for token %s", totalHolders, totalPages, tokenAddress)
	return th, nil
}

// CheckIfHolder checks if a given address is a holder of a given token in a given chain
// Returns the balance of the token if the address is a holder.
func (c *Client) CheckIfHolder(tokenAddress common.Address, blockchain string, address common.Address) (uint64, error) {
	chain, ok := c.blockchainToTokenBlockchain(blockchain)
	if !ok {
		return 0, fmt.Errorf("invalid blockchain provided")
	}
	resp, err := c.holderOf(tokenAddress, chain, address)
	if err != nil {
		return 0, fmt.Errorf("cannot check token ownership from Airstack: %w", err)
	}
	if len(resp.TokenBalances.TokenBalance) == 0 {
		return 0, nil
	}
	return uint64(resp.TokenBalances.TokenBalance[0].FormattedAmount), nil
}

// holderOf is a wrapper around the generated function for GraphQL query checkTokenOwnership.
func (c *Client) holderOf(
	tokenAddress common.Address,
	blockchain gql.TokenBlockchain,
	address common.Address,
) (*gql.CheckTokenOwnershipResponse, error) {
	cctx, cancel := context.WithTimeout(c.ctx, apiTimeout)
	defer cancel()
	r := 0
	var err error
	resp := &gql.CheckTokenOwnershipResponse{}
	for r < maxAPIRetries {
		resp, err = gql.CheckTokenOwnership(cctx, c.Client, address.Hex(), tokenAddress, blockchain)
		if err != nil {
			r += 1
			time.Sleep(time.Second * 3)
			continue
		}
		return resp, nil
	}
	return nil, fmt.Errorf("max GraphQL retries reached, error: %w", err)
}
