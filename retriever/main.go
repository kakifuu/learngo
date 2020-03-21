package main

import (
	"fmt"
	"learngo/retriever/mock"
	kaki "learngo/retriever/real"
	"time"
)

const url = "http://www.baidu.com"

var info = map[string]string{
	"name":   "kaki",
	"gender": "male",
}

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, info)
}

func session(s RetrieverPoster) string {
	s.Post(url, info)
	return s.Get(url)
}

func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "Hi there. This is a fake baidu."}
	inspect(r)
	r = &kaki.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   2 * time.Minute,
	}
	//fmt.Println(download(r))

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock retriever")
	}

	inspect(r)

	fmt.Println(session(&mock.Retriever{Contents: "Hello"}))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Println("> Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *kaki.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
