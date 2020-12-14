package dao

import (
	"database/sql"
	"fmt"
	"hitokoto-go/common"
	"testing"
)

func TestGetCountFromDB(t *testing.T) {
	common.ConnectDB("D:\\GOPATH\\src\\hitokoto-go\\hitokoto.db")
	var stat *sql.Stmt
	for i := 0; i < 1000; i++ {
		sql := fmt.Sprint("insert into sentence (word) values('测试语句", i, "');")
		stat, _ = common.DB.Prepare(sql)
		stat.Exec()
	}
	stat.Close()
}
