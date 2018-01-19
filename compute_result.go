package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	args := os.Args //获取用户输入的所有参数
	if args == nil || len(args) < 1 {
		Usage() //如果用户没有输入,或参数个数不够,则调用该函数提示用户
		return
	}
	startStr := args[1] //获取输入的第一个参数
	startValue, err := strconv.ParseFloat(startStr, 64)
	if err != nil {
		fmt.Println("请输入有效数字！")
	} else {
		endValue := startValue * 1.005
		fmt.Println("初始价格是：", startValue)
		fmt.Println("最低出手价：", endValue)
	}
}

var Usage = func() {
	fmt.Println("初始价格是？")
}
