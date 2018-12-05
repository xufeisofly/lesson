package main

import (
	"fmt"
	"lession/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, contents map[string]string) string
}

type Session interface {
	Retriever
	Poster
}

const url = "https://www.niltouch.cn"

func session(s Session) string {
	s.Post(url, map[string]string{
		"contents": "Another mock niltouch.cn",
	})
	return s.Get(url)
}

func main() {
	r := &mock.Retriever{"This is a mock content"}

	fmt.Println(session(r))
}
