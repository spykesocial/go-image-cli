/*
	Copyright © 2022 Spyke Social Private Limited.

*/
package image

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path"
)

/*
	Reference: https://golangbyexample.com/download-image-file-url-golang/
*/

func getFile(URL string) (http.Response, error) {
	// get file through http from URL
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	// defer res.Body.Close()

	// if there is an http error
	if res.StatusCode != 200 {
		return nil, errors.New("Error downloading file. Received statusCode:", res.StatusCode)
	}

	return res, nil
}

func SaveFile(URL, filename string) (err error) {
	res, err := getFile(URL)
	defer res.Body.Close()
	if err != nil {
		return
	}

	fname := path.Base(filename)

	// create an empty file
	file, err := os.Create(fname)
	defer file.Close()
	if err != nil {
		return
	}

	// copy the bytes into the file
	_, err := io.Copy(file, res.Body)
	if err != nil {
		return
	}

	return nil
}
