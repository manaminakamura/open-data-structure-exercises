package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		入力を1行ずつ読み取り、42行め以降で空白を見つけたら、その42行前の行を出力せよ。
		例えば、242行目が空行であれば、200行目を出力せよ。なお、プログラム実行中に43行以上の行を保持してはならない。
	*/
	data, _ := os.Open("chapter1/input.txt")
	scanner := bufio.NewScanner(data)
	array := []string{}
	lineNum := 0
	for scanner.Scan() {
		lineNum += 1
		if lineNum < 43 {
			array = append(array, scanner.Text())
		} else {
			if scanner.Text() == "" {
				fmt.Println(array[0])
			}
			array = append(array[1:], scanner.Text())
		}
	}
}
