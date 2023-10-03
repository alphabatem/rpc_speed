package main

import (
	ctx "context"
	"fmt"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
	"os"
	"time"
)

/**
 * RpcTest is the v1 RPC Speed Test
 * - Doesnt account for latency
 * - Simple ping test more than anything
 */
type RpcTest struct {
	Name   string
	Client *rpc.Client

	Count        int64
	lastRuntime  time.Duration
	totalRuntime time.Duration
}

func (rt *RpcTest) AverageRuntime() time.Duration {
	avg := rt.totalRuntime.Nanoseconds() / rt.Count
	return time.Duration(avg)
}

func (rt *RpcTest) Test() error {
	tn := time.Now()
	_, err := rt.Client.GetLatestBlockhash(ctx.Background(), "confirmed")
	if err != nil {
		return err
	}
	tne := time.Now().Sub(tn)

	rt.lastRuntime = tne
	rt.totalRuntime += tne

	log.Printf("%s: %s", rt.Name, tne)
	return nil
}

func main() {
	rpc1 := &RpcTest{
		Name:   os.Getenv("RPC_1_NAME"),
		Client: rpc.New(os.Getenv("RPC_1_URL")),
	}
	rpc2 := &RpcTest{
		Name:   os.Getenv("RPC_2_NAME"),
		Client: rpc.New(os.Getenv("RPC_2_URL")),
	}
	rpc3 := &RpcTest{
		Name:   os.Getenv("RPC_3_NAME"),
		Client: rpc.New(os.Getenv("RPC_3_URL")),
	}
	rpc4 := &RpcTest{
		Name:   os.Getenv("RPC_4_NAME"),
		Client: rpc.New(os.Getenv("RPC_4_URL")),
	}
	rpc5 := &RpcTest{
		Name:   os.Getenv("RPC_5_NAME"),
		Client: rpc.New(os.Getenv("RPC_5_URL")),
	}
	rpc6 := &RpcTest{
		Name:   os.Getenv("RPC_6_NAME"),
		Client: rpc.New(os.Getenv("RPC_6_URL")),
	}
	rpc7 := &RpcTest{
		Name:   os.Getenv("RPC_7_NAME"),
		Client: rpc.New(os.Getenv("RPC_7_URL")),
	}

	testers := []*RpcTest{rpc1, rpc2, rpc3, rpc4, rpc5, rpc6, rpc7}
	loops := 100

	for i := 0; i < loops; i++ {
		//Test all testers
		for _, t := range testers {
			err := t.Test()
			if err != nil {
				panic(fmt.Sprintf("%s Failed: %s", t.Name, err))
			}
			t.Count++
		}
	}

	//Results
	for _, t := range testers {
		log.Printf("%s (%v) - Avg: %s", t.Name, t.Count, t.AverageRuntime())
	}
}
