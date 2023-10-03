package suite

import (
	"github.com/alphabatem/rpc_speed/endpoint"
)

func SimpleSuite() []endpoint.Callable {
	return []endpoint.Callable{
		endpoint.NewGetLatestBlockHash(),
	}
}
