package main

import (
	"os"
	"strings"
//	"net/http"
)

type (
	InvalidArg struct {
		Value int
		Exists bool
	}
)

var (
	args = os.Args[1:]
	silent bool
	url string
	invalidArg InvalidArg
	help bool

	helpStr = []string{
		"\033[38mGOw\033[0m --> \033[33mhelp\033[0m",
		"  \033[32m-h\033[0m",
		"    help (responds with this)",
		"  \033[35m-s\033[0m",
		"    silent (outputs only response body",
	}
)

func main() {
	strRemaining := false
	for i := 0; i < len(args); i++ {
		curArg := args[i]
		if !strRemaining {
			switch curArg[0] {
			case '-':
				strSplit := strings.Split(curArg, "")
				strRemaining = readArgChars(strSplit)
				if invalidArg.Exists {
					errOut(wr.mkerr("arg"), curArg, invalidArg.Value)
				}
				if help {
					wr.i("\033[0m")
					wr.help()
				}
			default:
				url = curArg
			}
		} else {
			url += curArg
		}
	}
}

func readArgChars(arg []string) bool { 
	strRemaining := false
	for i := 0; i < len(arg); i++ {
		switch (arg[i]) {
		case "-":
			strRemaining = true
			break
		case "s":
			silent = true
		case "h":
			help = true
		default:
			invalidArg.Value = i
			invalidArg.Exists = true
			help = true
			return strRemaining
		}
	}
	return strRemaining
}

func errOut(err error, str any, ext ...interface{}) {
	if index, ok := ext[0].(int); ok && err.Error() == "arg" {
		var pointer string
		for i := 0; i < index; i++ {
			pointer += " "
		}
		pointer += "\033[1;31m^\033[0m"
		str, ok := str.(string)
		if !ok {
			wr.errl("failed to convert str type of "+
							"any to type of string in errOut()")
			os.Exit(1)
		}

		str = str[:index] + "\033[1;31m" +
					string(str[index]) +
					"\033[0m" + str[index+1:]

		wr.errl("invalid arg:\n" +
						"  " + str + "\n" +
			      "  " + pointer + "\n")
	} else {
		
	}
}
