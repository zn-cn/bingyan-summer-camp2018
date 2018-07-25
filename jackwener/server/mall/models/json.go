package models

// 定义了返回的统一的Json格式

type UserJson struct {
	User   User
	Status int
	Message string
}

type GoodJson struct {
	Good    []Good
	Status int
	Message string
	Url     string
}

type SearchJson struct {
	Content  []string
	Status   int
	Message string
}