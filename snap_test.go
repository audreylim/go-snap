package main

import (
	"bytes"
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

	//Where flip param is valid.
	output := bytes.NewBuffer(nil)
	cameraStill(output, "")

	if !bytes.Equal(output.Bytes(), imageBytes) {
		t.Errorf("Test failed when flip param is empty")
	}

	//Where flip param is invalid.
	output = bytes.NewBuffer(nil)
	cameraStill(output, "asdf")

	if bytes.Equal(output.Bytes(), imageBytes) {
		t.Errorf("Test failed when flip param is invalid.")
	}
}
