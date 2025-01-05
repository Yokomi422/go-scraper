package source

import (
	"fmt"
	"io"
	"net/http"
)

func Body(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		
		return bodyString, nil
	}

	return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
}