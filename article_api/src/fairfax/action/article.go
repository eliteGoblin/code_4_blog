package action

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	log "fairfax/logging"
	"fairfax/toolkit"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"fairfax/core"
	"fairfax/search_helper"
)

func GetArticleById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	articleId := vars["article_id"]
	if "" == articleId {
		log.Error("empty articleId")
		toolkit.RenderError(res, toolkit.InvalidParam)
		return
	}
	log.Infof("GetArticleById: %s", articleId)
	if article, err := search_helper.DbInstance().GetDocById(articleId); err != nil {
		toolkit.RenderError(res, err.Error())
		return
	}else {
		toolkit.RenderResponse(res, article)
	}
}


func UpsertArticle(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error(err)
		return
	}
	var article core.Article
	err = json.Unmarshal(body, &article)
	if err != nil || !article.IsValid() {
		toolkit.RenderError(res, err.Error())
		return
	}
	article.CreatedTime = fmt.Sprintf("%d", time.Now().Unix())
	log.Infof("%+v", article)
	if err := search_helper.DbInstance().AddDocToDb(&article); err != nil {
		toolkit.RenderError(res, err.Error())
	}
}