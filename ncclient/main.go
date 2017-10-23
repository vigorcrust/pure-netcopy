package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/vigorcrust/pure-netcopy/shared"
)

func main() {
	log.SetPrefix("[ncclient] ")
	log.Println("start...")

	for _, match := range os.Args[1:] {
		_, err := os.Stat(match)
		if err != nil {
			log.Fatalln("Error existing:", err)
		}
		processFile(match)
	}

	log.Println("...end")
}

func processFile(singleFileName string) {
	start := time.Now()
	log.Println("processing:", singleFileName)
	conn, err := net.Dial("tcp", "192.168.0.108" + shared.SERVERPORT)
	if err != nil {
		log.Println("Error dialing:", err)
		return
	}

	defer conn.Close()
	log.Println("Connected to server, start sending file")

	file, err := os.Open(singleFileName)
	if err != nil {
		log.Println("Error opening:", err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Error fileinfo:", err)
		os.Exit(1)
	}

	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 16)
	fileName := fillString(fileInfo.Name(), 256)

	log.Println("Sending filename and filesize!")
	conn.Write([]byte(fileSize))
	conn.Write([]byte(fileName))
	sendBuffer := make([]byte, shared.BUFFERSIZE)
	log.Println("Start sending file!")
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		conn.Write(sendBuffer)
	}
	log.Println("File has been sent, closing connection!")

	elapsed := time.Since(start)
	bytesPerSecond := float64(fileInfo.Size()) / elapsed.Seconds() / 1024

	log.Println("Transfered", fileInfo.Size(), "bytes, within", elapsed, "-", bytesPerSecond, "KB/s")

	return
}

func fillString(returnString string, toLength int) string {
	for {
		lengtString := len(returnString)
		if lengtString < toLength {
			returnString = returnString + ":"
			continue
		}
		break
	}
	return returnString
}
