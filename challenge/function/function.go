// Q: Write a function that gets a URL and returns the value of Content-Type response HTTP header.
// The function should return an error if it can't perform a GET request.
// The signature of the function is contentType, it gets a URL as a string, and returns a string and an error.
// Use net HTP package GET function to make an HTP call. Use the resp.Header.Get method to get the value of a header.
// And make sure the response body is closed properly.

package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	ct := response.Header.Get("Content-Type")
	if ct == "" {
		return "", fmt.Errorf("Can't find content-type header")
	}

	return ct, nil
}

func main() {
	ct, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(ct)
	}
}
