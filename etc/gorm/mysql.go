package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type NArFilterCategory struct {
	Idx  int
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "circusceo:circusARDB7@tcp(circuscms.cwgeharnfkuv.ap-northeast-2.rds.amazonaws.com)/circusceo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var nArFilterCategory NArFilterCategory
	db.Raw("SELECT Idx, Name FROM nArFilterCategory WHERE Idx = ?", 1).Scan(&nArFilterCategory)
	//db.First(&result, "Idx = ?", 1)

	log.Printf("%#v", nArFilterCategory)
}
