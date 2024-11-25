package main

import (
	"fmt"

	"github.com/LavaJover/Marketplace/user-service/internal/config"
)

func main(){
	cfg := config.MustLoad()
	fmt.Println(cfg)
}