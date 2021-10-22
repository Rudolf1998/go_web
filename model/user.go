package model

// User 用户格式
type User struct {
        ID       int    `json:"id"`
        Mobile   string `json:"mobile"`
        Name     string `json:"name"`
        Password string `json:"password"`
}

