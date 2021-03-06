package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func sender(conn net.Conn, infohash string, filename string) {
	buffer := make([]byte, 2048)
	words := "get_peers "
	words += infohash
	conn.Write([]byte(words))
	fmt.Println("send over")
	n, _ := conn.Read(buffer)
	fmt.Println(string(buffer[:n]))
	str_list := strings.Split(string(buffer[:n]), "_")
	fmt.Println(len(str_list))
	strList := str_list[0 : len(str_list)-1]
	fmt.Println(strList)
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Intn(len(strList))
	// 随机选择一个节点下载文件
	download(strList[index], filename)

}
func download(addr string, filename string) {

	buffer := make([]byte, 4096)
	tcpaddr, _ := net.ResolveTCPAddr("tcp", addr)
	conn, _ := net.DialTCP("tcp", nil, tcpaddr)
	f, _ := os.Create(filename)
	defer f.Close()
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		f.Write(buffer[0:n])
	}
}
func showinfo(conn net.Conn) {
	buffer := make([]byte, 2048)
	words := "get_info"
	conn.Write([]byte(words))
	fmt.Println("send ok")
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	str := string(buffer[0:n])
	str_list := strings.Split(str, "_")
	// 123_test.sh_345_kkkk.sh_
	for i := 0; i < len(str_list)-1; i += 2 {
		fmt.Printf("INFO:%s filename: %s \n", str_list[i], str_list[i+1])
	}
}
func upload(conn net.Conn, path string) {
	word := "openTcp " + path
	conn.Write([]byte(word))
	fmt.Println("test")
}

func main() {
	server := "127.0.0.1:2333"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println("connect success")
	fmt.Println("please enter your choice, /q to quit, /help to show help info")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>>")
		choice := ""
		strBytes, _, err := reader.ReadLine()
		choice = string(strBytes)
		if err != nil {
			fmt.Println(err)
			break
		}
		if strings.HasPrefix(choice, "/q") {
			break
		} else if strings.HasPrefix(choice, "/help") {
			fmt.Println("/q to quit this program")
			fmt.Println("/help to show this info")
			fmt.Println("/showinfo to show this machine has recved infohash")
			fmt.Println("/download infohash filename to download this infohash's file to file")
			fmt.Println("/upload filename : upload a file to DHT network")
		} else if strings.HasPrefix(choice, "/showinfo") {
			showinfo(conn)
		} else if strings.HasPrefix(choice, "/download") {
			fmt.Println(choice)
			list := strings.Split(choice, " ")
			infohash := list[1]
			filename := list[2]
			sender(conn, infohash, filename)
		} else if strings.HasPrefix(choice, "/upload") {
			path := strings.Split(choice, " ")[1]
			upload(conn, path)
		}
	}

}
