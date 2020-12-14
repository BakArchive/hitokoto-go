package dao

import (
	_ "github.com/mattn/go-sqlite3"
	"hitokoto-go/common"
	"log"
	"math/rand"
)

func GetCountFromDB() int {
	count := 0
	rows, err := common.DB.Query("select count() count from sentence")
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return 0
	}
	for rows.Next() {
		rows.Scan(&count)
	}
	return count
}

func GetStringFromDB(source string) string {
	index := rand.Intn(GetCountFromDB()) + 1
	rows, err := common.DB.Query("select word from sentence where id=?", index)
	if err != nil {
		log.Println(err)
		return ""
	}
	var str string
	for rows.Next() {
		rows.Scan(&str)
	}
	return str
}
