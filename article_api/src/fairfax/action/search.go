package action

import (
	"net/http"
	"github.com/gorilla/mux"
	log "fairfax/logging"
	"fairfax/toolkit"
	"fairfax/search_helper"
)

func RetrieveArticle(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tagName := vars["tag_name"]
	date := vars["date"]
	if "" == tagName || date == "" {
		log.Error("empty articleId")
		toolkit.RenderError(res, toolkit.InvalidParam)
		return
	}
	log.Infof("RetrieveArticle: %s %s", tagName, date)
	if article, err := search_helper.DbInstance().SearchByTag(tagName, date); err != nil {
		toolkit.RenderError(res, err.Error())
		return
	}else {
		toolkit.RenderResponse(res, article)
	}
}