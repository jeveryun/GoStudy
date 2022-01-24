package tree

import (
	"fmt"
	"study1/retriever/mock"
	"study1/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) string {
	return poster.Post("http://www.imooc.com", map[string]string{
		"name": "aaa",
		"course": "golang"})
}

type RetrieverPoster interface {
	Retriever
	Poster
	// Connect(host string)
}

func session(s RetrieverPoster) string{
	s.Post(url, map[string]string{"contents": "another faked imooc.com"})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	
	var r Retriever
	r2  := mock.Retriever{"this is mooc"}
	r = &mock.Retriever{"this is 111111"}
	inspect(r)
	r = &real.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	inspect(r)

	if realRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(realRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//try session
	fmt.Println("Try a session")
	fmt.Println(session(&r2))
}