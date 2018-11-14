package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		_, err := readLine(r)
		if err != nil {
			break
		}
	}

}

func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}

func main() {

	f, _ := os.Open("./order.sql")
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		e, err := readLine(r)
		if err != nil {
			break
		}
		fmt.Println("==================================================================================================================================================")
		fmt.Println(e)
	}
}
