package main

import (
	"log"

	domain "github.com/itratos/go-kit/domain_checker"
)

func main() {
	info, err := domain.CheckDom("envioskv.com")
	if err != nil {
		log.Println(err)
	}
	log.Println(info)
}
