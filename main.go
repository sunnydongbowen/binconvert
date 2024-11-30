package main

import (
	"fmt"
	"os"
)

// @program:   binconvert
// @file:      main.go
// @author:    bowen
// @time:      2024-11-18 20:32
// @description: 进制转换器

// 进制转换
func dToo() {
	// 将十进制转换成其他进制
	fmt.Print("请输入要转换的数字: ")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("该数的二进制为:%b\n", num)
	fmt.Printf("该数的八进制为:0%o\n", num)
	fmt.Printf("该数的十六进制为:0x%x\n", num)
}

func xToo() {
	fmt.Print("请输入要转换的数字(以0x开头): ")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("该数的二进制为:%b\n", num)
	fmt.Printf("该数的八进制为:0%o\n", num)
	fmt.Printf("该数的十进制为:%d\n", num)

}
func oToo() {
	fmt.Print("请输入要转换的数字(以0开头): ")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("该数的二进制为:%b\n", num)
	fmt.Printf("该数的十六进制为:0x%x\n", num)
	fmt.Printf("该数的十进制为:%d\n", num)
}

func bToo() {
	fmt.Print("请输入要转换的数字(以0b开头): ")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("该数的八进制为:0%o\n", num)
	fmt.Printf("该数的十六进制为:0x%x\n", num)
	fmt.Printf("该数的十进制为:%d\n", num)
}

func main() {
	for {
		//var num int
		var choice int
		fmt.Println(`请选择您要进行的操作:
            1、十进制转换成其他进制
            2、八进制转换成其他进制
            3、十六进制转换成其他进制
            4、二进制转换成其他进制
            0、退出程序`)
		//fmt.Println(·1、十进制转换成其他进制" +
		//	" 2、八进制转换成其他进制" +
		//	" 3、十六进制转换成其他进制" +
		//	" 4、二进制转换成其他进制" +
		//	" 0、退出程序")
		fmt.Print("请选择:")
		fmt.Scanln(&choice)
		if choice == 1 {
			dToo()
		} else if choice == 2 {
			oToo()
		} else if choice == 3 {
			xToo()
		} else if choice == 4 {
			bToo()
		} else if choice == 0 {
			os.Exit(1)
		} else {
			fmt.Println("无效输入")
		}
	}
}
