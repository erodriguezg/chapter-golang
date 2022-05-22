package benchmark

import (
	"fmt"
	"time"

	"github.com/erodriguezg/chapter-golang/pkg/utils/randomutils"
)

type something struct {
	ID            string
	Index         int64
	GUID          string
	IsActive      bool
	Balance       string
	Picture       string
	Age           int64
	EyeColor      string
	Name          string
	Gender        string
	Company       string
	Email         string
	Phone         string
	Address       string
	About         string
	Registered    string
	Latitude      float64
	Longitude     float64
	Greeting      string
	FavoriteFruit string
}

type mapValByValue struct {
	firstSomething something
	arraySomething []something
}

type mapValByRef struct {
	firstSomething *something
	arraySomething []*something
}

func BenchmarkByValue(maxDeep int, arraySize int) {
	auxMap := map[int64]mapValByValue{}
	deep := 0
	thing := generateSomethingByValue()

	start := time.Now()

	iterateByValue(auxMap, thing, deep, maxDeep, arraySize)

	elapsed := time.Since(start)

	fmt.Printf("\nBenchmarkByValue. MaxDeep: %d, ArraySize: %d took %d ms\n", maxDeep, arraySize, elapsed.Milliseconds())
}

func iterateByValue(auxMap map[int64]mapValByValue, thing something, deep int, maxDeep int, arraySize int) {
	var thingArray []something
	for i := 0; i < arraySize; i++ {
		thing := generateSomethingByValue()
		thingArray = append(thingArray, thing)
	}
	auxMap[int64(deep)] = mapValByValue{
		firstSomething: thing,
		arraySomething: thingArray,
	}
	if deep < maxDeep {
		iterateByValue(auxMap, thingArray[0], deep+1, maxDeep, arraySize)
	} else {
		return
	}
}

func BenchmarkByRef(maxDeep int, arraySize int) {
	auxMap := map[int64]mapValByRef{}
	deep := 0
	thing := generateSomethingByRef()

	start := time.Now()

	iterateByRef(auxMap, thing, deep, maxDeep, arraySize)

	elapsed := time.Since(start)

	fmt.Printf("\nBenchmarkByRef. MaxDeep: %d, ArraySize: %d took %d ms\n", maxDeep, arraySize, elapsed.Milliseconds())

}

func iterateByRef(auxMap map[int64]mapValByRef, thing *something, deep int, maxDeep int, arraySize int) {
	var thingArray []*something
	var thingAux *something
	for i := 0; i < arraySize; i++ {
		thingAux = generateSomethingByRef()
		thingArray = append(thingArray, thingAux)
	}
	auxMap[int64(deep)] = mapValByRef{
		firstSomething: thing,
		arraySomething: thingArray,
	}
	if deep < maxDeep {
		iterateByRef(auxMap, thingArray[0], deep+1, maxDeep, arraySize)
	} else {
		return
	}
}

func generateSomethingByValue() something {
	return something{
		ID:            randomutils.GenText(),
		Index:         randomutils.GenNumber(),
		GUID:          randomutils.GenText(),
		IsActive:      randomutils.GenBool(),
		Balance:       randomutils.GenText(),
		Picture:       randomutils.GenText(),
		Age:           randomutils.GenNumber(),
		EyeColor:      randomutils.GenText(),
		Name:          randomutils.GenText(),
		Gender:        randomutils.GenText(),
		Company:       randomutils.GenText(),
		Email:         randomutils.GenText(),
		Phone:         randomutils.GenText(),
		Address:       randomutils.GenText(),
		About:         randomutils.GenText(),
		Registered:    randomutils.GenText(),
		Latitude:      randomutils.GenDecimal(),
		Longitude:     randomutils.GenDecimal(),
		Greeting:      randomutils.GenText(),
		FavoriteFruit: randomutils.GenText(),
	}
}

func generateSomethingByRef() *something {
	thing := something{
		ID:            randomutils.GenText(),
		Index:         randomutils.GenNumber(),
		GUID:          randomutils.GenText(),
		IsActive:      randomutils.GenBool(),
		Balance:       randomutils.GenText(),
		Picture:       randomutils.GenText(),
		Age:           randomutils.GenNumber(),
		EyeColor:      randomutils.GenText(),
		Name:          randomutils.GenText(),
		Gender:        randomutils.GenText(),
		Company:       randomutils.GenText(),
		Email:         randomutils.GenText(),
		Phone:         randomutils.GenText(),
		Address:       randomutils.GenText(),
		About:         randomutils.GenText(),
		Registered:    randomutils.GenText(),
		Latitude:      randomutils.GenDecimal(),
		Longitude:     randomutils.GenDecimal(),
		Greeting:      randomutils.GenText(),
		FavoriteFruit: randomutils.GenText(),
	}
	return &thing
}
