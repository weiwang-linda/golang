package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "regexp"
    "strings"
)


const HELP = `
    go run main.go [-s|-m|-x][-c] PATTERNS FILE
    Search for PATTERNS in one FILE.

    Example:
        go run main.go -s -c "ERROR" file/engine.log

    PATTERNS can contain multiple strings separated by comma.
`
const COLORPAINT = "\033[34m"
const COLORERR = "\033[31m"
const COLORRESET = "\033[0m"


func check(e error) {

    if e != nil {

        panic(e)

    }
}



func main() {

    //help documnet
    flag.Bool("Usage", false, HELP)

    flag.Bool("x", false, "Using regexp search in a file")

    flag.Bool("c", false, "Color the output of the function")

    flag.Bool("m", false, "Searching for multiple strings in a file." )

    flag.Bool("s", true,"Searching for a string in a file.")

    flag.Parse()


    color := false   //colour bool
    
    args :=make([]string, len(os.Args)-1)    //args slice

    copy(args, os.Args[1:])

    //no args or missing required parameters
    if len(args) == 0 || strings.Contains(args[len(args)-2], "-") { 

        fmt.Println(COLORERR, HELP, COLORRESET)

    }

    index := isValueInList("-c", args)    //location of parameter '-c'

    argsNew := []string{"","","",""}

    if index != 0{

        color = true

        argsNew = append(args[:index], args[index+1:]...)    //remove parameter '-c'

        SearchStr(argsNew[0], argsNew[len(argsNew)-2], argsNew[len(argsNew)-1], color)
        
    }else {

        argsNew = append(args)

        SearchStr(argsNew[0], argsNew[len(argsNew)-2], argsNew[len(argsNew)-1], color)

    }

}

/* colour output function need to be improved later */
func SearchStr(opt string, patterns string, destFile string, colour bool) {

	var clr string

	if colour {

		clr = COLORPAINT

	}

    f, err := os.Open(destFile)

    check(err)

    defer f.Close()

    scanner := bufio.NewScanner(f)

    reg := regexp.MustCompile(patterns)

    for scanner.Scan() {

		var line string = scanner.Text()

        if opt == "-s" {

            if strings.Contains(line, patterns) {

                fmt.Println(clr, line, COLORRESET)

            }

        }else if opt == "-m" {

		    keys := strings.Split(patterns, ",")

            if len(keys) <=1 {

                fmt.Println(COLORERR, "'-m' must have more than 1 searching strings.", COLORRESET)
                
                break

            }else if isStringsInLine(line, keys) {

                fmt.Println(clr, line, COLORRESET)

            }

        }else if opt == "-x" {

            txt := reg.FindString(line)

            if txt == "" {

                continue
            }
                
            fmt.Println(clr, line, COLORRESET)

        }
        
    }

    if err := scanner.Err(); err != nil {

        check(err)

    }
}


func isStringsInLine (line string, arry []string) bool {

	for _, key := range arry {

		if strings.Contains(line, key) {
			
			continue

		}else {

			return false

		}

	}

	return true

}

func isValueInList(value string, list []string) int {

    for i, v := range list {

        if v == value {

            return i

        }

    }

    return 0

}