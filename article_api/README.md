

#### Set up (ubuntu 16.04)

```shell
git clone https://github.com/eliteGoblin/article_api.git
cd article_api/src
docker-compose build
sudo sysctl -w vm.max_map_count=262144
docker-compose up
```


#### Tool scripts


*   curl -XPOST "localhost:8080/articles" -d@article_1.json
*   curl "localhost:8080/articles/1"
*   curl "localhost:8080/tags/health/2018-08-11"

to delete es index use:
curl -XDELETE localhost:9200/articles

#### About the solution

##### structures of code: 

```shell
src/fairfax
├── action
│   ├── article.go
│   ├── index.go
│   └── search.go
├── core
│   └── article.go
├── Dockerfile
├── Gopkg.lock
├── Gopkg.toml
├── http_server
│   ├── router.go
│   └── server.go
├── logging
│   ├── formatter.go
│   ├── logger.go
│   └── log_hook.go
├── main.go
├── search_helper
│   ├── db_interface.go
│   └── elastic_search_db.go
├── toolkit
│   └── resp.go
└── vendor
```

*  **action**: for http handler
*  **code**: for core data structures
*  **vendor** for all dep packages: using godep to manage deps
*  **logging** for customized log function: function name, file name, line number aded
*  **http_server** for http router
*  **search_helper** for elastic helper: interface for decouple client code and elastic library code

##### About the packages 

*  elastic search, using package: github.com/olivere/elastic
*  logging: customized from github.com/sirupsen/logrus
*  http router using github.com/gorilla/mux


#### Assumptions

about the GET /tag/{tagName}/{date} API assume it should behave like following: 

all articles from same day: 20160922

```javascript
{
  "id": "1",
  "tags" : ["health", "sports"]
}

{
  "id": "2",
  "tags" : ["health", "science"]
}

{
  "id": "3",
  "tags" : ["health"]
}

{
  "id": "4",
  "tags" : ["cooking"]
}
```

after GET /tag/health/20160922

for the result field returned:

*  related_tags should return: ["science", "sports"]
*  count should be 3 (3 articles have health label)
*  articles should be ["1", "2"] sorted by created time



#### About testing


It is an very interesting and practical task; before i have no knowledge in ElasticSearch, but i wanted to see how it works long times ago, because full text search is so useful today.  I am glad that after the test, i learned some basic stuff about it, a very good subject for my [tech blog](http://elitegoblin.github.io)

In my last role, we build a FAQ system mainly use bm25, word2vec, it is said that better to use elasticsearch alongside it, because es provide some quick and powerful searching functions.

Is took me around one day and half to learn some basic concept about ElasticSearch, its APIs and how to put everything altogether.