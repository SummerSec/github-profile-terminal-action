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
		if i > 5 {
			break
		} else {
			text := element.SelectElement("loc").Text()
			// text 含有resources字符串 则跳过 否则添加到maps中 并且i++ 如果text 和https://sumsec.me/ 一样 则跳过
			if !strings.Contains(text, "resources") {
				if !strings.EqualFold(text, "https://sumsec.me/") {
					maps = append(maps, text)
					i++
				}
			}
			//maps = append(maps, element.SelectElement("loc").Text())
			//fmt.Println(element.SelectElement("loc").Text())
			//i++
		}
	}
	return maps
}
