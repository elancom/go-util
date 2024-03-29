package main

import (
	"flag"
	"fmt"
	"github.com/elancom/go-util/lang"
	"log"
)

func main() {
	log.Println(lang.GetEnv())
	var (
		host = flag.String("host", "", "go flag test")
		port = flag.Int("port", 0, "请输入端口号")
	)
	flag.Parse() // 解析参数
	fmt.Printf("%s:%d\n", *host, *port)
}
