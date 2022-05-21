package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/erodriguezg/chapter-golang/pkg/config"
	"github.com/erodriguezg/chapter-golang/pkg/problems"
)

func main() {

	example := flag.String("example", "", "example to execute")
	delete := flag.Bool("delete", false, "delete the test data")
	fail := flag.Bool("fail", false, "force the fail of the transaction")

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
