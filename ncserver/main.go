package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/vigorcrust/pure-netcopy/shared"
)

func main() {
	log.SetPrefix("[ncserver] ")
	log.Println("start...")

	server, err := net.Listen("tcp", shared.SERVERPORT)

	if err != nil {
		log.Fatalln("error listening", err)
	}

	defer server.Close()
	log.Println("server started...")
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln("error connecting", err)
		}
		go saveFile(conn)
	}
}

func saveFile(conn net.Conn) {

	bufferFileSize := make([]byte, 16)
	bufferFileName := make([]byte, 256)

	defer conn.Close()

	conn.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	conn.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	if _, err := os.Stat(fileName); err == nil {
		log.Println("File exists, skipping", fileName)
		return
	}

	log.Println("Start saving", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	var receivedBytes int64

	for {
		if (fileSize - receivedBytes) < shared.BUFFERSIZE {
			io.CopyN(file, conn, (fileSize - receivedBytes))
			conn.Read(make([]byte, (receivedBytes+shared.BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(file, conn, shared.BUFFERSIZE)
		receivedBytes += shared.BUFFERSIZE
	}
	log.Println("Received file completely!")
}
