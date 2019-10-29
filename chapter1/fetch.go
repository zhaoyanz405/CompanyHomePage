package main

//e1.5
import (
	"bufio"
	//"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "http://www.bing.cn"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	input := bufio.NewScanner(os.Stdout)
	for input.Scan() {
		fmt.Println("---------------------")
		fmt.Println("size: %d", b)

		fmt.Printf("%s", input.Text())
	}

}
