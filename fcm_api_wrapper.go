package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const fcmSendURL = "https://fcm.googleapis.com/fcm/send"

type FCMAPIWrapper struct {
	Authorization string
}

type fCMRequestBody struct {
	To   string         `json:"to"`
	Data fCMRequestData `json:"data"`
}

type fCMRequestData struct {
	Message string `json:"message"`
}

type fCMError struct {
	StatusCode int
	Body       string
}

func (e *fCMError) Error() string {
	return fmt.Sprintf("Error on notfication request to FCM: \nStatus %d\nBody>\n%s", e.StatusCode, e.Body)
}

// Send sends push notification request
func (w *FCMAPIWrapper) Send(to string, message string) error {
	data := fCMRequestData{message}
	body := fCMRequestBody{to, data}
	byteData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", fcmSendURL, bytes.NewReader(byteData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", w.Authorization)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		err = &fCMError{
			StatusCode: res.StatusCode,
			Body:       string(bodyBytes),
		}
	}
	return err
}

func NewFcmApiWrapper(authorization string) *FCMAPIWrapper {
	wrapper := new(FCMAPIWrapper)
	wrapper.Authorization = authorization
	return wrapper
}
