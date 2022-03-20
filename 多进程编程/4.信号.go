package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

func getCmdText(cmd *exec.Cmd) string {
	var buf bytes.Buffer
	buf.WriteString(cmd.Path) //写入路径
	for _, arg := range cmd.Args[1:] {
		buf.WriteRune(' ')
		buf.WriteString(arg)
	}
	return buf.String() //返回字符串
}

func RunCmds(cmds []*exec.Cmd) ([]string, error) {
	if cmds == nil || len(cmds) == 0 {
		return nil, errors.New("cmds为空")
	}
	first := true
	var output []byte
	for _, cmd := range cmds {
		fmt.Println("★ ☎☎☎☎☎ run Command", getCmdText(cmd)) //☏　☎　✆
		if !first {
			var stdinBuf bytes.Buffer
			stdinBuf.Write(output) //写入
			cmd.Stdin = &stdinBuf  //设置输入
		}
		var stdoutbuf bytes.Buffer
		cmd.Stdout = &stdoutbuf             //设定输入
		if err := cmd.Start(); err != nil { //开启命令，
			return nil, err
		}
		if err := cmd.Wait(); err != nil {
			return nil, err
		}
		output = stdoutbuf.Bytes() //写入
		//第一个命令的记过当作二次命令的写入
		if first {
			//first=false
		}

	}
	var lines []string
	var outputbuf bytes.Buffer
	outputbuf.Write(output) //写入
	for {
		line, err := outputbuf.ReadBytes('\n')
		fmt.Println("line:", string(line))
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		lines = append(lines, string(line)) //读取
	}
	return lines, nil
}

//处理进程编号
func GetPids(strs []string) ([]int, error) {
	var pids []int
	for _, str := range strs {
		fmt.Println("getpid", str)
		pid, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, err
		}
		pids = append(pids, pid)
	}
	return pids, nil
}

//处理进程编号
func GetPidByName(strs []string, name string) ([]int, error) {
	var pids []int
	for _, str := range strs {
		fmt.Println("getpid", str)
		pid := 88888
		if strings.Contains(str, name) {
			//找出 pid
			fmt.Println("☎ ☎ ☎ ☎ ☎ 找到可疑进程：", str)
			field := strings.Fields(str)
			pidstr := field[1]
			pid, _ = strconv.Atoi(pidstr)
		}
		if pid != 88888 {
			pids = append(pids, pid)
		}
	}
	return pids, nil
}

func SendSingal() {
	cmds := []*exec.Cmd{
		//exec.Command("tasklist",),
		//exec.Command("ipconfig",),
		exec.Command("ps", "-ef"),
		//exec.Command("ifconfig",),
	}
	output, err := RunCmds(cmds) //执行命令
	if err != nil {
		fmt.Println(err)
		return
	}
	//pids,err:=GetPids(output)//获取进程编号
	pids, err := GetPidByName(output, "印象") //获取进程编号
	fmt.Println("☎ ☎ ☎ ☎ ☎ 找到的可疑进程：", pids)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, pid := range pids {
		proc, err := os.FindProcess(pid) //抓取进程
		if err != nil {
			fmt.Println(err)
			//sig:=syscall.SIGQUIT//退出
			//fmt.Println("发送中断信号",sig,pid)
			//proc.Signal(sig)//发送中断信号
			return
		}
		sig := syscall.SIGQUIT //退出
		fmt.Println("发送信号", sig, pid)
		err = proc.Signal(sig) //发送信号
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//tasklist,ipconfig
func HandleSingal() {

	///☎☎☎☎☎  接收者进程使用了两个信道（os.Signal）,并使用 signal.Notify(sigRecv1, sigs1...)  发送 操作系统信号，
	//处理接收信号
	sigRecv1 := make(chan os.Signal, 1)                   //创建管道
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT} //信号
	fmt.Println("sig1", sigs1)                            //接收
	signal.Notify(sigRecv1, sigs1...)                     //通知

	sigRecv2 := make(chan os.Signal, 1)
	sigs2 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Println("sig2", sigs2)
	signal.Notify(sigRecv2, sigs2...)

	///☎☎☎☎☎ 然后操作两个协程 等待信道接收消息
	var wg sync.WaitGroup
	wg.Add(2) //等待两个
	go func() {
		for sig := range sigRecv1 {
			fmt.Println("☎☎☎☎☎ Rece1", sig)
		}
		fmt.Println("☎☎☎☎☎ Rece1", "over")
		wg.Done()
	}()
	go func() {
		for sig := range sigRecv2 {
			fmt.Println("☎☎☎☎☎ Rece2", sig)
		}
		fmt.Println("☎☎☎☎☎ Rece2", "over")
		wg.Done()
	}()

	///最后在 25 秒 和 26 秒后close(sigRecv1) //关闭管道关闭两个信道
	fmt.Println("wait for 25 seconds")
	time.Sleep(time.Second * 25)
	fmt.Println("stop  Notify")
	signal.Stop(sigRecv1)
	close(sigRecv1) //关闭管道
	fmt.Println("sigRecv1 over")

	fmt.Println("wait for 26 seconds")
	time.Sleep(time.Second * 26)
	fmt.Println("stop  Notify")
	signal.Stop(sigRecv2)
	close(sigRecv2) //关闭管道
	fmt.Println("sigRecv2  over")

	wg.Wait()
}

func main() {
	//发送信号
	go func() {
		time.Sleep(time.Second * 5)
		SendSingal()
	}()
	//接收信号
	HandleSingal()
}
