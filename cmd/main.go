package main

import (
	"fmt"
	Config "github.com/Oracl0s/Mini-prodject.git/internal/config"
)

func main() {
	cfg := Config.Mustlaod()
	fmt.Println(cfg)
	//todo:init logger
	//todo:init storage
	//todo:init router
	//todo:run server
	//todo
}
