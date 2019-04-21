package main

import (
	"os"
	"testing"
)

func TestData_Send(t *testing.T) {
	d := &Data{
		Host:      "speedtest.tele2.net:21",
		Username:  "Anonymous",
		Password:  "Anonymous",
		Directory: "",
		File:      "512KB.zip",
	}
	d.Send()

	err := os.Remove(d.File)
	if err != nil {
		t.Fatal(err)
	}
}
