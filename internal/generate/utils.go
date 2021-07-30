package generate

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/go-courier/packagesx"
)

func logCost() func() {
	startedAt := time.Now()

	return func() {
		log.Printf("costs %s", color.GreenString("%0.0f ms", float64(time.Now().Sub(startedAt)/time.Millisecond)))
	}
}

type Generator interface {
	Output(cwd string)
}

func RunGenerator(createGenerator func(pkg *packagesx.Package) Generator) {
	defer logCost()()

	cwd, _ := os.Getwd()

	pkg, err := packagesx.Load(cwd)
	if err != nil {
		panic(err)
	}

	g := createGenerator(pkg)
	g.Output(cwd)
}
