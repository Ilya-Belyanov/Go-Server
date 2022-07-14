package main

type RegisterStruct struct {
	Type     int    `json:"type"`
	Message  string `json:"message"`
	Id       int64  `json:"id"`
	Is_staff bool   `json:"is_staff"`
}
