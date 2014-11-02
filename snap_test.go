package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func init() {
	Path = os.Getenv("PATH_TO_RASPISTILL-MOCK")
}

func TestCameraStill(t *testing.T) {
	imageBytes, err := ioutil.ReadFile("helpers/raspipic.jpg")
	if err != nil {
		t.Error(err)
	}

	//where param is valid
	flip := ""
	output := bytes.NewBuffer(nil)
	cameraStill(output, flip)

	if bytes.Equal(output.Bytes(), imageBytes) {
		fmt.Println("test1 passed")
	} else {
		t.Errorf("test1 failed")
	}

	//where param is invalid
	flip = "asdf"
	output = bytes.NewBuffer(nil)
	cameraStill(output, flip)

	if !bytes.Equal(output.Bytes(), imageBytes) {
		fmt.Println("test2 passed")
	} else {
		t.Errorf("test2 failed")
	}

}
