package gozone

import (
	"fmt"

	"github.com/fame-git/go-zone/gozone/internal/secret"
)

func SayHelloGoZone() {
	fmt.Println("Hello Go-zone")
	secret.SaySecret()
}
