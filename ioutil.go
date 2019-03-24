package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	READY = "READY\n"
	OK    = "RESULT 2\nOK"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
)

func initIOBuffers() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
}

func emptyFile(file *os.File) (bool, error) {
	fs, err := file.Stat()
	if err != nil {
		return false, err
	}
	return fs.Size() == 0, nil
}

func readHeaderData() (string, error) {
	if ok, err := emptyFile(os.Stdin); ok || err != nil {
		return "", err
	}
	return readHeaderDataFromReader(in)
}

func readEventData(len int64) (string, error) {
	if ok, err := emptyFile(os.Stdin); ok || err != nil {
		return "", err
	}
	return readEventDataFromReader(in, len)
}

func replyOk() {
	out.WriteString(OK)
}

func replyReady() {
	out.WriteString(READY)
}

func readHeaderDataFromReader(in *bufio.Reader) (string, error) {
	var headerbuilder strings.Builder

	for {
		bytes, isPrefix, err := in.ReadLine()
		if err != nil {
			return "", fmt.Errorf("Could not read event header")
		}
		headerbuilder.Write(bytes)
		if !isPrefix {
			break
		}
	}
	return headerbuilder.String(), nil
}

func readEventDataFromReader(in *bufio.Reader, len int64) (string, error) {
	bodybytearray := make([]byte, len)
	var bodystringbuilder strings.Builder
	if _, err := io.ReadFull(in, bodybytearray); err != nil {
		return "", fmt.Errorf("Could not read event body")
	}
	bodystringbuilder.Write(bodybytearray)
	return bodystringbuilder.String(), nil
}
