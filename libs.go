package main

import (
	"fmt"
	"os"
	"errors"
)

type (
	Writer struct{}
)
var (
	wr = Writer{}
)

func (w Writer) l(a ...interface{}) {
	fmt.Println(a...)
}

func (w Writer) f(str string, a ...interface{}) {
	fmt.Printf(str, a...)
}

func (w Writer) sf(str string, a ...interface{}) {
	fmt.Sprintf(str, a...)
}

func (w Writer) s(a ...interface{}) {
	fmt.Sprint(a...)
}

func (w Writer) mkerr(a string) error {
	return errors.New(a)
}

func (w Writer) mkerrf(str string, a ...interface{}) {
	fmt.Errorf(str, a...)
}

func (w Writer) err(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
}

func (w Writer) errl(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func help() {
	for i := 0; i < len(helpStr); i++ {
		wr.l(helpStr[i])
	}
	os.Exit(0)
}
