package main

import (
    "flag"
    "os"
    "fmt"
    "greputil"
    "strings"
)

const HELP = `
    go run main.go [-s|-m|-x][-c] PATTERNS FILE
    Search for PATTERNS in one FILE.

    Example:
        go run main.go -s -c "hello" file/test.txt

    PATTERNS can contain multiple strings separated by comma.
`

func main() {

    flag.Bool("Usage", false, HELP)

    flag.Bool("x", false, "Using regexp search in a file")

    flag.Bool("c", false, "Color the output of the function")

    flag.Bool("m", false, "Searching for multiple strings in a file." )

    flag.Bool("s", true,"Searching for a string in a file.")

    flag.Parse()

    color := 0
    
    args :=make([]string, len(os.Args)-1)    //args slice

    copy(args, os.Args[1:])

    if len(args) == 0 || strings.Contains(args[len(args)-2], "-") {    //need to add other conditions

        fmt.Println(greputil.COLORERR, HELP, greputil.COLORRESET)

    }

    if isValueInList("-c", args) {

        color = 1
    
        if isValueInList("-s", args) {

            greputil.SearchSingleStr(args[len(args)-2], args[len(args)-1], color)

        }else if isValueInList("-m", args) {

            greputil.SearchMultipleStr(args[len(args)-2], args[len(args)-1], color)

        }else if isValueInList("-x", args) {

            greputil.SearchRegexp(args[len(args)-2], args[len(args)-1], color)

        }else {

            fmt.Println(greputil.COLORERR, "Parameter '-c' cannot be used alone", greputil.COLORRESET)

        }

    }else {

        if isValueInList("-s", args) {

            greputil.SearchSingleStr(args[len(args)-2], args[len(args)-1], color)

        }else if isValueInList("-m", args) {

            greputil.SearchMultipleStr(args[len(args)-2], args[len(args)-1], color)

        }else if isValueInList("-x", args) {

            greputil.SearchRegexp(args[len(args)-2], args[len(args)-1], color)

        }

    }

}

func isValueInList(value string, list []string) bool {

    for _, v := range list {

        if v == value {

            return true

        }

    }

    return false

}