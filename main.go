package main

import (
	//log "github.com/sirupsen/logrus"
	"github.com/vitaminwater/daryl_db"
)

type Test struct {
	Slug string `db:"slug" access:"i,s"`
	Name string `db:"name" access:"i,u,s"`
}

func init() {
	daryl_db.Init()
}

func main() {
	test := Test{"eR@3rlka", "lolname"}
	daryl_db.Insert(test)
}
