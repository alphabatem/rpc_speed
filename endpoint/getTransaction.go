package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type GetTransaction struct {
	endpoint
	Signature solana.Signature          `json:"signature"`
	Options   *rpc.GetTransactionOpts   `json:"options"`
	Results   *rpc.GetTransactionResult `json:"results"`
}

func NewGetTransaction(signature solana.Signature, opts *rpc.GetTransactionOpts) *GetTransaction {
	return &GetTransaction{
		Signature: signature,
		Options:   opts,
	}
}

func (call *GetTransaction) Name() string {
	return call.EndpointName(fmt.Sprintf("getTransaction:%s", call.Signature))
}

func (call *GetTransaction) Run(ctx context.Context, c *rpc.Client) ([]byte, error) {
	call.Start()
	defer call.Stop()

	resp, err := c.GetTransaction(ctx, call.Signature, call.Options)
	if err != nil {
		return nil, err
	}

	call.Results = resp
	return json.Marshal(resp)
}
