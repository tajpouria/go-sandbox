package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 0, Title: "Foo", Content: "About Foo"},
	article{ID: 0, Title: "Bar", Content: "About Bar"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
