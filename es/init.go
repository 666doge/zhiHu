package es

import (
	"context"
	"zhiHu/logger"
	"fmt"
	"strconv"
	mappings "zhiHu/es/mappings"

	elastic "github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
	ctx context.Context
)

func Init() (err error) {
	client, err = elastic.NewClient()
	if err != nil {
		logger.Error("init es failed, err: %v", err)
		return
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		logger.Error("ping es failed, err: %v", err)
		return
	}
	fmt.Printf("ES returned with code %d and version %s\n", code, info.Version.Number)
	createIndexes()
	return
}

func createIndexes() {
	// 先写死两个索引
	createIndex("zhihu_question", mappings.QuestionMap)
	createIndex("zhihu_answer", mappings.AnswerMap)
}

func createIndex(indexName string, mapping string) (err error) {
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		logger.Error("es: IndexExists  failed, indexName: %v, err: %v", indexName, err)
		return
	}
	if !exists {
		// Create a new index.
		newIndex, err1 := client.CreateIndex(indexName).Body(mapping).Do(context.Background())
		if err1 != nil {
			err = err1
			logger.Error("es: create index failed, indexName: %v, err: %v", indexName, err)
			return
		}
		if !newIndex.Acknowledged {
			logger.Error("es: create index not Acknowledged, indexName: %v", indexName)
			return
		}
	}
	return
}

func InsertDoc(indexName string, data []byte, id int64, v interface{})(err error) {

	strId := strconv.FormatInt(id, 10)

	put, err := client.Index().
		Index(indexName).
		Id(strId).
		BodyJson(v).
		Do(context.Background())
	if err != nil {
		logger.Error("es: insert doc failed, indexName: %v", indexName)
		return
	}
	fmt.Printf("Indexed %s to index: %s, type: %s\n", put.Id, put.Index, put.Type)
	return
}

func SearchByMatchQuery(index string, termKey string, termValue interface{}) (result []*elastic.SearchHit, err error){

	matchQuery := elastic.NewMatchQuery(termKey, termValue)
	searchResult, err := client.Search().
		Index(index).
		Query(matchQuery).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	result = searchResult.Hits.Hits
	return
}
