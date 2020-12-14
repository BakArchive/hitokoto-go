package common

import "database/sql"

var (
	Port    int     //端口号
	Path    string  //访问路径
	Sources string  //源目录
	Mode    int     //启动模式：0简易模式，1数据库模式
	DB      *sql.DB //数据库
)
