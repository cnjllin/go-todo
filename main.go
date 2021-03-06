package main

import (
	_ "github.com/pibigstar/go-todo/controller"
	_ "github.com/pibigstar/go-todo/controller/admin"

	"github.com/gogf/gf/frame/g"
	"github.com/pibigstar/go-todo/config"
)

func main() {
	s := g.Server()
	port := config.ServerConfig.Port
	s.SetPort(port)
	host := config.ServerConfig.Host
	s.Domain(host)
	// 开启日志
	s.SetLogPath("log/todo.log")
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	// 开启https
	s.EnableHTTPS("https/ssl.pem", "https/ssl.key")
	s.SetHTTPSPort(443)
	// 开启性能分析，可访问页面/debug/pprof
	s.EnablePprof()
	s.Run()
}
