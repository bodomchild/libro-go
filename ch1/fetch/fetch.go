package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	switch strings.Split(os.Args[1], "-f=")[1] {
	case "1":
		fetch1()
	case "2": // Ej 1.7
		fetch2()
	case "3": // Ej 1.8
		fetch3()
	case "4": // Ej 1.9
		fetch4()
	default:
		fmt.Println("Incorrect value for argument -f")
	}
}

func fetch1() {
	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v'n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func fetch2() {
	for _, url := range os.Args[2:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v'n", url, err)
			os.Exit(1)
		}
	}
}

func fetch3() {
	for _, url := range os.Args[2:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v'n", url, err)
			os.Exit(1)
		}
	}
}

func fetch4() {
	for _, url := range os.Args[2:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v'n", url, err)
			os.Exit(1)
		}
		fmt.Println("Status:", resp.Status)
	}
}
