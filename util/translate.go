package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

const UNKNOWN_ERROR = "An error occured, please try again"

var goClient = &http.Client{Transport: &http.Transport{
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
	IdleConnTimeout:     90 * time.Second,
}, Timeout: 120 * time.Second}

type ErrorResponse struct {
	Detail string `json:"detail"`
}

func SendImageDataToMl(
	mlServerAddress string,
	imageData []byte,
	fileName string,
) ([]byte, string, error) {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", fileName) // Replace with actual filename
	if err != nil {
		err = errors.New(UNKNOWN_ERROR)
		return nil, "0", err
	}
	_, err = part.Write(imageData)
	if err != nil {
		err = errors.New(UNKNOWN_ERROR)
		return nil, "0", err
	}

	err = writer.Close()
	if err != nil {
		err = errors.New(UNKNOWN_ERROR)
		return nil, "0", err
	}
	goRequest, err := http.NewRequest(
		http.MethodPost,
		mlServerAddress,
		buffer,
	)
	if err != nil {
		err = errors.New(UNKNOWN_ERROR)
		return nil, "0", err
	}
	goRequest.Header.Set("Content-Type", writer.FormDataContentType())
	goResponse, err := goClient.Do(goRequest)
	if err != nil {
		fmt.Println("Error sending Go request", err)
		err = errors.New(UNKNOWN_ERROR)
		return nil, "0", err
	}

	defer goResponse.Body.Close()

	if goResponse.StatusCode != http.StatusOK {
		fmt.Println("Status not ok", goResponse.StatusCode)
		var res ErrorResponse
		if err := json.NewDecoder(goResponse.Body).Decode(&res); err != nil {
			err = errors.New(UNKNOWN_ERROR)
			return nil, "0", err
		}

		err = errors.New(res.Detail)
		return nil, "0", err
	}

	cost := goResponse.Header.Get("X-Cost")
	fmt.Println("cost", cost)
	imageBytes, err := io.ReadAll(goResponse.Body)
	if err != nil {
		fmt.Println("Error reading bytes image data:", err)
		err = errors.New(UNKNOWN_ERROR)
		return nil, "0", err
	}

	return imageBytes, cost, nil
}
