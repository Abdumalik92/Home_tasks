package main

type User struct {
	ID       int    `gorm:"column:id_us;primary_key"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}
type Log struct {
	ID       int    `gorm:"column:id;primary_key"`
	Username string `gorm:"column:username"`
}
type Journal struct {
	ID       int    `gorm:"column:id_me;primary_key"`
	Username string `gorm:"column:username"`
	Messag   string `gorm:"column:messag"`
	Time     string `gorm:"column:time"`
}
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}
