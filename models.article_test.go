package main

import "testing"

func TestGetAllArticles(t *testing.T){
	alist := getAllArticles()

	if len(alist) != len(articleList){
		t.Fatal()
	}

	for i,v := range alist{
		if v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title{
			t.Fatal()
			break
		}
	}
}
