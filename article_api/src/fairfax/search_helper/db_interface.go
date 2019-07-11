package search_helper

import "fairfax/core"

type SearchByTagResult struct {
	Tag string 	 			`json:"tag"`
	Count int 				`json:"count"`
	Articles []string 			`json:"articles"`
	RelatedTags []string 	`json:"related_tags"`
}

type DbHelper interface {
	AddDocToDb(article *core.Article) error
	GetDocById(id string) (article *core.Article, err error)
	SearchByTag(tag string, date string) (res *SearchByTagResult, err error)
}

