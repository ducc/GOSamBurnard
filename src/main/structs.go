package main

type (
	config struct {
		Http struct {
			Port string `json:"port"`
		} `json:"http"`
		Templates struct {
			Directory string `json:"directory"`
		} `json:"templates"`
		Database struct {
            Host string `json:"host"`
            Database string `json:"database"`
            Username string `json:"username"`
            Password string `json:"password"`
            Ssl bool `json:"ssl"`
		} `json:"database"`
	}
)
