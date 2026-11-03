//package main
//
//import (
//	"fmt"
//	"os"
//
//	"github.com/mdp/qrterminal/v3"
//)
//
//func main() {
//	if len(os.Args) < 2 {
//		fmt.Println("用法: go run main.go <文件路径>")
//		fmt.Println("示例: go run main.go ~/.ssh/id_ed25519.pub")
//		os.Exit(1)
//	}
//
//	filePath := os.Args[1]
//	data, err := os.ReadFile(filePath)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "读取文件失败: %v\n", err)
//		os.Exit(1)
//	}
//
//	// 打印文件内容对应的终端二维码
//	qrterminal.Generate(string(data), qrterminal.L, os.Stdout)
//
//	// !! 注意安全清理:
//	// tmux 下使用: ctrl-b :clear-history
//	// 终端重置: reset
//}
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

	// 保持第一版原汁原味的纯黑白背景块，只改 QuietZone 边距让它变小
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK, // 强制纯黑背景色块（绝不发绿）
		WhiteChar: qrterminal.WHITE, // 强制纯白背景色块
		QuietZone: 1,                // 只改这里：白边从默认的巨大外圈缩到最小，整体立刻变小！
	}

	qrterminal.GenerateWithConfig(string(data), config)
}

