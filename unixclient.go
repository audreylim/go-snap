package main

import (
	"io"
	"log"
	"net"
	"os"
	"syscall"
)

func main() {
	addr := "/tmp/unixdomaincli"
	a := "unix"

	if err := syscall.Unlink(addr); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	laddr := net.UnixAddr{addr, a}
	conn, err := net.DialUnix(a, &laddr, &net.UnixAddr{"/tmp/unixdomain", a})
	if err != nil {
		panic(err)
	}
	defer os.Remove(addr)

	out, err := os.Create("unixpic.jpg")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(out, conn)
	if err != nil {
		log.Fatal(err)
	}
}
