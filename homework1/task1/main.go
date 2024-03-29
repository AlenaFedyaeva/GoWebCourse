package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func findStrInURL(subsrt string, url string) bool {
	rez := false
	b, body := readURL(url)

	if b {
		//fmt.Println(body)
		rez = find(subsrt, body)
	}
	return rez
}

// Поиск подстроки в строке (* переписать функцию на поиск в массиве)
func find(subsrt string, str string) bool {
	res := strings.Contains(str, subsrt)
	return res
}

//Чтение body в string
func readURL(url string) (bool, string) {
	var rezStr string
	var b bool
	httpCli := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := httpCli.Do(req)
	if err != nil {
		fmt.Println(err)
		return b, rezStr
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//log.Fatal(errors.New("readURL: status code not StatusOK "))
		return b, rezStr
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return b, rezStr
	}
	return true, string(bodyBytes)
}

func main() {
	var urlArray [3]string
	urlArray[0] = "https://mail.ru"
	urlArray[1] = "https://golang.org/"
	urlArray[2] = "https://google.com"

	substr := "Go is an open source programming "

	for _, url := range urlArray {
		rez := findStrInURL(substr, url)
		if rez {
			fmt.Printf("Find in page :  %s", url)
		}
	}
	fmt.Println("\nbye")
}
