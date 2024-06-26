package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/eugene-fedorenko/ring"
	"github.com/valyala/fasthttp"
	"os"
	"time"
)

var (
	storageURI       = flag.String("myGaruStorageURI", "https://segments.mygaru.com/upload/", "Segment API upload endpoint")
	storageSecretKey = flag.String("myGaruSecretKey", "", "Per client secret authentication key")
	filePath         = flag.String("file", "", "CSV File for uploading")
)

func main() {

	flag.Parse()

	st := time.Now()

	if *filePath == "" {
		fmt.Println("err: csv file for uploading is not specified")
		return
	}

	file, err := os.Open(*filePath)
	if nil != err {
		fmt.Printf("error while opening file %q, err: %v", *filePath, err)
		return
	}

	fmt.Printf("reading file: %q", file.Name())
	scanner := bufio.NewScanner(file)
	var dataset [][]byte
	for i := 0; scanner.Scan(); i++ {
		uid := scanner.Bytes()
		dataset = append(dataset, uid)

	}

	fmt.Printf("\nreading done, dataset size: %d", len(dataset))
	fmt.Printf("\nconverting to the bloom filter, FPR = 0.001")

	// create bloom filter with capacity usersTotal + 20%
	bloom, _ := ring.Init(int(float64(len(dataset))*1.2), 0.001)
	for i := 0; i < len(dataset); i++ {
		bloom.Add(dataset[i])
	}

	bloomBinFormat, err := bloom.MarshalBinary()
	if nil != err {
		fmt.Printf("\nerror while marshaling bloom filter, err: %v", err)
		return
	}

	fmt.Printf("\nbloom filter size %d bytes", len(bloomBinFormat))
	fmt.Printf("\nsending data to the myGaruStorageURI: %q", *storageURI)

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("POST")
	req.SetRequestURI(fmt.Sprintf("%s?key=%s&filename=%q", *storageURI, *storageSecretKey, file.Name()))
	req.SetBody(bloomBinFormat)

	res := fasthttp.AcquireResponse()
	err = fasthttp.DoTimeout(req, res, time.Minute*10)
	if nil != err {
		fmt.Printf("\nerror while sending data to the myGaruStorageURI, err: %v", err)
		return
	}

	if res.StatusCode() != 204 {
		fmt.Printf("\nunespected response, code: %d, body: %q", res.StatusCode(), res.Body())
		return
	}

	fmt.Printf("\nsuccess! took: %s", time.Now().Sub(st))

}
