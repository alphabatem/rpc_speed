package endpoint

import (
	"github.com/gagliardetto/solana-go/rpc"
	"time"
)

type endpoint struct {
	customName string

	startTime time.Time
	endTime   time.Time

	client *rpc.Client
}

func (e *endpoint) SetCustomName(name string) {
	e.customName = name
}

func (e *endpoint) EndpointName(name string) string {
	if e.customName != "" {
		return e.customName
	}
	return name
}

func (e *endpoint) Start() {
	e.startTime = time.Now()
}

func (e *endpoint) Stop() {
	e.endTime = time.Now()
}

func (e *endpoint) Elapsed() time.Duration {
	return e.endTime.Sub(e.startTime)
}
