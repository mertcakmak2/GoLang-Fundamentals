package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	client := &http.Client{}

	url := "https://jsonplaceholder.typicode.com/posts"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	json.Unmarshal(responseBody, &posts)

	fmt.Println(resp.Status)
	for _, post := range posts {
		fmt.Println(post.Id, post.UserId, post.Title)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		response := &Response{StatusCode: resp.StatusCode, Status: resp.Status, Body: posts}
		res, _ := json.Marshal(response)
		w.Write(res)
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

type Response struct {
	StatusCode int    `json:"statusCode"`
	Status     string `json:"status"`
	Body       []Post `json:"body"`
}

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
