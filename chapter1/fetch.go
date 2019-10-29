package main

//e1.5
import (
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
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	//	os.Exit(1)
	//}
}
