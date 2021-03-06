package reptile

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"news_watch_notice/utils"
	"strings"
	"time"
)

/*
* @Author:15815
* @Date:2019/4/30 0:10
* @Name:reptile
* @Function:
 */

// 爬虫CoNews网页
func GetNewsContent(publishTime time.Time) (e error, content []string) {
	var baseUrl string
	c := colly.NewCollector()
	//t:=time.Now().Add(-time.Hour*time.Duration(24))
	data := publishTime.Format("2006-01-02")
	dateOther := publishTime.Format("2006-01-2")
	// Find and visit all links
	c.OnHTML("div > h4 > a", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, data) {
			baseUrl = e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, baseUrl)
		} else if strings.Contains(e.Text, dateOther) {
			baseUrl = e.Attr("href")
			fmt.Printf("Link found: %q -> %s\n", e.Text, baseUrl)
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	e = c.Visit("https://gocn.vip/question/category-14")

	if e != nil {
		return e, nil
	}
	if baseUrl == "" {
		return errors.New("news not update"), nil
	}
	b := colly.NewCollector()

	// Find and visit all links
	i := 0
	contentList := make([]string, 15)
	b.OnHTML("div.mod-body > div > ol > li", func(e *colly.HTMLElement) {
		if e.Text != "" {
			contentList[i] = utils.TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text))
			i++
			fmt.Printf("%d:%q\n", i, utils.TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text)))
		}
	})
	b.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	e = b.Visit(baseUrl)
	if e != nil {
		return e, nil
	}
	var flag bool
	for _, c := range contentList {
		if c != "" {
			flag = true
			break
		}
	}
	if !flag {
		c := colly.NewCollector()
		c.OnHTML("div.mod-body > div > p", func(e *colly.HTMLElement) {
			if e.Text != "" {
				contentList[i] = utils.TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text))
				i++
				fmt.Printf("%d:%q\n", i, utils.TrimQuotes(fmt.Sprintf("%d. %s\n\n", i+1, e.Text)))
			}
		})
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})
		e = c.Visit(baseUrl)
		if e != nil {
			return e, nil
		}
	}
	return nil, contentList

}
