package config

// 配置
var ServerAppConfig = map[string]string {
	"app_mode": "debug",								// 项目运行模式(debug, release, test)
	"host": "127.0.0.1",							// 项目ip
	"port": "9000",									// 项目端口
	"html_path": "app/Http/view/*",					// 视图文件加载路径
	"delim_left": "{{{",								// 视图变量渲染符（左）
	"delim_right": "}}}",							// 视图变量渲染符（右）
}

// get config
func App (key string) (value string) {
	value = ServerAppConfig[key]
	return
}