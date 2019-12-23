package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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

// Send sends push notification request
func (w *FCMAPIWrapper) Send(to string, message string) error {
	log.Println("to: ", to)
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
		log.Println("Reuqest is handled incorrectly")
		log.Println(res.StatusCode, res.Header)
	}
	log.Println("body: ", string(byteData))
	resBytes, _ := ioutil.ReadAll(res.Body)
	reqBytes, _ := ioutil.ReadAll(req.Body)
	log.Println("reqBody: ", string(reqBytes))
	log.Println("resBody: ", string(resBytes))
	return nil
}

func NewFcmApiWrapper(authorization string) *FCMAPIWrapper {
	wrapper := new(FCMAPIWrapper)
	wrapper.Authorization = authorization
	return wrapper
}
