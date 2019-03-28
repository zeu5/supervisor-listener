package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	ready = "READY\n"
	ok    = "RESULT 2\nOK"
)

var (
	in  *bufio.Reader
	out *bufio.Writer
	err *bufio.Writer
)

func initLogger(verbose bool) {
	if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}

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
	return readHeaderDataFromReader(in)
}

func readEventData(len int64) (string, error) {
	return readEventDataFromReader(in, len)
}

func replyOk() {
	out.WriteString(ok)
	out.Flush()
}

func replyReady() {
	out.WriteString(ready)
	out.Flush()
}

func readHeaderDataFromReader(in *bufio.Reader) (string, error) {
	s, err := in.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Could not read header form input")
	}
	return strings.TrimSuffix(s, "\n"), nil
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
