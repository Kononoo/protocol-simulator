package cmd

import (
	"flag"
	"github.com/sirupsen/logrus"
	"protocol-simulator/internal"
)

func main() {
	// 定义命令行参数
	var (
		port     int
		logLevel string
		protocol string
	)

	flag.IntVar(&port, "p", 8080, "服务监听端口")
	flag.StringVar(&logLevel, "l", "info", "日志级别")
	flag.StringVar(&protocol, "protocol", "kubernetes", "模拟的协议名称")

	// 初始化日志
	internal.InitLog(logLevel)

	// 启动服务器
	var err error
	switch protocol {
	case "kubernetes":
		err = protocols.StartKubernetesServer(port)
	case "wsman":
		err = protocols.StartWSMANServer(port)
	default:
		logrus.Fatalf("不支持的协议: %s", protocol)
	}

	if err != nil {
		logrus.Fatalf("启动服务器失败: %v", err)
	}
}
