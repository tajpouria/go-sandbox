package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	artlist := getAllArticles()

	if len(artlist) != len(articleList) {
		t.Fail()
	}

	for i, a := range artlist {

		if a.Contnet != articleList[i].Contnet ||
			a.ID != articleList[i].ID ||
			a.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}

}
