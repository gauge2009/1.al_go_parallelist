package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

//返回字符串
func (c Counter) String() string {
	return fmt.Sprintf("{count=%d}", c.count)
}

var mapChannel = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChannel; ok {
				c := elem["zrz1898zrz的余额"]
				c.count++
			} else {
				break
			}
		}
		fmt.Println("stop 接收")

		syncChan <- struct{}{}
	}()

	go func() {
		//countmap:=map[string]Counter{"count":Counter{}}
		countmap := map[string]*Counter{"count": &Counter{}}
		for i := 0; i < 5; i++ {
			mapChannel <- countmap //压入数据
			time.Sleep(time.Microsecond)
			fmt.Printf("countmap %v\n", countmap)
		}
		close(mapChannel)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}
