package rpc_speed

import "github.com/gagliardetto/solana-go/rpc"

type Endpoint struct {
	Name   string      `json:"name"`
	Client *rpc.Client `json:"-"`
}
