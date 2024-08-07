package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	majorFlag := flag.Bool("M", false, "Major release x.0.0")
	minorflag := flag.Bool("m", false, "Minor release 1.x.0")
	revflag := flag.Bool("r", false, "Revision 1.0.x")
	flag.Parse()

	var bs []bool
	if !*majorFlag && !*minorflag && !*revflag {
		//Revision is default if no flags set.
		bs = append(bs, false, false, true)
	} else {
		bs = append(bs, *majorFlag, *minorflag, *revflag)
	}
	incrementVersion("VERSION", bs)
}

func incrementVersion(filename string, bs []bool) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	saned := strings.TrimSpace(strings.ReplaceAll(string(file), "\n", ""))
	split := strings.Split(saned, ".")

	if len(split) != 3 {
		fmt.Println("Version format incorrect. should be x.x.x eg 420.6.9")
		panic(err)
		return
	}

	var is []int
	for _, i := range split {
		conv, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Something went wrong, version has a letter in it?")
			panic(err)
			return
		}
		is = append(is, conv)
	}

	for index, _ := range is {
		is[index] += Btoi(bs[index])
	}

	com := fmt.Sprintf("%v.%v.%v", is[0], is[1], is[2])

	fmt.Printf("Version: %v -> %v", string(file), com)

	wfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer wfile.Close()
	_, err = wfile.WriteString(com)
	if err != nil {
		panic(err)
	}
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
