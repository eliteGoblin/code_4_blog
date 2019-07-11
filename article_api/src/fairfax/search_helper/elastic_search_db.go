package search_helper

import (
	"fairfax/core"
	"sync"
	"github.com/olivere/elastic"
	"time"
	log "fairfax/logging"
	"os"
	"context"
	"reflect"
	"fmt"
)


const (
	esConnRetryTimes = 20
)

type ESDb struct {
	esClient *elastic.Client
}



func NewESDb(esClient *elastic.Client) DbHelper {
	return &ESDb{
		esClient : esClient,
	}
}

var once sync.Once
var dbClient DbHelper

func DbInstance() DbHelper {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("%v", r)
			os.Exit(1)
		}
	}()
	once.Do(func() {
		for i := 0; i < esConnRetryTimes; i ++ {
			if client, err := elastic.NewClient(
				elastic.SetURL("http://elasticsearch:9200"),
				elastic.SetSniff(false),
			); err == nil {
				_, err = client.CreateIndex("article").Do(context.Background())
				if err != nil {
					panic(err)
				}
				dbClient = NewESDb(client)
				log.Info("successfully connected to es, cheers")
				break
			}else {
				log.Infof("fail to connect to es %s, retry in 3 sec", err)
				time.Sleep(3 * time.Second)
			}
		}
		if dbClient == nil {
			panic("cannot connect to es, aborting...")
		}
	})
	return dbClient
}


func (selfPtr *ESDb)AddDocToDb(article *core.Article) error {
	_, err := selfPtr.esClient.Index().
		Index("article").
		Type("doc").
		Id(article.Id).
		BodyJson(article).
		Refresh("wait_for").
		Do(context.Background())
	return err
}


func (selfPtr *ESDb)GetDocById(id string)  (article *core.Article, err error) {
	query := elastic.NewBoolQuery()
	q1 := elastic.NewMatchQuery("id", id)
	query.Must(q1)

	searchResult, err := selfPtr.esClient.Search().
		Index("article").             // search in index "article"
		Query(query).           			// specify the query
		Pretty(true).                // pretty print request and response JSON
		Do(context.Background())            // execute
	if err != nil {
		return nil, err
	}

	var esItem core.Article
	for _, item := range searchResult.Each(reflect.TypeOf(esItem)) {
		if t, ok := item.(core.Article); ok {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("id %s not found", id)
}


func (selfPtr *ESDb)SearchByTag(tag string, date string) (res *SearchByTagResult, err error){
	query := elastic.NewBoolQuery()
	q1 := elastic.NewTermsQuery("tags", tag)
	q2 := elastic.NewMatchQuery("date", date)
	query.Must(q1, q2)

	searchResult, err := selfPtr.esClient.Search().
		Index("article").             		// search in index "article"
		Query(query).           					// specify the query
		Pretty(true).                		// pretty print request and response JSON
		Sort("created_time.keyword", false). 	// sort by "id" field, ascending
		Do(context.Background())            		// execute
	if err != nil {
		return nil, err
	}
	response := generateResponse(tag, searchResult)
	return response, nil
}

func generateResponse(tag string, searchResult *elastic.SearchResult) *SearchByTagResult {
	var response SearchByTagResult
	var esItem core.Article
	response.Tag = tag
	mapTagsMap := make(map[string]bool)
	for _, item := range searchResult.Each(reflect.TypeOf(esItem)) {
		if t, ok := item.(core.Article); ok {
			for _, v := range t.Tags {
				mapTagsMap[v] = true
			}
			if response.Count < 10 {
				response.Articles = append(response.Articles, t.Id)
			}
			response.Count ++
		}
	}
	delete(mapTagsMap, tag)
	for tagName := range mapTagsMap {
		response.RelatedTags = append(response.RelatedTags, tagName)
	}
	return &response
}