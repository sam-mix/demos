package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l") // 创建一个执行 "ls -l" 命令的子进程
	err := cmd.Start()              // 启动子进程
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process ID: %d\n", cmd.Process.Pid) // 打印子进程的进程 ID

	err = cmd.Wait() // 等待子进程结束
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Command finished.")
}
