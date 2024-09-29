package main

import (
	"fmt"
)

func main() {

	str := "NEGIE12"

	strRev := reverseFunc(str)
	fmt.Println(str)
	fmt.Println(strRev)
	fmt.Println()

	str2 := "Saya sangat senang mengerjakan soal algoritma"
	longest, length := longestString(str2)
	fmt.Println(str2)
	fmt.Println(longest, length)
	fmt.Println()

	str3Input := []string{`xc`, `dz`, `bbb`, `dz`}
	str3Query := []string{`bbb`, `ac`, `dz`}
	queryCount := queryEachInput(str3Input, str3Query)
	output := ""
	for index, teks := range queryCount {
		if teks < 1 && index != len(queryCount)-1 {
			output = output + fmt.Sprintf("kata  '%s' tidak ada %v pada INPUT, ", str3Query[index], teks)
		} else if index != len(queryCount)-1 {
			output = output + fmt.Sprintf("kata  '%s'  terdapat %v pada INPUT, ", str3Query[index], teks)
		}
		if teks < 1 && index == len(queryCount)-1 {
			output = output + fmt.Sprintf("kata  '%s' tidak ada %v pada INPUT ", str3Query[index], teks)
		} else if index == len(queryCount)-1 {
			output = output + fmt.Sprintf("kata  '%s'  terdapat %v pada INPUT ", str3Query[index], teks)
		}
	}
	fmt.Println("INPUT ", str3Input)
	fmt.Println("QUERY ", str3Query)
	fmt.Println("OUTPUT ", queryCount, " karena ", output)
	fmt.Println()

	str4InputMatrix := [][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}
	firstDiag, secondDiag := sumDiagonal(str4InputMatrix)
	firstSum := 0
	secondSum := 0
	outputFirstDiagonal := ""
	outputSecondDiagonal := ""
	//fmt.Println("diagonal pertama = ")
	for index, v := range firstDiag {
		if index == len(firstDiag)-1 {
			outputFirstDiagonal = outputFirstDiagonal + fmt.Sprintf("%d", v)
		} else {
			outputFirstDiagonal = outputFirstDiagonal + fmt.Sprintf("%d + ", v)
		}
		firstSum = firstSum + v
	}
	for index, v := range secondDiag {
		if index == len(firstDiag)-1 {
			outputSecondDiagonal = outputSecondDiagonal + fmt.Sprintf("%d", v)
		} else {
			outputSecondDiagonal = outputSecondDiagonal + fmt.Sprintf("%d + ", v)
		}
		secondSum = secondSum + v
	}
	fmt.Printf("diagonal pertama = %s ", outputFirstDiagonal)
	fmt.Println()
	fmt.Printf("diagonal kedua = %s ", outputSecondDiagonal)
	fmt.Println()
	fmt.Printf("maka hasilnya adalah %v - %v = %v", firstSum, secondSum, (firstSum - secondSum))
}
