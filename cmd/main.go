package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/benchmark"
	"github.com/erodriguezg/chapter-golang/pkg/benchmark2"
	"github.com/erodriguezg/chapter-golang/pkg/config"
	"github.com/erodriguezg/chapter-golang/pkg/problems"
)

func main() {

	// main param
	example := flag.String("example", "", "example to execute")

	// sql params
	delete := flag.Bool("delete", false, "delete the test data")
	fail := flag.Bool("fail", false, "force the fail of the transaction")

	// pointer benchmark params
	maxDeep := flag.Int("max-deep", 1, "the max deep of recursion")
	arraySize := flag.Int("array-size", 1, "the size of process of the array in the recursion")

	// pointer benchmark2 param
	iterations := flag.Int("iterations", 1, "amount of iterations for the benchmark 2")

	flag.Parse()

	switch *example {

	case "problem-float32":
		problems.Float32ExampleProblem()

	case "problem-config":
		problems.Config()

	case "tx":
		mainSqlTemplate(true, *delete, *fail)

	case "no-tx":
		mainSqlTemplate(false, *delete, *fail)

	case "pointer-bench-by-val":
		benchmark.BenchmarkByValue(*maxDeep, *arraySize)

	case "pointer-bench-by-ref":
		benchmark.BenchmarkByRef(*maxDeep, *arraySize)

	case "pointer-bench-2-by-val":
		benchmark2.BenchmarkByVal(*iterations)

	case "pointer-bench-2-by-ref":
		benchmark2.BenchmarkByRef(*iterations)
	}

}

func mainSqlTemplate(isTx bool, delete bool, fail bool) {

	defer config.CloseDemoSqlAll()

	config.ConfigDemoSqlAll()

	fmt.Printf("====> Params: isTx: %t, Delete: %t, Fail: %t", isTx, delete, fail)

	service := config.GetDemoTxService()

	var err error
	if isTx {
		err = service.ProcessWithTx(delete, fail)
	} else {
		err = service.ProcessWithoutTx(delete, fail)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "An error has ocurred: \n%v\n", err)
	}

}
