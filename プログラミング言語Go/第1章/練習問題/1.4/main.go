// 入力に2回以上現れた行の数とテキストを表示。標準入力か指定されたファイル一覧から読み込む。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countFileLines(f, arg, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileNames[line])
		}
	}
}

func countFileLines(f *os.File, fileName string, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		file, _ := fileNames[input.Text()]
		checkFlg := 0
		if len(file) == 0 {
			fileNames[input.Text()] = append(fileNames[input.Text()], fileName)
		} else {
			for _, fileCheck := range file {
				if fileCheck == fileName {
					checkFlg = 1
					break
				}
			}
			if checkFlg == 0 {
				fileNames[input.Text()] = append(fileNames[input.Text()], fileName)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
