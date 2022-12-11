package parser

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)



var packagesToSkip = map[string]bool {
	"cloud.google.com/go/home": true,
	"cloud.google.com/go/getting-started": true,
	"cloud.google.com/go/docs/reference": true,
}

var subpackageRe = regexp.MustCompile(`cloud.google.com/go/[a-z-/0-9]+`)

func DiscoverSubpackages() ([]string, error) {
	var subpackages = make([]string, 0)
	resp, err := http.Get("https://pkg.go.dev/cloud.google.com/go")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	matches := subpackageRe.FindAllString(string(body), -1)
	for _, match := range matches {
			if strings.Contains(match, "internal") {
				continue
			}
			if _, ok := packagesToSkip[match]; !ok {
				subpackages = append(subpackages, match)
			}
	}

	return subpackages, nil
}
