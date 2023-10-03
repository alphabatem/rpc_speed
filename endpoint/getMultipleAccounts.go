package endpoint

import (
	"context"
	"encoding/json"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type GetMultipleAccounts struct {
	endpoint
	Accounts []solana.PublicKey             `json:"accounts"`
	Options  *rpc.GetMultipleAccountsOpts   `json:"options"`
	Results  *rpc.GetMultipleAccountsResult `json:"results"`
}

func NewGetMultipleAccounts(accounts []solana.PublicKey, opts *rpc.GetMultipleAccountsOpts) *GetMultipleAccounts {
	return &GetMultipleAccounts{
		Accounts: accounts,
		Options:  opts,
	}
}

func (call *GetMultipleAccounts) Name() string {
	return call.EndpointName("getMultipleAccounts")
}

func (call *GetMultipleAccounts) Run(ctx context.Context, c *rpc.Client) ([]byte, error) {
	call.Start()
	defer call.Stop()

	resp, err := c.GetMultipleAccountsWithOpts(ctx, call.Accounts, call.Options)
	if err != nil {
		return nil, err
	}

	call.Results = resp
	return json.Marshal(resp)
}
