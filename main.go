package main

import (
	"os"
	"time"
	"io/ioutil"
	"strings"
	"net/http"
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
	secure bool
	url string
	invalidArg InvalidArg
	help bool
	headKeys []string
	headVals []string

	helpStr = []string{
		"\033[1;36mGOw\033[0m \033[1m-->\033[0m \033[1;33mhelp\033[0m",
		"  \033[1;32m-h\033[0m",
		"    help (responds with this)",
		"  \033[1;35m-s\033[0m",
		"    silent (outputs only response body)",
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
				strRemaining = readArgChars(strSplit, i)
				if invalidArg.Exists {
					errOut(wr.mkerr("arg"), curArg, invalidArg.Value)
				}
				if help {
					wr.help()
				}
			default:
				url = curArg
			}
		} else {
			url += curArg
		}
	}
 
	if len(args) == 0 {
		wr.errl("not enough args")
		os.Exit(1)
	}

	if url == "" {
		wr.errl("no url provided")
		os.Exit(1)
	} 

	res, status, code, err := mkReq()
	if !silent {
		wr.lf("status:  %s / %d", status, code)
	}
	er.fan(err, res)

	wr.l(res)
}

func readArgChars(arg []string, cur int) bool { 
	strRemaining := false
	for i := 0; i < len(arg); i++ {
		switch (arg[i]) {
		case "-":
			strRemaining = true
			break
		case "s":
			silent = true
		case "S":
			secure = true
		case "h":
			help = true
		case "H":
			headKeys[len(headKeys)] = args[cur+1]
			headVals[len(headVals)] = args[cur+2]
		default:
			invalidArg.Value = i
			invalidArg.Exists = true
			help = true
			return strRemaining
		}
	}
	return strRemaining
}

func mkReq() (string, string, int, error){
	wr.l(url)
	urlSplit := strings.Split("://", url)
	if len(urlSplit) == 1 {
		urlSplit = append(urlSplit, url)
		reqType, ok := op.tern(secure, "https", "http").(string)
		if !ok {
			return "converting type of any to " +
						 "string for setting request type",
							"" , 0, er.mk("")
		}
		urlSplit[0] = reqType
	}

	url = urlSplit[0] + "://" + urlSplit[1]

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	wr.l(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "err creating request", "", 0, err
	}

	req.Header.Add("User-Agent", "someCrappyCliTool")
	for i := 0; i < len(headKeys); i++ {
		req.Header.Add(headKeys[i], headVals[i])
	}

	resp, err := client.Do(req)
	stat := resp.StatusCode
	statTxt := http.StatusText(stat)
	if err != nil {
		return "err doing request", statTxt, stat, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "err reading body", statTxt, stat, err
	}


	return string(body), statTxt, stat, nil
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
