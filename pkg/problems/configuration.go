package problems

import (
	"github.com/erodriguezg/chapter-golang/pkg/config"
)

func Config() {
	config.ConfigAll()

	barService1 := config.GetBarService1()
	barService2 := config.GetBarService2()

	barService1.DoBar()
	barService2.DoBar()
}
