package cmd

import (
	"time"
)

type Config struct {
	Host	string
	Port	int
	TLS	bool
	Token	string
}

type Subscription struct {
	Id string
	Description string `json:"description"`
	Subject     struct {
		Entities []struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"entities"`
		Condition struct {
			Attrs []string `json:"attrs"`
		} `json:"condition"`
	} `json:"subject"`
	Notification struct {
		HTTP struct {
			URL string `json:"url"`
		} `json:"http"`
		Attrs []string `json:"attrs"`
	} `json:"notification"`
	Expires    time.Time `json:"expires"`
	Throttling int       `json:"throttling"`
}
