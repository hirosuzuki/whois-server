package main

import (
	"fmt"

	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

func main() {
	result, err := whois.Whois("google.com")
	if err == nil {
		fmt.Println(result)
		result, err := whoisparser.Parse(result)
		if err == nil {
			fmt.Println(result.Domain.Punycode)
		}
	}
}
