package model

type Example struct {
	Id   int64  `column:"gorm:id"`
	Name string `column:"gorm:name"`
}
