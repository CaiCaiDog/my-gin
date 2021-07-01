package config

// 配置
var ServerDatabaseConfig = map[string]string{
	"host": "127.0.0.1",
	"port": "3306",
	"database": "chart",
	"username": "root",
	"password": "root",
	"charset": "utf8",
}

// get config
func Database (key string) (value string) {
	value = ServerDatabaseConfig[key]
	return
}