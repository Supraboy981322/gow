package main

import (
	"fmt"
	"os"
	"errors"
)

type (
	Writer struct{}
	Eror struct{}
	Misc struct{}
)
var (
	wr = Writer{}
	er = Eror{}
	op = Misc{}
)

func (w Writer) l(a ...interface{}) {
	fmt.Fprintln(os.Stdout, a...)
}

func (wr Writer) lf(str string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(str, a...))
}

func (w Writer) i(a ...interface{}) {
	fmt.Print(a...)
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

func (w Writer) help() {
	for i := 0; i < len(helpStr); i++ {
		wr.l(helpStr[i])
	}
	os.Exit(0)
}

func (er Eror) han(err error, str string) {
	if err != nil {
		wr.errl(str)
		wr.errl(err)
	}
}

func (er Eror) fan(err error, str string) {
	if err != nil {
		er.han(err, str)
		os.Exit(1)
	}
}

func (er Eror) mk(str string) error {
	return errors.New(str)
}

func (er Eror) ok(ok bool, str string) {
	if !ok {
		err := er.mk(str)
		wr.errl(err)
	}
}

func (er Eror) okMkErr(ok bool, str string) error {
	if !ok {
		return er.mk(str)
	}
	return nil
}

func (op Misc) tern(con1 bool, val1 any, val2 any) interface{} {
	if con1 {
		return val1
	} else {
		return val2
	}
}

func (op Misc) arrToStr(arr []string) string {
	var str string
	for _, val := range arr {
		str = fmt.Sprint(str, val)
	}
	return str
}
