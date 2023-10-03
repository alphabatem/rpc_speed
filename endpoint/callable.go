package endpoint

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"
	"time"
)

type runner func(ctx context.Context, c *rpc.Client) ([]byte, error)

type Callable interface {
	Name() string
	Run(ctx context.Context, c *rpc.Client) ([]byte, error)
	Elapsed() time.Duration
}
