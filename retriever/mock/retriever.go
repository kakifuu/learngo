package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever : {Contents=%s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["name"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	fmt.Println("URL:", url)
	return r.Contents
}
