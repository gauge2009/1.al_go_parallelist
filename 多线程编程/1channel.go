package main

import (
	"fmt"
	"time"
)

//make和new的区别
//new的作用是初始化一个指向类型的指针（*T）。使用new函数来分配空间，传递给new函数的是一个类型，不是一个值。返回的是指向这个新分配的零值的指针。
//make的作用是为slice、map或chan初始化并返回引用（T）。make仅仅用于创建slice、map和channel，并返回它们的实例。
func main() {
	//1G文件，10份，搜索yincheng,结果归并
	var strChan = make(chan string, 3)
	//上面是 chan 通道的声明，对比下面数组的声明
	var strCham = make([]string, 3)
	fmt.Printf(" chan 通道的声明，对比数组的声明：strCham length = %v\n", len(strCham))
	dic := make(map[string]int)
	fmt.Println(" chan 通道的声明，对比字典的声明:", dic)
	syncChannel1 := make(chan struct{}, 1)
	syncChannel2 := make(chan struct{}, 2)
	//读取
	go func() {
		<-syncChannel1
		fmt.Println("syncChannel1收到了信号")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("收到", elem, "recv")
			} else {
				break
			}
		}

		fmt.Println("停止接收")
		syncChannel2 <- struct{}{}
	}()
	//写入
	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem //压入数据到
			fmt.Println("Send", elem, "sender")
			if elem == "c" {
				syncChannel1 <- struct{}{}
				fmt.Println("send signal sender")
			}
		}
		fmt.Println("wait 2 s")
		time.Sleep(time.Second * 2)
		close(strChan) //关闭
		syncChannel2 <- struct{}{}

	}()
	<-syncChannel2
	<-syncChannel2

}
