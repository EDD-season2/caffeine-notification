package main

import "testing"

import "encoding/json"

type body struct {
	To   string `json:"to"`
	Data data   `json:"data"`
}

type data struct {
	Message string `json:"message"`
}

func TestMarshal(t *testing.T) {
	data := data{"hello"}
	bytes, err := json.Marshal(data)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(data)
	t.Log(bytes)
	t.Log(string(bytes))
}

func TestMarshalNested(t *testing.T) {
	data := data{"hello"}
	body := body{"abcedf", data}
	bytes, err := json.Marshal(body)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(string(bytes))
}
