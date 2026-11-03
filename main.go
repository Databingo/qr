package main

import (
	"fmt"
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run main.go <文件路径>")
		fmt.Println("示例: go run main.go ~/.ssh/id_ed25519.pub")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读取文件失败: %v\n", err)
		os.Exit(1)
	}

	// 打印文件内容对应的终端二维码
	qrterminal.Generate(string(data), qrterminal.L, os.Stdout)

	// !! 注意安全清理:
	// tmux 下使用: ctrl-b :clear-history
	// 终端重置: reset
}
