package main

import (
	"flag"
	"hitokoto-go/common"
	"hitokoto-go/handler"
	"log"
	"net/http"
	"strconv"
)

func init() {
	flag.StringVar(&common.Path, "P", "/", "设置访问路径")
	flag.IntVar(&common.Port, "p", 9000, "设置端口号")
	flag.StringVar(&common.Sources, "s", common.GetPathHere()+common.Separator+"hitokoto.txt", "设置数据源")
	flag.Parse()
	common.ModeSet(common.Sources)
	if common.Mode == 1 {
		common.ConnectDB(common.Sources)
	}
}

func main() {
	http.HandleFunc(common.Path, handler.GetOneHitokoto)
	err := http.ListenAndServe(":"+strconv.FormatInt(int64(common.Port), 10), nil)
	if err != nil {
		log.Fatal(err)
	}
	if common.Mode == 1 {
		common.DB.Close()
	}
}
