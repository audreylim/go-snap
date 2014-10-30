package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"syscall"
)

func cameraStill(w io.Writer, flip string) {
	var args []string

	if flip != "" {
		args = append(args, "-"+flip, "-o", "-")
	} else {
		args = append(args, "-o", "-")
	}

	cmd := exec.Command("raspistill", args...)

	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(w, out)
	if err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}

func main() {
	//listen on Unix Socket
	go func() {
		addr := "/tmp/unixdomain"
		if err := syscall.Unlink(addr); err != nil && !os.IsNotExist(err) {
			panic(err)
		}
		ln, err := net.ListenUnix("unix", &net.UnixAddr{addr, "unix"})
		if err != nil {
			fmt.Println(err)
		}
		defer os.Remove(addr)
		for {
			conn, err := ln.AcceptUnix()
			if err != nil {
				fmt.Println(err)
				continue
			}
			go cameraStill(conn, "")
		}
	}()

	//listen on HTTP
	http.HandleFunc("/snap", func(w http.ResponseWriter, r *http.Request) {
		f := r.URL.Query().Get("flip")
		w.WriteHeader(http.StatusOK)
		cameraStill(w, f)
	})
	http.ListenAndServe(":3000", nil)
}
