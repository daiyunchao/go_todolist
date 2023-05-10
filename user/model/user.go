package model

type User struct {
	Id       string `json:"id" bson:"id"`
	Nickname string `json:"nickname" bson:"nickname"`
	Password string `json:"-" bson:"password"`
}
