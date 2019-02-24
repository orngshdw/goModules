package utility

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// GetRequest returns []byte from page
func GetRequest(aboutendpoint string, username string, password string) ([]byte, error) {
	// https://tutorialedge.net/golang/consuming-restful-api-with-go/
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second)}
	req, err := http.NewRequest("GET", aboutendpoint, nil)
	CheckIfError(err)
	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte("request failed"), err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		msg := fmt.Sprintf("status %d, request failed", resp.StatusCode)
		return []byte(msg), err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	if s == "" {
		empty := fmt.Sprintf("%s: body text is empty", aboutendpoint)
		CheckIfError(errors.New(empty))
	}
	return bodyText, err
}
