package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type envVars struct {
}

func main() {
	// 除程序本身外，第一个参数即为模板文件路径
	tplFile := os.Args[1]

	// 获取文件basename
	baseName := filepath.Base(tplFile)

	ev := new(envVars)

	// 自定义函数
	funcMap := template.FuncMap{"GetVar": ev.GetVar}

	// 模板文件装载自定义函数
	tpl := template.Must(template.New(baseName).Funcs(funcMap).ParseFiles(tplFile))

	// 渲染模板生成最终文件
	finalFile, _ := os.OpenFile(strings.Split(tplFile, ".tpl")[0], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	tpl.Execute(finalFile, ev)
}

// 获取环境变量
func (ev *envVars) GetVar(key string) string {
	return strings.ReplaceAll(os.Getenv(key), "\\n", "\n")
}
