package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	// 用于发送数据的管道
	C    chan string
	Name string
	Addr string
}

// 保存在线用户
var onlineMap map[string]Client

// 用户发送的消息
var message = make(chan string)

func main() {

	// 监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	onlineMap = make(map[string]Client)

	// 转发消息，只要有消息来了，就遍历 map 给每个成员发送消息
	go Manager()

	// 主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		// 处理当前用户连接
		go HandleConn(conn)
	}

	defer listener.Close()
}

/**
  转发消息，只要有消息来了，就遍历 map 给每个成员发送消息
*/
func Manager() {
	for {
		// 如果没有消息，就会阻塞
		msg := <-message
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

/**
  处理当前用户连接
*/
func HandleConn(conn net.Conn) {
	cliAddr := conn.RemoteAddr().String()
	// 为当前的连接创建结构体
	client := Client{make(chan string), cliAddr, cliAddr}
	onlineMap[cliAddr] = client

	// 新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(client, conn)
	// 当前客户端发送上线广播
	message <- MakeMsg(client, "上线了")

	// 当前客户端是否主动退出
	isQuit := make(chan bool)
	// 超时
	hasData := make(chan bool)
	// 新建一个协程，接收用户发送过来的消息
	go func() {
		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// 对方断开或者出问题
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:n-2])
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				client.Name = name
				onlineMap[cliAddr] = client
				conn.Write([]byte("修改名成功\n"))
			} else {
				// 转发此内容
				message <- MakeMsg(client, msg)
			}
			// 有数据
			hasData <- true
		}
	}()

	// 当前客户端用不断线
	for {
		select {
		case <-isQuit:
			// 移除当前用户
			delete(onlineMap, cliAddr)
			// 广播下线
			message <- MakeMsg(client, "login out")
			return
		case <-hasData:
		case <-time.After(10 * time.Second):
			delete(onlineMap, cliAddr)
			message <- MakeMsg(client, "time out leave out")
			conn.Close()
		}
	}

	defer conn.Close()
}

/**
新开一个协程，专门给当前客户端发送信息
*/
func WriteMsgToClient(client Client, conn net.Conn) {
	for msg := range client.C {
		// 给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

/**
工具函数
*/
func MakeMsg(client Client, msg string) string {
	return "[" + client.Addr + "]" + client.Name + ": " + msg
}
