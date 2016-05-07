// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FlyingFries/ketchup/sms/simpay"
)

var key = flag.String("key", "", "Auth key")
var secret = flag.String("secret", "", "Auth secret")

var service_id = flag.String("service_id", "", "Usluga, id w systemie, nie kod sms")
var number = flag.String("number", "", "Numer")
var code = flag.String("code", "", "Kod zwrotny")

func main() {
	flag.Parse()
	if *key == "" || *secret == "" || *service_id == "" || *number == "" || *code == "" {
		fmt.Println("One of flags is empty!")
		flag.Usage()
		os.Exit(1)
	}

	var auth = simpay.Auth{*key, *secret}
	price, err := simpay.Check(auth, *service_id, *number, *code)
	fmt.Printf("Check: %d err: %v\n", price, err)
}
