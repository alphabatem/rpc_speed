package util

import (
	"encoding/json"
	"github.com/alphabatem/rpc_speed"
	"github.com/gagliardetto/solana-go/rpc"
	"os"
)

type fileRpc struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func LoadEndpoints() ([]*rpc_speed.Endpoint, error) {
	d, err := os.ReadFile("./endpoints.json")
	if err != nil {
		return nil, err
	}

	var out []*fileRpc
	err = json.Unmarshal(d, &out)
	if err != nil {
		return nil, err
	}

	endpoints := make([]*rpc_speed.Endpoint, len(out))
	for i, o := range out {
		endpoints[i] = &rpc_speed.Endpoint{
			Name:   o.Name,
			Client: rpc.New(o.Url),
		}
	}
	return endpoints, nil
}
