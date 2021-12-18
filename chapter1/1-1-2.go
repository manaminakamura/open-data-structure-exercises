package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		先頭から50行の入力を読み、それを逆順で出力せよ。その後、続く50行を読み、それを逆順で出力せよ。
		これを読み取る行がなくなるまで繰り返し、最後に残った行（50行未満かもしれない）もやはり逆順で出力せよ。
		なお、プログラムの実行中に50行より多くの行を保持してはならない。
	*/
	data, err := os.Open("chapter1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	flag := true
	for {
		flag = scanUpto50lines(scanner, flag)
		if !flag {
			break
		}
	}
}

func scanUpto50lines(scanner *bufio.Scanner, flag bool) bool {
	array := []string{}
	for i := 0; i < 50; i++ {
		if scanner.Scan() {
			array = append(array, scanner.Text())
		} else {
			flag = false
			break
		}
	}
	for i := len(array); i > 0; i-- {
		fmt.Println(array[i-1])
	}
	return flag
}
