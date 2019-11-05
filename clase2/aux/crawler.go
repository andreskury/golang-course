package main

import (
	"errors"
	"net/http"

	"golang.org/x/net/html"
)

func GetLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("El request no retornó 200 OK")
	}

	z := html.NewTokenizer(resp.Body)

	links := []string{}
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return links, nil
		case tt == html.StartTagToken:
			t := z.Token()

			isAnchor := t.Data == "a"
			if isAnchor {
				for _, a := range t.Attr {
					if a.Key == "href" {
						links = append(links, a.Val)
						break
					}
				}
			}
		}
	}

	return links, nil
}
