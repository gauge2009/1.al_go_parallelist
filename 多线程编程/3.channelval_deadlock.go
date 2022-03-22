package main

import (
	"fmt"
	"time"
)

var mapChanl = make(chan map[string]int, 1) //引用类型

func main() {
	syncChan := make(chan struct{}, 2)        // 该通道用于流水线同步【即：生产者和消费者都完成后，主线程再结束】
	syncChanB := make(chan map[string]int, 1) //该通道用于消费者通知生产者篡改结果
	go func() {
		for {
			if elem, ok := <-mapChanl; ok {
				elem["zrz1898zrz的余额"]++
			} else {
				break
			}
		}
		fmt.Println("stop 接收")

		////我们改写一下，让生产者生产后改变map 中 count 键的 值为 10，并用另一个通道 syncChanB 来通知 消费者：
		if elem, ok := <-mapChanl; ok {
			elem["count"] = 100

			fmt.Println("篡改完成")
			syncChanB <- elem //struct{}{} //文儿压入：通知生产者消费结束，且数据被修改
		}
		syncChan <- struct{}{} //双小吱 压入：消费者通知线程

	}()

	go func() {
		countmap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChanl <- countmap //压入数据
			time.Sleep(time.Microsecond)
			fmt.Printf("countmap %v万\n", countmap)
		}
		close(mapChanl)

		////我们改写一下，让生产者生产后改变map 中 count 键的 值为 10，并用另一个通道 syncChanB 来通知 消费者：
		if newMap, isSuccess := <-syncChanB; isSuccess { //文儿等待读取：收到消费完成通知
			fmt.Printf("篡改完成效果： countmap %v\n", countmap)

			fmt.Printf("篡改完成效果： 通道内新值 %v\n", newMap)
		} else {
			return
		}
		//syncChan <- struct{}{}
		syncChan <- struct{}{} //双小吱压入：生产者通知线程

	}()

	<-syncChan //双小吱等待读取
	<-syncChan //双小吱等待读取

}
