package endpoint

import (
	"context"
	"encoding/json"
	"github.com/gagliardetto/solana-go/rpc"
)

type GetLatestBlockHash struct {
	endpoint
	Result     *rpc.LatestBlockhashResult `json:"result"`
	Commitment rpc.CommitmentType         `json:"commitment"`
}

func NewGetLatestBlockHash() *GetLatestBlockHash {
	return &GetLatestBlockHash{
		Commitment: rpc.CommitmentConfirmed,
	}
}

func (call *GetLatestBlockHash) Name() string {
	return call.EndpointName("getLatestBlockHash")
}

func (call *GetLatestBlockHash) Run(ctx context.Context, c *rpc.Client) ([]byte, error) {
	call.Start()
	defer call.Stop()

	bh, err := c.GetLatestBlockhash(ctx, call.Commitment)
	if err != nil {
		return nil, err
	}

	call.Result = bh.Value
	return json.Marshal(bh.Value)
}
