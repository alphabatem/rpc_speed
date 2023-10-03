package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/alphabatem/rpc_speed"
	"github.com/alphabatem/rpc_speed/endpoint"
	"github.com/alphabatem/rpc_speed/suite"
	"github.com/alphabatem/rpc_speed/util"
	"log"
	"os"
	"time"
)

func main() {
	st := speedTest{}
	if err := st.Configure(); err != nil {
		panic(err)
	}

	if err := st.Start(); err != nil {
		panic(err)
	}
}

type testResults []speedResult

type speedResult struct {
	Elapsed time.Duration `json:"elapsed"`
	Result  []byte        `json:"result"`
	Error   error         `json:"error"`
}

type speedTest struct {
	endpoints []*rpc_speed.Endpoint
	suite     suite.Suite

	//Amount of cycles to run per test
	cycles int

	cycleResults [][]testResults
	resultMatrix map[string]testResults
}

func (t *speedTest) Configure() (err error) {
	t.resultMatrix = map[string]testResults{}

	t.cycles = 3
	t.cycleResults = make([][]testResults, t.cycles)

	//Load our test suite for now
	t.suite = suite.DefaultSuite()

	t.endpoints, err = util.LoadEndpoints()
	if err != nil {
		return err
	}

	log.Printf("%v Tests loaded", len(t.suite))
	log.Printf("%v Endpoints loaded", len(t.endpoints))

	return nil
}

func (t *speedTest) Start() (err error) {
	for i := 0; i < t.cycles; i++ {
		log.Printf("Cycle: %v", i)
		t.cycleResults[i] = t.executeCycle()
	}

	return t.printResults()
}

func (t *speedTest) executeCycle() []testResults {
	testResult := make([]testResults, len(t.suite))
	for i, c := range t.suite {
		tr := t.executeTest(c)
		testResult[i] = tr
	}
	return testResult
}

func (t *speedTest) executeTest(tc endpoint.Callable) []speedResult {
	results := make([]speedResult, len(t.endpoints))

	for i, c := range t.endpoints {
		cx := context.TODO()
		log.Printf("%v Running %s - %s - %s", i, c.Name, tc.Name(), tc.Elapsed())

		d, err := tc.Run(cx, c.Client)

		results[i] = speedResult{
			Elapsed: tc.Elapsed(),
			Result:  d,
			Error:   err,
		}
	}
	return results
}

func (t *speedTest) avgResults() map[string][]int64 {
	//map[testName]map[endpoint]testDuration
	endpointTiming := map[string][]int64{}

	for ti, tst := range t.suite {
		avgTimes := make([]int64, len(t.endpoints))

		for ei := range t.endpoints {
			totDur := time.Duration(0)

			//Get times across cycles
			for x := 0; x < t.cycles; x++ {
				totDur += t.cycleResults[x][ti][ei].Elapsed
			}
			avgTimes[ei] = totDur.Milliseconds() / int64(t.cycles)
		}

		endpointTiming[tst.Name()] = avgTimes
	}

	return endpointTiming
}

func (t *speedTest) printResults() error {
	f, err := os.Create("speed_results.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	cf := csv.NewWriter(f)
	defer cf.Flush()

	headers := make([]string, len(t.endpoints)+1)
	headers[0] = "RPC / Method"
	for i, e := range t.endpoints {
		headers[i+1] = e.Name
	}
	err = cf.Write(headers)
	if err != nil {
		return err
	}

	avg := t.avgResults()
	for testName, av := range avg {
		row := make([]string, len(av)+1)
		row[0] = testName
		for i, x := range av {
			row[i+1] = fmt.Sprintf("%v", x)
		}
		err = cf.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
