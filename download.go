// download.go
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if flag.NFlag() < 1 {
		fmt.Println("Download target URL from Easynews and save as Filename")
		fmt.Println("Use -url= and -filename=")

		return
	}
	urlPtr := flag.String("url", "", "Source URL")
	filenamePtr := flag.String("filename", "downloaded", "Target filename")
	flag.Parse()

	url := urlPtr
	filename := filenamePtr

	FetchFromEasynews(*url, *filename)
}

func FetchFromEasynews(url string, filename string) bool {
	username := "<username>"
	password := "<password>"

	client := &http.Client{}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return true
}
