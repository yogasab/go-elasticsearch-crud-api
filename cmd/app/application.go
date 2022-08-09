package app

import (
	"log"
	"net/http"
	"yogasab/go-elasticsearch-crud-api/internal/pkg/storage/elasticsearch"
	"yogasab/go-elasticsearch-crud-api/internal/post"

	"github.com/julienschmidt/httprouter"
)

func StartApplication() {
	elastic, err := elasticsearch.New([]string{"http://127.0.0.1:9206"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err = elastic.CreateIndex("posts"); err != nil {
		log.Fatalln(err)
	}

	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	if err != nil {
		log.Fatalln(err)
	}
	postService := post.NewPostService(postStorage)
	postHandler := post.NewPostHandler(postService)

	router := httprouter.New()
	router.HandlerFunc("POST", "/api/v1/posts", postHandler.InsertDocument)
	log.Fatalln(http.ListenAndServe(":5000", router))
}
