package suite

import (
	"github.com/alphabatem/rpc_speed/endpoint"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func DefaultSuite() []endpoint.Callable {
	HUNDRED := uint64(100)
	ONE := uint64(1)

	historicalSig := solana.MustSignatureFromBase58("4QKAYZmXF8aNdECfHMkjw1nEsJSkPcPki4o1CBpmw4ZTHjc2BH1ghmPYuRENry4M833rCwz2VDFUu61hhPBVZiSL")

	return []endpoint.Callable{
		endpoint.NewGetLatestBlockHash(),
		endpoint.NewGetMultipleAccounts([]solana.PublicKey{solana.TokenProgramID}, nil),
		endpoint.NewGetProgramAccounts(solana.StakeProgramID, &rpc.GetProgramAccountsOpts{DataSlice: &rpc.DataSlice{Length: &HUNDRED}}),
		endpoint.NewGetTransaction(historicalSig, &rpc.GetTransactionOpts{MaxSupportedTransactionVersion: &ONE}),
	}
}
