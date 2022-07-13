package main

import (
        "bufio"
	"fmt"
	"log"
	"net/http"
        "os"
	"strings"
)

func GoNgl(Url string, teks string) {
	req, err := http.Post(Url, "application/json", strings.NewReader(`{"question": " `+ teks +`"}`))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	if err != nil || req.StatusCode != 200 {
		log.Println("StatusCode: ", req.StatusCode, " || StatusText: Gagal")
		return
	}
	defer req.Body.Close()
	log.Println("StatusCode: ", req.StatusCode, " || StatusText: Berhasil")
}

func init() {
	fmt.Println(`
██████╗  ██████╗               ███╗   ██╗ ██████╗ ██╗     
██╔════╝ ██╔═══██╗              ████╗  ██║██╔════╝ ██║     
██║  ███╗██║   ██║    █████╗    ██╔██╗ ██║██║  ███╗██║     
██║   ██║██║   ██║    ╚════╝    ██║╚██╗██║██║   ██║██║     
╚██████╔╝╚██████╔╝              ██║ ╚████║╚██████╔╝███████╗
╚═════╝  ╚═════╝               ╚═╝  ╚═══╝ ╚═════╝ ╚══════╝`)
	fmt.Printf("Create By : %sVnia%s - %sgithub.com/fckvania%s\n", "\033[31m", "\033[0m", "\033[34m", "\033[0m")
}

func main() {
	var (
		url   string
		teks  string
		total int
	)
	fmt.Print("[+] Please Enter Url: ")
	if _, err := fmt.Scanln(&url); err != nil {
		panic("Please Masukan Url")
	}
	fmt.Print("[+] Please Enter Teks: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		teks = string(scanner.Text())
	}
        fmt.Print("[+] Total Spam: ")
	if _, err := fmt.Scanln(&total); err != nil {
		panic(err)
	}
	for i := 0; i < total; i++ {
		GoNgl(url, teks)
	}
}
