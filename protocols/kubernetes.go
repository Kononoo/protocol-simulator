package protocols

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"protocol-simulator/internal"
)

// Kubernetes Param 配置参数
var (
	Kubernetes_SessionTime int    = 3000
	Kubernetes_Version     string = "v1.20"
)

func StartKubernetesServer(port int) error {
	return internal.StartServer(port, kubernetesHandler)
}

func kubernetesHandler(conn net.Conn) {
	defer conn.Close()
	logrus.Infof("[Kubernetes] 接收连接来自 %v", conn.RemoteAddr())

	// 模拟Kubernetes协议的握手过程
	handshake := fmt.Sprintf("Kubernetes Version: %s\n", Kubernetes_Version)
	_, err := conn.Write([]byte(handshake))
	if err != nil {
		logrus.Errorf("[Kubernetes] 发送握手信息失败: %v", err)
		return
	}
	logrus.Infof("[Kubernetes] 发送握手信息: %v", handshake)

	// 模拟Kubernetes协议交互过程
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			logrus.Errorf("[Kubernetes] 读取客户端信息失败: %v", err)
			break
		}
		logrus.Infof("[Kubernetes] 接收到客户端信息: %s", message)

		response := fmt.Sprintf("Echo: %s", message)
		_, err = conn.Write([]byte(response))
		if err != nil {
			logrus.Errorf("[Kubernetes] 发送响应信息失败: %v", err)
			break
		}
		logrus.Infof("[Kubernetes] 发送响应信息: %s", response)
	}
}

/*================================================================================*/
/*
Protocol: Kubernetes

Kubernetes协议是一个用于管理容器化应用的开源系统，用于自动化部署、扩展和管理应用程序。该模拟程序实现了基本的握手和回显交互过程。

备注：该模拟程序实现了v1.20版本的Kubernetes协议。

版本号: v1.20

params:
1. Kubernetes_SessionTime: 交互时间（毫秒）
2. Kubernetes_Version: 协议版本

Process:
s1: 客户端连接到服务器
s2: 服务器发送握手信息（包含协议版本）
s3: 客户端发送消息到服务器
s4: 服务器回显客户端消息

协议的数据结构：

|-----------------------------------------------------------------------------------
| Header ----------------------------------------------------------------------------
| Payload ----------------------------------------------------------------------------
| Type ------------- 4 bytes
| Version ---------- 6 bytes
| -----------------------------------------------------------------------------------
*/
/*================================================================================*/
