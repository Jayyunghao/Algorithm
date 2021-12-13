package Concurrent

import (
	"fmt"
	"sync"
	"time"
)

/*
交替打印数字和字.
问题描述
使.两个 goroutine 交替打印序列，.个 goroutine 打印数字， 另外.
个 goroutine 打印字.， 最终效果如下：
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/

type token struct {
}
var wg sync.WaitGroup
var n int = 1
func AlternatePrinting() {
	taskName := []string{"number","words"}
	tokens := []chan token{make(chan token,1),make(chan token,1)}
	for i:=0;i<len(tokens);i++ {
		wg.Add(1)
		go printCore(taskName[i], tokens[i], tokens[(i+1)%2])
	}
	tokens[0] <- struct{}{}
	wg.Wait()
}

func printCore(name string, token chan token, next chan token) {
	defer wg.Done()
	switch name {
	case "number":
		for {
			select {
			case t := <-token:
				for i :=0; i < 2;i++ {
					fmt.Print(n)
					n++
				}
				next <- t
				time.Sleep(time.Second * 2)
			}
		}
	case "words":
		count := 0
		for {
			select {
			case t := <-token:
				for i:=0;i<2;i++ {
					s := string('A' + count)
					fmt.Print(s)
					count = (count+1)%26
				}
				next <- t
				time.Sleep(time.Second * 2)
			}
		}
	}
}

/*
func main() {
	AlternatePrinting()
}
*/
