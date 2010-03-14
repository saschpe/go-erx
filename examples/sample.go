package main

import (
	"fmt"
	"erx"
	"os"
)

const (
	ERROR_ONE = iota
	ERROR_TWO
	ERROR_THREE
	ERROR_FOUR
)

const scope = "main"

var errors = map[uint]string {
	ERROR_ONE   : "error one",
	ERROR_TWO   : "error two",
	ERROR_THREE : "error three",
	ERROR_FOUR  : "error four",
}

type MyType struct {
    v int
}

func (m *MyType) String() string {
    return "Hello error!"
}

func init() {
	erx.RegisterScopeMessages(scope, errors)
	erx.AddPathCut("/home/steplg/quickdoc/workspaces/jtt/")
}

func main() {
	var m MyType
	_, osError := os.Open("nonExistedFile.tmp", os.O_RDONLY, 0000)
	err := erx.NewSequent(scope, ERROR_ONE, osError)
	err.AddV("var1", "444")
	err.AddV("var2", &m)
	err1 := erx.NewSequent(scope, ERROR_TWO, err)
	formatter := erx.NewStringFormatter("  ")
	fmt.Println(formatter.Format(err))
	fmt.Println(formatter.Format(err1))
}
