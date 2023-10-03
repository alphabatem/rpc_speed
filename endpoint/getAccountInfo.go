package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type GetAccountInfo struct {
	endpoint
	Account solana.PublicKey          `json:"account"`
	Options *rpc.GetAccountInfoOpts   `json:"options"`
	Results *rpc.GetAccountInfoResult `json:"results"`
}

func NewGetAccountInfo(account solana.PublicKey, opts *rpc.GetAccountInfoOpts) *GetAccountInfo {
	return &GetAccountInfo{
		Account: account,
		Options: opts,
	}
}

func (call *GetAccountInfo) Name() string {
	return call.EndpointName(fmt.Sprintf("getAccountInfo:%s", call.Account))
}

func (call *GetAccountInfo) Run(ctx context.Context, c *rpc.Client) ([]byte, error) {
	call.Start()
	defer call.Stop()

	resp, err := c.GetAccountInfoWithOpts(ctx, call.Account, call.Options)
	if err != nil {
		return nil, err
	}

	call.Results = resp
	return json.Marshal(resp)
}
