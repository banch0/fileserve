package main

import (
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

// BufferSize ...
const BufferSize = 4096

func main() {
	connection, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(connection)
	}
	defer connection.Close()

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)

	connection.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	connection.Read(bufferFileName)
	fileName := strings.Trim(string(bufferFileName), ":")

	newFile, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	var receivedBytes int64

	for {
		if (fileSize - receivedBytes) < BufferSize {
			io.CopyN(newFile, connection, (fileSize - receivedBytes))
			connection.Read(make([]byte, (receivedBytes+BufferSize)-fileSize))
			break
		}
		io.CopyN(newFile, connection, BufferSize)
		receivedBytes += BufferSize
	}
}
