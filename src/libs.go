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
	wr = Writer{} //fmt
	er = Eror{}   //error
	op = Misc{}   //operations
)

//read it
func (w Writer) l(a ...interface{}) {
	fmt.Fprintln(os.Stdout, a...)
}

//read it
func (wr Writer) lf(str string, a ...interface{}) {
	fmt.Println(fmt.Sprintf(str, a...))
}

//read it
func (w Writer) i(a ...interface{}) {
	fmt.Print(a...)
}

//read it
func (w Writer) f(str string, a ...interface{}) {
	fmt.Printf(str, a...)
}

//read it
func (w Writer) sf(str string, a ...interface{}) {
	fmt.Sprintf(str, a...)
}

//read it
func (w Writer) s(a ...interface{}) {
	fmt.Sprint(a...)
}

//read it
func (w Writer) mkerr(a string) error {
	return errors.New(a)
}

//read it
func (w Writer) mkerrf(str string, a ...interface{}) {
	fmt.Errorf(str, a...)
}

//read it
func (w Writer) err(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
}

//read it
func (w Writer) errl(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

//fmt.Printf("\n")
func (w Writer) b() {
	wr.l()
}

//fmt.Println("...") on each item of string array
func (w Writer) help() {
	for i := 0; i < len(helpStr); i++ {
		wr.l(helpStr[i])
	}
	os.Exit(0)
}

//if err != nil { fmt.Fprintln(os.Stderr, "..."); fmt.Fprintln(err) }
func (er Eror) han(err error, str string) {
	if err != nil {
		wr.errl(str)
		wr.errl(err)
	}
}

//if err != nil { fmt.Fprintln(os.Stderr, "..."); fmt.Fprintln(err); os.Exit(1) }
func (er Eror) fan(err error, str string) {
	if err != nil {
		er.han(err, str)
		os.Exit(1)
	}
}

//errors.New("...")
func (er Eror) mk(str string) error {
	return errors.New(str)
}

//if !ok { fmt.Fprintln(os.Stderr, "...") }
func (er Eror) ok(ok bool, str string) {
	if !ok {
		err := er.mk(str)
		wr.errl(err)
	}
}

//if !ok { return errors.New("...") }
func (er Eror) okMkErr(ok bool, str string) error {
	if !ok {
		return er.mk(str)
	}
	return nil
}

//ternary-like func
func (op Misc) tern(con bool, val1 any, val2 any) interface{} {
	if con {
		return val1
	} else {
		return val2
	}
}

//array to string
func (op Misc) arrToStr(arr []string) string {
	var str string
	for _, val := range arr {
		str = fmt.Sprint(str, val)
	}
	return str
}
