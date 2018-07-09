package main

import (
	"flag"
	"fmt"

	"github.com/ComputePractice2018/videohosting/backend/utils"
)

func main() {
	name := flag.String("name", "Michael", "имя для приветствия")
	flag.Parse()
	fmt.Println(utils.GetHelloWorldString(*name))
}
