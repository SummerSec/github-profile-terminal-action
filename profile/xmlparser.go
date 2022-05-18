package profile

import (
	"github.com/beevik/etree"
	"io/ioutil"
	"net/http"
	"strings"
)

func Parser(res string) []string {
	resp, err := http.Get(res)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	xml := string(body)
	doc := etree.NewDocument()
	err = doc.ReadFromString(xml)
	root := doc.SelectElement("urlset")
	var maps []string
	i := 0
	for _, element := range root.SelectElements("url") {
		if i > 7 {
			break
		} else {
			text := element.SelectElement("loc").Text()
			if !strings.Contains(text, "resources") || strings.Compare(text, "https://sumsec.me/") != 0 {
				maps = append(maps, text)
				i++
			}

		}
	}
	return maps
}
