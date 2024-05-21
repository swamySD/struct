package main 

import (
	"fmt"
	"io"
	"compress/gzip"
	"compress/zlib"
	"bytes"
)

type compressor interface{
	Compress(data []bytes) ([]bytes,error)
	Decompress(data []bytes) ([]bytes,error)
}

type GzipComprssor struct{}

func main(){
	fmt.Println("compressor")

	
}

func gzipcompress(data byte[])([]byte,error){
	



}