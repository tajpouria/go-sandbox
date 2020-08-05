package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
)

type student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

const (
	elasticSearchURL = "http://localhost:9200"
)

func main() {
	ctx := context.Background()
	esClient, err := getESClient()

	if err != nil {
		fmt.Println("Connection err!")
		panic(err)
	}

	if false { // https://github.com/olivere/elastic/wiki/Index
		newStudent := student{
			Name:         "Foo",
			Age:          12,
			AverageScore: 18.6,
		}

		dataJSON, err := json.Marshal(newStudent)
		js := string(dataJSON)

		idx, err := esClient.Index().Index("students").Type("_doc").BodyJson(js).Do(ctx)

		if err != nil {
			panic(err)
		}

		fmt.Println(idx, "Created!")
	}

	if true { // https://github.com/olivere/elastic/wiki/Search

		searchSource := elastic.NewSearchSource()
		searchSource.Query(elastic.NewMatchQuery("name", "Foo"))

		/* this block will basically print out the es query */
		queryStr, err1 := searchSource.Source()
		queryJs, err2 := json.Marshal(queryStr)

		if err1 != nil || err2 != nil {
			fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
		}
		fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))
		/* until this block */

		searchService := esClient.Search().Index("students").SearchSource(searchSource)

		searchResult, err := searchService.Do(ctx)
		if err != nil {
			panic(err)
		}

		var students []student

		for _, hit := range searchResult.Hits.Hits {
			var s student
			err := json.Unmarshal(*hit.Source, &s)

			if err != nil {
				fmt.Println("[Getting Students][Unmarshal] Err=", err)
			}

			students = append(students, s)
		}

		if err != nil {
			fmt.Println("Fetching student fail: ", err)
		} else {
			for _, s := range students {
				fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
			}
		}

	}

}

func getESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(elasticSearchURL), elastic.SetSniff(false), elastic.SetHealthcheck(false))

	fmt.Println("Es initialized...")

	return client, err
}
