package Hough_Transform

import (
	"crypto/md5"
	"io/ioutil"
	"os"
	"testing"
)

func TestHough(t *testing.T) {
	if h := Hough("Hough.png", "test.png", 460, 360); h != true {
		t.Fatal(h)
	}
	s, err := ioutil.ReadFile("test.png"); if err != nil {
		t.Fatal(err)
	}
	w, err := ioutil.ReadFile("example.png"); if err != nil {
		t.Fatal(err)
	}
	m, err := ioutil.ReadFile("Hough.png"); if err != nil {
		t.Fatal(err)
	}
	if md5.Sum(s) != md5.Sum(w) {
		t.Fatal()
	}
	if md5.Sum(s) == md5.Sum(m) {
		t.Fatal()
	}
	err = os.Remove("test.png")
	if err != nil {
		t.Fatal(err)
	}
}
