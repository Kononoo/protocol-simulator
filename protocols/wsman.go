package protocols

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"protocol-simulator/internal"
)

// WSMAN_Params 配置参数
var (
	WSMAN_SessionTime int    = 5000
	WSMAN_Version     string = "v2.3"
)

func StartWSMANServer(port int) error {
	return internal.StartServer(port, wsmanHandler)
}

func wsmanHandler(conn net.Conn) {
	defer conn.Close()
	logrus.Infof("[WSMAN] 接受连接来自 %v", conn.RemoteAddr())

	// 模拟WSMAN协议的握手过程
	handshake := fmt.Sprintf("WSMAN Version: %s", WSMAN_Version)
	_, err := conn.Write([]byte(handshake))
	if err != nil {
		logrus.Errorf("[WSMAN] 发送握手信息失败: %v", err)
		return
	}
	logrus.Infof("[WSMAN] 发送握手信息: %s", handshake)

	// 模拟WSMAN协议的交互过程
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			logrus.Errorf("[WSMAN] 读取客户端信息失败: %v", err)
			break
		}
		logrus.Infof("[WSMAN] 接收到客户端信息: %s", message)

		response := fmt.Sprintf("Echo: %s", message)
		_, err = conn.Write([]byte(response))
		if err != nil {
			logrus.Errorf("[WSMAN] 发送响应信息失败: %v", err)
			break
		}
		logrus.Infof("[WSMAN] 发送响应信息: %s", response)
	}
}

/*================================================================================*/
/*
Protocol: WSMAN

WSMAN协议是Web服务管理协议，主要用于远程管理和监控系统。该模拟程序实现了基本的握手和回显交互过程。

备注：该模拟程序实现了v2.3版本的WSMAN协议。

版本号: v2.3

params:
1. WSMAN_SessionTime: 交互时间（毫秒）
2. WSMAN_Version: 协议版本

Process:
s1: 客户端连接到服务器
s2: 服务器发送握手信息（包含协议版本）
s3: 客户端发送消息到服务器
s4: 服务器回显客户端消息

格式复杂的协议标记出协议的数据结构：

|-----------------------------------------------------------------------------------
| Header ----------------------------------------------------------------------------
| Payload ----------------------------------------------------------------------------
| Type ------------- 4 bytes
| Version ---------- 6 bytes
| -----------------------------------------------------------------------------------
*/
/*================================================================================*/
