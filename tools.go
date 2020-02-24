package main

import (
	"github.com/common-nighthawk/go-figure"
)

func printName(name string) {
	myFigure := figure.NewFigure(name, "", true)
	myFigure.Print()
}
