package profile

import (
	"github.com/beevik/etree"
	"io/ioutil"
	"net/http"
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
		if i > 8 {
			break
		} else if i < 4 {
			i++
		} else {
			maps = append(maps, element.SelectElement("loc").Text())
			i++
		}
	}
	return maps
}
