package main

import (
	"fmt"
	"io"
	"os"
)

type CountWriter struct {
	Writer io.Writer
	Count  int
}

func (cw *CountWriter) Write(content []byte) (int, error) {
	n, err := cw.Writer.Write(content)
	//fmt.Println(" wirter: ", cw.Writer)
	if err != nil {
		return n, err
	}
	cw.Count += n
	return n, nil
}

func CountingWriter(writer io.Writer) (io.Writer, *int) {
	cw := CountWriter{
		Writer: writer,
	}
	return &cw, &(cw.Count)
}

func main() {
	f, err := os.OpenFile("/home/alan/go/src/github.com/yufeifly/gopl.io/ch7/ex7-2/out", os.O_WRONLY, 0)
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	cw, _ := CountingWriter(f)
	fmt.Fprintf(cw, "%s", "fool")
	cw.Write([]byte("def"))
	//_, err = f.Write([]byte("hello,yufeifly"))
	//if err != nil {
	//	fmt.Printf("write file err: %v\n", err)
	//}
}
