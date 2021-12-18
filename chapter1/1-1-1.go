package main

import (
	"bufio"
	"fmt"
	"os"
)

// 入力を1行ずつ読み、その逆順で出力せよ。すなわち、最後の入力行を最初に書き出し、最後から2行めを2番めに書き出す、というように出力せよ
func main() {
	data, _ := os.Open("input.txt")
	defer data.Close()
	scanner := bufio.NewScanner(data)
	array := []string{}
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	for i := len(array); i > 0; i-- {
		fmt.Println(array[i-1])
	}
}
