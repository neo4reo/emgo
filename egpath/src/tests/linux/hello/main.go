package main

import (
	"fmt"
	"os"
	"sync"
	"syscall"
)

func checkErr(err error) {
	if err != nil {
		os.Stderr.WriteString("Fatal error: ")
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString(".\n")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Args:")
	for i, a := range os.Args {
		fmt.Printf("%d: %s\n", i, a)
	}

	fmt.Println("Env:")
	for i, e := range os.Env {
		fmt.Printf("%d: %s\n", i, e)
	}
	fmt.Printf("EGPATH = '%s'\n", os.Getenv("EGPATH"))

	/*
		buf := make([]byte, 80)
		n := copy(buf, os.Args[0])
		fmt.Printf("Program name: %s\n", buf[:n])

		n, err := os.Stdin.Read(buf)
		checkErr(err)

		f, err := os.OpenFile(
			"file.txt",
			os.O_CREATE|os.O_RDWR|os.O_EXCL,
			0660,
		)
		checkErr(err)

		n, err = f.Write(buf[:n])
		checkErr(err)

		f.Close()
		fmt.Printf("%d bytes written.\n", n)
	*/

	sd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	checkErr(err)
	err = syscall.SetsockoptInt(sd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, 524288)
	checkErr(err)

	sa := syscall.RawSockaddrInet4{
		Family: syscall.AF_INET,
		Port:   syscall.Htons(1234),
	}
	err = syscall.Bind(sd, &sa)
	checkErr(err)

	con := os.NewFile(uintptr(sd), "")

	var buf [2048]byte
	var m sync.Mutex
	for {
		m.Lock()
		n, err := con.Read(buf[:])
		checkErr(err)
		var tp syscall.Timespec
		checkErr(syscall.ClockGettime(syscall.CLOCK_MONOTONIC, &tp))

		fmt.Printf("%d.%09d pkt %d B\n", tp.Sec, tp.Nsec, n)
		m.Unlock()
	}

}
