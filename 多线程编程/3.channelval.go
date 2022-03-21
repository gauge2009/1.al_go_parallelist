package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1) //引用类型

func main() {
	syncChan := make(chan struct{}, 2) // 该通道用于流水线同步【即：生产者和消费者都完成后，主线程再结束】

	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["zrz1898zrz的余额"]++
			} else {

				break
			}
		}
		fmt.Println("stop 接收")

		syncChan <- struct{}{} //双小吱 压入：消费者通知线程

	}()

	go func() {
		countmap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countmap //压入数据
			time.Sleep(time.Microsecond)
			fmt.Printf("countmap %v万\n", countmap)
		}
		close(mapChan)

		syncChan <- struct{}{} //双小吱压入：生产者通知线程

	}()

	<-syncChan //双小吱等待读取
	<-syncChan //双小吱等待读取

}
