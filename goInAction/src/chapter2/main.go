package main

import (
	_ "chapter2/sample/matchers"
	"chapter2/sample/search"
	"log"
	"os"
)

func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
