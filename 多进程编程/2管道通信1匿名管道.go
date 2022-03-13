package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

//文件缓冲区
func RunCMD(cmd string, args string) {
	//cmd0:=exec.Command("tasklist")
	//cmd0:=exec.Command("ls","-l")
	cmd0 := exec.Command(cmd, args)
	//cmd0:=exec.Command("ping","www.qq.com")
	stdout0, err := cmd0.StdoutPipe() //输出
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd0.Start(); err != nil { //开始执行
		fmt.Println(err)
		return
	}
	useBufferIO := false
	if !useBufferIO {
		var outputBuf0 bytes.Buffer //二进制保存
		for {
			tempoutput := make([]byte, 256)
			n, err := stdout0.Read(tempoutput) //读取二进制
			if err != nil {
				if err == io.EOF {
					break //跳出循环
				} else {
					fmt.Println(err)
					return
				}
			}
			if n > 0 {
				outputBuf0.Write(tempoutput[:n]) //读取到的写入
			}
		}

		fmt.Println(outputBuf0.String())

	} else {
		outputbuf0 := bufio.NewReader(stdout0)
		output0, _, err := outputbuf0.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(output0)) //读取行

	}

}
func RunCMDPipe(cmd string, args string) {
	//cmd1:=exec.Command("tasklist")
	//cmd1:=exec.Command("ls","-l")
	cmd1 := exec.Command(cmd, args)
	var ouputbuf1 bytes.Buffer //输出
	cmd1.Stdout = &ouputbuf1   //设置输入
	if err := cmd1.Start(); err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", ouputbuf1.Bytes())

}

func RunCMDPipeThenWriteOut2File(cmd string, args string) {
	//cmd1:=exec.Command("tasklist")
	//cmd1:=exec.Command("ls","-l")
	cmd1 := exec.Command(cmd, args)
	var ouputbuf1 bytes.Buffer //输出
	cmd1.Stdout = &ouputbuf1   //设置输入
	if err := cmd1.Start(); err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", ouputbuf1.Bytes())

	filePath := "/Users/gauge/T2_22/Golang十四章经/1.al_go_parallelist/多进程编程/cmdout.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	//for i := 0; i < 5; i++ {
	//	write.WriteString("http://c.biancheng.net/golang/ \n")
	//}
	write.Write(ouputbuf1.Bytes())
	//Flush将缓存的文件真正写入到文件中
	write.Flush()

}

func RunCMDPipeThenWriteOut2FileV2(cmd string, args string) {
	//cmd1:=exec.Command("tasklist")
	//cmd1:=exec.Command("ls","-l")
	cmd1 := exec.Command(cmd, args)
	var ouputbuf1 bytes.Buffer //输出
	cmd1.Stdout = &ouputbuf1   //设置输入
	if err := cmd1.Start(); err != nil {
		fmt.Println(err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", ouputbuf1.Bytes())
	filePath := "/Users/gauge/T2_22/Golang十四章经/1.al_go_parallelist/多进程编程/cmdout2.txt"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	n, err := file.Write(ouputbuf1.Bytes()) //二进制写入文件
	if err != nil {
		fmt.Println("写入文件失败", err)
	}
	if n > 0 {
	}

	//Flush将缓存的文件真正写入到文件中
	file.Close()

}

func main1() {
	//RunCMD("ping","192.168.0.1")
	RunCMD("pwd", "  ")
	RunCMD("cd", " /Users/gauge/T2_22/Golang十四章经/1.al_go_parallelist/多进程编程")
	RunCMD("ls", "-l")
	//go RunCMD()
	// RunCMDPipe()
	RunCMDPipe("pwd", "  ")
	RunCMDPipe("cd", " /Users/gauge/T2_22/Golang十四章经/1.al_go_parallelist/多进程编程")
	RunCMDPipe("ls", "-l")
}

func main2() {
	//RunCMD("ping","192.168.0.1")
	RunCMD("ps", "-ef")

}

func main3() {
	//RunCMD("ping","192.168.0.1")
	RunCMDPipe("ps", "-ef")

}

func main() {
	//RunCMD("ping","192.168.0.1")
	//  RunCMDPipeThenWriteOut2File("ps","-ef")
	//RunCMDPipeThenWriteOut2FileV2("ifconfig","-a")
	RunCMDPipeThenWriteOut2FileV2("ps", "-ef")

}
