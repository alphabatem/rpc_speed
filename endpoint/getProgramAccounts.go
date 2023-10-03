package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type GetProgramAccounts struct {
	endpoint
	Account solana.PublicKey              `json:"account"`
	Options *rpc.GetProgramAccountsOpts   `json:"options"`
	Results *rpc.GetProgramAccountsResult `json:"results"`
}

func NewGetProgramAccounts(account solana.PublicKey, opts *rpc.GetProgramAccountsOpts) *GetProgramAccounts {
	return &GetProgramAccounts{
		Account: account,
		Options: opts,
	}
}

func (call *GetProgramAccounts) Name() string {
	return call.EndpointName(fmt.Sprintf("getProgramAccounts:%s", call.Account))
}

func (call *GetProgramAccounts) Run(ctx context.Context, c *rpc.Client) ([]byte, error) {
	call.Start()
	defer call.Stop()

	resp, err := c.GetProgramAccountsWithOpts(ctx, call.Account, call.Options)
	if err != nil {
		return nil, err
	}

	call.Results = &resp
	return json.Marshal(resp)
}
