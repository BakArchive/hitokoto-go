package handler

import (
	"fmt"
	"hitokoto-go/common"
	"hitokoto-go/dao"
	"log"
	"math/rand"
	"net/http"
)

func GetOneHitokoto(w http.ResponseWriter, r *http.Request) {
	if common.Mode == 0 {
		log.Println("file request")
		result, err := dao.GetStringFromFile(common.Sources)
		if err != nil {
			log.Fatal(err)
		}
		size := len(result)
		index := rand.Intn(size)
		fmt.Fprintf(w, result[index])
	} else {
		log.Println("db request")
		str := dao.GetStringFromDB(common.Sources)
		fmt.Fprintln(w, str)
	}
}
