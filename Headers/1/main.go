package main

import "net/http"

func main() {

}

func getContentType(res *http.Response) string {
	return res.Header.Get("Content-Type")
}
