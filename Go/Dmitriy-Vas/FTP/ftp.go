package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

type Data struct {
	Host      string
	Username  string
	Password  string
	Directory string
	File      string
}

func (d *Data) Send() {

	conn, err := ftp.DialWithOptions(d.Host, ftp.DialWithTimeout(time.Second*10))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Quit()
	fmt.Println(conn)

	if err = conn.Login(d.Username, d.Password); err != nil {
		log.Fatal(err)
	}
	// Uncomment this to change directory
	/*if err = conn.ChangeDir(d.Directory); err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn.CurrentDir())*/
	files, err := conn.List(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Printf("%v %12d %v %v\n", f.Time, f.Size, f.Type, f.Name)
	}

	r, err := conn.Retr(d.File)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	f, err := os.Create(d.File)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, err := io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wrote", n, "bytes to", d.File)
}
