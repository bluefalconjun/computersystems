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


/*
练习题2.5 思考下面对show_bytes的三次调用
int ival = 0x87654321
byte_pointer pval = (byte_pointer) &ival;

show_bytes(pval, 1);
show_bytes(pval, 2);
show_bytes(pval, 3);

在大端和小端机器上,输出分别是什么?
*/


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

