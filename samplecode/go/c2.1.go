package main

import "fmt"

/*
练习题2.1 完成以下的数字转换
a. 0x39A7F8转换为二进制
b. 二进制[1100100101111011]转换为十六进制
c. 0xD5E4C转换为二进制
d. 二进制[1001101110011110110101]转换为十六进制
*/
func C2_1() {
	const hexa int = 0x39A7F8
	fmt.Printf("dex:%x bin:%b\n", hexa, hexa)
}
