package main

import (
	"os"
	"time"
	"io/ioutil"
	"strings"
	"slices"
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
	method = "GET"
	url string
	invalidArg InvalidArg
	help bool
	headKeys []string
	headVals []string
	parsedArgs []int

	helpStr = []string{
		"\033[1;36mGOw\033[0m \033[1m-->\033[0m \033[1;33mhelp\033[0m",
		"  \033[1;32m-h\033[0m",
		"    help (responds with this)",
		"  \033[1;35m-s\033[0m",
		"    silent (outputs only response body)",
		"  \033[1;33m-S\033[0m",
		"    secure (https only)",
		"  \033[1;34m-H\033[0m",
		"    header",
		"      usage:  \033[1;34m-H \"[key]\" \"[value]\"\033[0m",
		"      eg:  \033[1;37mgow \033[1;34m-H "+
						"\"Content-Type\" \"text/json\""+
						"\033[1;37mhttps://example.com\033[0m",
	}
)

func main() {
	var strRemaining bool
	var parseHeader bool
	for i := 0; i < len(args); i++ {
		curArg := args[i]
		parsed := slices.Contains(parsedArgs, i)
		if !strRemaining && !parsed {
			switch curArg[0] {
			case '-':
				strSplit := strings.Split(curArg, "")[1:]
				strRemaining, parseHeader = readArgChars(strSplit, i)
				if invalidArg.Exists {
					errOut(wr.mkerr("arg"), curArg, invalidArg.Value)
				}
				if help {
					wr.help()
				}
				if parseHeader {
					if i+2 > len(args)-1 {
						wr.errl("\033[1;31mheader incomplete\033[0m")
						os.Exit(1)
					} else {
						headKeys = append(headKeys, args[i+1])
						headVals = append(headVals, args[i+2])
						parsedArgs = append(parsedArgs, i+1, i+2)
						parseHeader = false
					}
				}
			default:
				url = curArg
			}
		} else if !parsed {
			url += args[i] + "%20"
		}
	}

	//if the args were stringed together
	//  remove the last %20
	if strRemaining {
		url = url[:len(url)-3]
	}
 
	if len(args) == 0 {
		wr.errl("\033[1;31mnot enough args\033[0m")
		os.Exit(1)
	}

	if url == "" {
		wr.errl("\033[1;31mno url provided\033[0m")
		os.Exit(1)
	} 

	res, status, code, err := mkReq()
	if !silent {
		wr.l(url)
		wr.lf("status:  %s / %d", status, code)
	}
	er.fan(err, res)

	wr.l(res)
}

func readArgChars(arg []string, cur int) (bool, bool) { 
	var parseForHeader bool
	var strRemaining bool
	if arg[0] == "-" {
		if len(arg) == 1 {
			strRemaining = true
		} else {
			argCut := op.arrToStr(arg[1:])
			switch argCut {
			case "PUT":
				method = argCut
			default:
				wr.errl("invalid arg:  \033[1;31m-" +
						op.arrToStr(arg) + "\033[0m\n")
				wr.help()
			}
		}
		strRemaining = false
	} else {
		for i := 0; i < len(arg); i++ {
			if strRemaining {
				break
			}
			switch (arg[i]) {
			case "-":
				strRemaining = true
			case "s":
				silent = true
			case "S":
				secure = true
			case "h":
				help = true
			case "p":
				method = "POST"
			case "g":
				method = "GET"
			case "H":
				parseForHeader = true
			default:
				invalidArg.Value = i+1
				invalidArg.Exists = true
				help = true
				return strRemaining, false
			}
		}
	}
	return strRemaining, parseForHeader
}

func mkReq() (string, string, int, error){
	urlSplit := strings.Split(url, "://")
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

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "err creating request", "", 0, err
	}

	req.Header.Add("User-Agent", "someCrappyCliTool")
	for i := 0; i < len(headKeys); i++ {
		req.Header.Add(headKeys[i], headVals[i])
	}

	resp, err := client.Do(req)
	if err != nil {
		return "err doing request", "client failed", 0, err
	}

	stat := resp.StatusCode
	statTxt := http.StatusText(stat)

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
