/************************************************************************************/
/*This program uses google translator to translate whatever you want free,
/*it assembles an url, and analyzes the http result to extract the translated content.
/*Libin.Tian
/*May 23, 2013
/************************************************************************************/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const BASEURL string = "http://translate.google.cn/translate_a/t?"

// type TranslatePara struct {
// 	Client   string
// 	Text     string
// 	H1       string
// 	Sl       string
// 	Tl       string
// 	Ie       string
// 	Oe       string
// 	Multires string
// 	Ssel     string
// 	Tsel     string
// 	Sc       string
// }

type parameters map[string]string

const file = "data.tian"

var paraKeys = []string{"client", "hl", "sl", "tl", "ie",
	"oe", "multires", "ssel", "tsel", "sc"}

func ReadContent(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(content), nil
}

const PROXYURL = "http://web-proxy.rose.hp.com:8080"

func main() {
	content, err := ReadContent(file)
	if err != nil {
		return
	}
	lines := strings.Split(content, "\n")
	fmt.Println(lines[1])

	para := make(parameters)
	para["client"] = "t"
	para["hl"] = "zh-CN"
	para["sl"] = "zh-CN"
	para["tl"] = "en"
	para["ie"] = "UTF-8"
	para["oe"] = "UTF-8"
	para["multires"] = "1"
	para["ssel"] = "0"
	para["tsel"] = "0"
	para["sc"] = "1"

	var strUrl string = BASEURL
	for _, val := range paraKeys {
		strUrl = strUrl + val + "="
		strUrl = strUrl + para[val] + "&"
	}
	//strUrl = strings.TrimRight(strUrl, "&")
	//fmt.Println(strUrl)
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(PROXYURL)
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	for _, line := range lines {
		url := strUrl + "text=" + line
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		strBody := string(body)
		drop := strings.Index(strBody, "]]")
		strTranslation := strBody[3:drop]
		fmt.Println(strTranslation)
	}
}
