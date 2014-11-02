package main

import (
	"io/ioutil"
	"os"
)

func main() {
	read, _ := ioutil.ReadFile("helpers/raspipic.jpg")
	os.Stdout.Write(read)
}
