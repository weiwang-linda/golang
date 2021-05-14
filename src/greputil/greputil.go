package greputil

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"regexp"
)

const COLORPAINT string = "\033[34m"
const COLORERR string = "\033[31m"
const COLORRESET string = "\033[0m"

func check(e error) {

    if e != nil {

        panic(e)

    }
}

func SearchSingleStr(singleKey string, destFile string, colour int) {

	var clr string

	if colour == 1 {

		clr = COLORPAINT

	}

    f, err := os.Open(destFile)

    check(err)

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

		var line string = scanner.Text()
        
		if strings.Contains(line, singleKey){

			fmt.Println(clr, line, COLORRESET)

		}

    }

    if err := scanner.Err(); err != nil {

        check(err)

    }
}



func SearchMultipleStr(multiKey string, destFile string, colour int) {

	var clr string

	if colour == 1 {

		clr = COLORPAINT

	}

	f, err := os.Open(destFile)

    check(err)

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

		var line string = scanner.Text()

		keys := strings.Split(multiKey, ",")

		if len(keys) <=1 {

			fmt.Println(COLORERR, "'-m' must have more than 1 searching strings.", COLORRESET)
			
			break

		}else if isStringsInLine(line, keys) {

			fmt.Println(clr, line, COLORRESET)

		}
		
    }

    if err := scanner.Err(); err != nil {

        check(err)

    }

}

func SearchRegexp(regExpMatch string, destFile string, colour int) {

	var clr string

	if colour == 1 {

		clr = COLORPAINT

	}
	
	f, err := os.Open(destFile)

    check(err)

    defer f.Close()

    scanner := bufio.NewScanner(f)

	reg := regexp.MustCompile(regExpMatch)

    for scanner.Scan() {

		var line string = scanner.Text()
        
		txt := reg.FindString(line)

		if txt == "" {

			continue
		}
			
		fmt.Println(clr, line, COLORRESET)
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
