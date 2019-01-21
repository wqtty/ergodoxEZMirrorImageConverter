package main

import (
	"os"
	"bufio"
	"regexp"
	"strings"
	"unicode"
	"bytes"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		println("请输入待转换的ergodox ez生成的.c文件")
		os.Exit(1)
	}
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		println("打开文件" + os.Args[0] + "失败！")
		os.Exit(2)
	}
	outFile, err := os.OpenFile("keymap.c", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		println("打开文件keymap.c失败，请确认当前目录可写")
		os.Exit(3)
	}
	fs := bufio.NewScanner(inFile)
	regex, err := regexp.Compile(".*LAYOUT_ergodox\\(.*")
	if err != nil {
		println("请检查正则表达式语法")
		os.Exit(3)
	}
	lineNo := 0
	for fs.Scan() {
		txt := fs.Text()
		if regex.MatchString(txt) {
			splitted := strings.Split(strings.Map(func(r rune) rune {
				if unicode.IsSpace(r) {
					return -1
				}
				return r
			}, txt), ",")
			splitted[0] = splitted[0][strings.Index(splitted[0], "(")+1:]
			splitted[len(splitted)-2] = splitted[len(splitted)-2][:len(splitted[len(splitted)-2])-1]
			transform(splitted)
			//76 keys
			txt = strings.Join(splitted[:76], ",")
			var buffer bytes.Buffer
			buffer.WriteString("[")
			buffer.WriteString(strconv.Itoa(lineNo))
			buffer.WriteString("]=LAYOUT_ergodox(")
			buffer.WriteString(txt)
			buffer.WriteString("),")
			outFile.WriteString(buffer.String())
			lineNo++
		} else {
			outFile.WriteString(txt)
			outFile.WriteString("\n")
		}
	}
}

func transform(splitted []string) {
	//process Mod-Tap Keys
	for i := 0; i < len(splitted); i++ {
		element := splitted[i]
		if strings.Contains(element, "(") && ! strings.Contains(element, ")") {
			temp := splitted[:i]
			temp = append(temp, strings.Join(splitted[i:i+2], ","))
			temp = append(temp, splitted[i+2:]...)
			splitted = temp
		}
	}
	//generate mirror image for ergodox pro
	//the first row
	swap(splitted, 0, 44)
	swap(splitted, 1, 43)
	swap(splitted, 2, 42)
	swap(splitted, 3, 41)
	swap(splitted, 4, 40)
	swap(splitted, 5, 39)
	swap(splitted, 6, 38)
	//the second row
	swap(splitted, 7, 51)
	swap(splitted, 8, 50)
	swap(splitted, 9, 49)
	swap(splitted, 10, 48)
	swap(splitted, 11, 47)
	swap(splitted, 12, 46)
	swap(splitted, 13, 45)
	//the third row
	swap(splitted, 14, 57)
	swap(splitted, 15, 56)
	swap(splitted, 16, 55)
	swap(splitted, 17, 54)
	swap(splitted, 18, 53)
	swap(splitted, 19, 52)
	//the forth row
	swap(splitted, 20, 64)
	swap(splitted, 21, 63)
	swap(splitted, 22, 62)
	swap(splitted, 23, 61)
	swap(splitted, 24, 60)
	swap(splitted, 25, 59)
	swap(splitted, 26, 58)
	//the fifth row
	swap(splitted, 27, 69)
	swap(splitted, 28, 68)
	swap(splitted, 29, 67)
	swap(splitted, 30, 66)
	swap(splitted, 31, 65)
	//thumb pad
	swap(splitted, 32, 71)
	swap(splitted, 33, 70)
	swap(splitted, 34, 72)
	swap(splitted, 35, 75)
	swap(splitted, 36, 74)
	swap(splitted, 37, 73)
}

func swap(arr []string, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
