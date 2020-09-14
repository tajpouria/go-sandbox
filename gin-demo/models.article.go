package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Contnet string `json:"content"`
}

var articleList = []article{
	article{ID: 0, Title: "Foo", Contnet: "About Foo"},
	article{ID: 0, Title: "Bar", Contnet: "About Bar"},
}

func getAllArticles() []article {
	return articleList
}
