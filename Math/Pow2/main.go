package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_,err := fmt.Scanf("%d",&n)
	if err != nil {
		fmt.Sprintf("err :%s",err.Error())
		return
	}
	s := Pow2(n)
	fmt.Printf("2的%d次方为:%s",n,s)
	return
}

func Pow2(n int) string {
	if n < 0 {return ""}
	if n ==0 {return "1"}
	s := []byte{'1'}
	for i:=1;i<=n;i++ {
		flag := 0  //进位
		for j:=0;j<len(s);j++ {
			num ,_ := strconv.Atoi(string(s[j]))
			temp := num * 2 + flag
			flag = temp / 10
			temp = temp % 10
			s[j] = []byte(strconv.Itoa(temp))[0]
		}
		if flag != 0 {
			s = append(s,[]byte(strconv.Itoa(flag))[0])
		}
	}
	result :=""
	for i:=len(s)-1;i>=0;i-- {
		result += string(s[i])
	}
	return result
}
