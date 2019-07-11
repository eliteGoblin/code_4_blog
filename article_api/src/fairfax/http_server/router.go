package http_server

import (
	"net/http"
	"fairfax/action"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", action.Index).Methods("GET")
	router.HandleFunc("/articles/{article_id}", action.GetArticleById).Methods("GET")
	router.HandleFunc("/articles", action.UpsertArticle).Methods("POST")
	router.HandleFunc("/tags/{tag_name}/{date}", action.RetrieveArticle).Methods("GET")
	return router
}

