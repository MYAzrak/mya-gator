package main

import (
	"fmt"

	"github.com/MYAzrak/mya-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := cfg.SetUser("mya"); err != nil {
		fmt.Println(err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg.DbURL)
	fmt.Println(cfg.CurrentUserName)
}
