package main

import (
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// BufferSize ...
const BufferSize = 4096

func main() {
	addr := "0.0.0.0:9000"
	log.Println(os.Getenv("PORT"))
	log.Println("server starting")
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("server can't started error: %s", err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Println("Client connected!")

	file, err := os.Open("image.jpg")
	if err != nil {
		log.Println(err)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err)
		return
	}

	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)
	fileName := fillString(fileInfo.Name(), 64)
	log.Println("Sending filename and filesize!")
	conn.Write([]byte(fileSize))
	conn.Write([]byte(fileName))
	sendBuffer := make([]byte, BufferSize)
	for {
		_, err := file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		conn.Write(sendBuffer)
	}
	return
}

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}
