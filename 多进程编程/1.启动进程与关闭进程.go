package main

import "os/exec"

func main() {
	//cmd:=exec.Command("notepad")

	//cmd:=exec.Command("taskkill","/f","/im" ,"XMind")
	cmd := exec.Command("ls", "-l")
	cmd.Run()
}

//多进程适用于计算密集型，消耗CPU，
//exe
//exe
//exe

//多线程IO密集型号，爬虫抓取数据，网络速度，文件操作，
