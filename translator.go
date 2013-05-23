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

var paraKeys = []string{"client", "text", "hl", "sl", "tl", "ie",
	"oe", "multires", "ssel", "tsel", "sc"}

func main() {
	para := make(parameters)
	para["client"] = "t"
	para["text"] = "我喜欢你"
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
	strUrl = strings.TrimRight(strUrl, "&")
	fmt.Println(strUrl)

	resp, err := http.Get(strUrl)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
