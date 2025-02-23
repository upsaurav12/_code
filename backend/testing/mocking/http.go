package main

import "io"

func FetchData(client mockHTTPClient, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
