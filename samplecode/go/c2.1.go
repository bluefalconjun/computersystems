package main

import (
	"fmt"
	"unsafe"
)

/*
all in C2.1
*/
func C2() {
	c21()
	c24()
	c25()
}

/*
练习题2.1 完成以下的数字转换
a. 0x39A7F8转换为二进制
b. 二进制[1100100101111011]转换为十六进制
c. 0xD5E4C转换为二进制
d. 二进制[1001101110011110110101]转换为十六进制
*/
func c21() {
	fmt.Println("c21")
	defer fmt.Println("c21---done")
	const hexa int = 0x39A7F8
	fmt.Printf("dex:%x bin:%b\n", hexa, hexa)

	const binb string = "1100100101111011"
	//实现不写了,麻烦:(
}

/*
练习题2.4 实现以下c代码

typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, size_t len){
	for(i=0; i<len; i++){
		printf("%.2x", start[i]);
	}
	printf("\n");
}

void show_int(int x){
	show_bytes((byte_pointer) &x, sizeof(int));
}

void show_float(float x){
	show_bytes((byte_pointer) &x, sizeof(float));
}

void show_pointer(void* x){
	show_bytes((yte_pointer) &x, sizeof(void *));
}

int ival = 12345;
float fval = (float)ival;
int *pval = &ival;

show_int(ival);
show_float(fval);
show_pointer(pval);

*/
type bytePointer = byte

func showbytes(start *bytePointer, len int) {

	/*
		注意: go语言中不支持指针的其他运算, 只有 &/* 这两种操作,
		当前情况下, 需要将start地址向后移, 使用unsafe.Pointer获取地址,并使用uintptr来将地址转换为整数,后续的+1即将地址向后移动 one 字节.
		如果需要按照c语言中的按类型大小移动. 使用:
		unsafe.Pointer(uintptr(unsafe.Pointer(&p) + unsafe.Sizeof(p)))
	*/
	var curbyte *byte
	for i := 0; i < len; i++ {

		curbyte = (*byte)(unsafe.Pointer((uintptr(unsafe.Pointer(start)) + (uintptr)(i))))
		fmt.Printf("%.2x", *curbyte)
	}
	fmt.Printf("\n")
}

/*
考虑到同一接口,只是传入参数不同,这里应该使用Interface来实现.?
*/
func showint(i int) {
	showbytes((*byte)(unsafe.Pointer(&i)), (int)(unsafe.Sizeof(i)))
}

func showfloat(i float32) {
	showbytes((*byte)(unsafe.Pointer(&i)), (int)(unsafe.Sizeof(i)))
}

func showpointer(i uintptr) {
	showbytes((*byte)(unsafe.Pointer(&i)), (int)(unsafe.Sizeof(i)))
}

func c24() {
	fmt.Println("c24")
	defer fmt.Println("c24---done")

	const ival int = 12345
	fval := float32(ival)
	pval := uintptr(ival)

	showint(ival)
	showfloat(fval)
	showpointer(pval)

}

/*
练习题2.5 思考下面对show_bytes的三次调用
int ival = 0x87654321
byte_pointer pval = (byte_pointer) &ival;

show_bytes(pval, 1);
show_bytes(pval, 2);
show_bytes(pval, 3);

在大端和小端机器上,输出分别是什么?
*/
func c25() {
	fmt.Println("c25")
	defer fmt.Println("c25---done")

	var ival int
	ival = 0x87654321
	pval := (*byte)(unsafe.Pointer(&ival))

	showbytes(pval, 1)
	showbytes(pval, 2)
	showbytes(pval, 3)
}

/*
练习题2.7 下面对show_bytes的调用将输出什么结果

const char *s = "abcdef";
show_bytes((byte_pointer)s, strlen(s));
*/

/*
练习题2.10 对任一向量a 有 a^a=0, 应用这一属性, 考虑以下程序:

void inplace_swap(int *x, iny *y){
	*y = *x ^ *y;
	*x = *x ^ *y;
	*y = *x ^ *y;
}
*/

/*
练习题2.11 基于2.10的函数, 实现数组元素头尾对调

void reverse_array(int a[], int cnt){
	//reverse array a[]
}
*/

/*
练习题2.13
考虑没有布尔运算符 AND 和 OR 指令的机器中, 使用

int bis(int x, int m): 返回在m为1的每个位上,设置x对应位为1的值
int bic(int x, int m): 返回在m为1的每个位上,设置x对应位为0的值

来实现 | ^ 指令

int bool_or(int x, int y){
	return ____;
}

int bool_xor(int x, int y){
	return ____;
}
*/
