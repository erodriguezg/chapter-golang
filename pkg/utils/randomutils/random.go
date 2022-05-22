package randomutils

import (
	"math/rand"
)

const (
	loremimpsum = `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's 
	standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make 
	a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, 
	remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing 
	Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions 
	of Lorem Ipsum.`
)

func GenInt64(start int64, end int64) int64 {
	rate := rand.Float64()
	return start + (int64)((float64(end)-float64(start))*rate)
}

func GenText() string {
	end := GenInt64(0, (int64)(len(loremimpsum)-1))
	return loremimpsum[:end]
}

func GenBool() bool {
	aux := GenInt64(0, 100)
	if aux > 50 {
		return true
	} else {
		return false
	}
}

func GenNumber() int64 {
	return GenInt64(0, 1000000)
}

func GenDecimal() float64 {
	return rand.Float64()
}
