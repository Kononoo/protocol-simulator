package internal

import (
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

// StartServer 启动服务器并接收连接
func StartServer(port int, handler func(conn net.Conn)) error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer listener.Close()

	logrus.Infof("服务器启动，监听端口: %d", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Errorf("接收连接失败: %v", err)
			continue
		}
		go handler(conn)
	}

}
