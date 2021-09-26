package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
	Sex  int    `json:"sex"`
}
