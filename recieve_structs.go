package main

const (
	SERVER_REQ_DISC          = -1
	SERVER_PROOF_DISC        = 0
	SERVER_SUCCESS_REG       = 1
	SERVER_SUCCESS_AUTH      = 2
	SERVER_ERR_REG_AUTH      = 3
	SERVER_ALL_SONGS         = 4
	SERVER_ADD_SONG          = 5
	SERVER_ALL_TAGS          = 6
	SERVER_USER_SONGS        = 7
	SERVER_USER_BUYING_SONGS = 8
	SERVER_USER_BUY_SONG     = 9
)

type RegisterStruct struct {
	Type     int    `json:"type"`
	Message  string `json:"message"`
	Id       int64  `json:"id"`
	Is_staff bool   `json:"is_staff"`
}
