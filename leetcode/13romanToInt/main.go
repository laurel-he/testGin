package main

import "fmt"

func romanToInt(s string) int {
	/*
		需要处理六种额外情况：
		 IV=4 IX=9
		 XL=40 XC=90
		 CD=400 CM=900
		即每次判断是否是需要的数据时，判断前一位和当前位能否组成额外的含义
	*/
	var prev rune
	res := 0
	// I， V， X， L，C，D 和 M
	for _, v := range s {
		// 三种额外情况，单独计算，需要减两次，因为之前加过
		if prev == 'I' {
			if v == 'V' || v == 'X' {
				res--
				res--
			}
		}
		if prev == 'X' {
			if v == 'L' || v == 'C' {
				res -= 10
				res -= 10
			}
		}
		if prev == 'C' {
			if v == 'D' || v == 'M' {
				res -= 100
				res -= 100
			}
		}
		prev = v
		switch v {
		case 'I':
			res += 1
		case 'V':
			res += 5
		case 'X':
			res += 10
		case 'L':
			res += 50
		case 'C':
			res += 100
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
	}
	return res
}

func main() {
	res := romanToInt("MCMXCIV")
	fmt.Println(res)
}
