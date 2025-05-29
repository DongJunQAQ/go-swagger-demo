package api

type User struct {
	//Required: true
	Name     string `json:"name"` //该字段是必须的
	Nickname string `json:"nickname"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}
