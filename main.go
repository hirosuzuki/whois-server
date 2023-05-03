package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
)

type WhoisInfo struct {
	Result struct {
		Domain      string            `json:"domain"`
		Nameservers []string          `json:"nameservers"`
		Attributes  map[string]string `json:"attributes"`
		Raw         string            `json:"raw"`
	} `json:"result"`
	Cached  bool   `json:"cached"`
	Fetched string `json:"fetched"`
}

func getWhoisInfo(domain string) (*WhoisInfo, error) {
	result, err := whois.Whois("google.com")
	if err == nil {
		fmt.Println(result)
		result, err := whoisparser.Parse(result)
		if err == nil {
			fmt.Println(result.Domain.Punycode)
		}
	}
	info := WhoisInfo{}
	return &info, nil
}

func getDomainWhois(c echo.Context) error {
	domain := c.Param("domain")
	whoisInfo, err := getWhoisInfo(domain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error fetching Whois information"})
	}
	return c.JSON(http.StatusOK, whoisInfo)
}

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/whois/:domain", getDomainWhois)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
