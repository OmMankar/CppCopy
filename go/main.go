package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const chunksize = 24

type copy struct {
	Srcf, Dstf      *os.File
	TotalSizeToCopy int
	AccData         int
	OffsetDstF      int
	SizeOfSrcF      int
	DstfStat        os.FileInfo
}

func (*copy) Progress(num int, p copy) {

	if num < 0 {
		num = 0
	}
	if num > 100 {
		num = 100
	}
	i := int(0)
	data := "["
	for i = 0; i < 100; i++ {

		if i <= num {
			data += "#"
		} else {
			data += "."
		}
	}
	data += "]"

	fmt.Printf("\r copying %d %s ", num, data)
	time.Sleep(time.Millisecond * 100)

}
func (*copy) StatOfSourceF(src *string, p copy) (int, error) {
	info, err := p.Srcf.Stat()
	if err != nil {
		return -1, err
	}
	p.SizeOfSrcF = int(info.Size())
	return p.SizeOfSrcF, nil
}
func (*copy) UpdateDstF(src *string, dst *string, p copy) {
	// Open the file already present
	var err error
	p.Dstf, err = os.OpenFile(*dst, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error occured while opening the already present destination file ", err)
		os.Exit(100)
	}
	defer p.Dstf.Close()

	//finding the size of the dst file using Stat function using which finding the offset
	p.OffsetDstF = int(p.DstfStat.Size())
	//offset varies from [0 to n-1]
	p.OffsetDstF -= 1
	p.TotalSizeToCopy = int(p.SizeOfSrcF) - int(p.DstfStat.Size())
	p.SizeOfSrcF -= p.OffsetDstF
	p.AccData = 0

	for p.SizeOfSrcF != 0 {

		var cnt int
		if p.SizeOfSrcF < chunksize {
			cnt = int(p.SizeOfSrcF)
		} else {
			cnt = chunksize
		}
		b := make([]byte, cnt)

		n, err := p.Srcf.ReadAt(b, int64(p.OffsetDstF))
		if err != nil {
			fmt.Println("error occured while reading the file ", err)
			os.Exit(100)
		}

		lenWrite, err := p.Dstf.WriteAt(b, int64(p.OffsetDstF))
		if err != nil {
			fmt.Println("error occured while writing in the file", err)
			os.Exit(100)
		}

		p.OffsetDstF += n
		p.AccData += lenWrite
		p.SizeOfSrcF -= n
		//Tracking the progess of copy
		precent := (p.AccData * 100 / int(p.TotalSizeToCopy))
		p.Progress(precent, p)
	}
}
func (*copy) CopyAtDstF(dst *string, p copy) {

	var err error
	p.Dstf, err = os.Create(*dst)
	if err != nil {
		fmt.Println("Error occured while creating new file ", err)
		os.Exit(100)
	}
	defer p.Dstf.Close()
	// writing a loop to read 25 bytes at a time from the file until eof is reached

	p.TotalSizeToCopy = int(p.SizeOfSrcF)
	p.AccData = 0
	for p.SizeOfSrcF != 0 {

		var cnt int
		if p.SizeOfSrcF < chunksize {
			cnt = int(p.SizeOfSrcF)
		} else {
			cnt = chunksize
		}

		//to track the character read from source file
		b := make([]byte, cnt)
		n, err := p.Srcf.Read(b)
		if err != nil {
			fmt.Println("error occured while reading the file ", err)
			os.Exit(100)
		}

		lenWrite, err := p.Dstf.Write(b)
		if err != nil {
			fmt.Println("error occured while writing in the file")
		}

		p.AccData += lenWrite
		p.SizeOfSrcF -= n

		//Tracking the progess of copy
		precent := (p.AccData * 100 / int(p.TotalSizeToCopy))
		p.Progress(precent, p)
	}
}
func main() {

	src := flag.String("src", "/tmp/blah.txt", "Enter the source file path")
	dst := flag.String("dst", "/tmp/blah.ttx", "Enter the destination file path")
	flag.Parse()

	//check source and destination file are same
	if *src == *dst {
		fmt.Println("Both Src and Dst are same file")
		os.Exit(100)
	}

	var p copy
	var err error
	//opening the srcf
	p.Srcf, err = os.Open(*src)
	if err != nil {
		fmt.Println("Error occured while opening the source file ", err)
		os.Exit(100)
	}
	defer p.Srcf.Close()

	//finding the stat of source file
	p.SizeOfSrcF, err = p.StatOfSourceF(src, p)
	if err != nil {
		fmt.Println("Error occured in finding Stat of Srcf", err)
		os.Exit(100)
	}

	//check wether the file already exist or not

	if p.DstfStat, err = os.Stat(*dst); err == nil {
		fmt.Println("Dst File already exists")
		p.UpdateDstF(src, dst, p)
	} else if os.IsNotExist(err) {
		fmt.Println("Creating a new File at dst path")
		p.CopyAtDstF(dst, p)
	} else {
		fmt.Println("Error checking file:", err)
		os.Exit(100)
	}
	fmt.Println()
	os.Exit(0)
}
