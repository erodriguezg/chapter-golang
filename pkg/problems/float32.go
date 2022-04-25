package problems

import "fmt"

func Float32ExampleProblem() {

	var float32Data float32
	float32Data = 123456789

	fmt.Printf("data float32: %f \n", float32Data)

	var float64Data float64
	float64Data = 123456789

	fmt.Printf("\ndata float64: %f \n", float64Data)

	var float64DataCast float64
	float64DataCast = (float64(float32Data))
	fmt.Printf("\ndata float64 Cast From 32: %f \n", float64DataCast)

	var float32DataCast float32
	float32DataCast = float32(float64Data)
	fmt.Printf("\ndata float32 Cast From 64: %f \n", float32DataCast)
}
