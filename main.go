package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var grades = [6]string{"一年级", "二年级", "三年级", "四年级", "五年级", "六年级"}
	var subjects = [3]string{"语文", "数学", "英语"}
	const youxueHome string = `https://video.ijiwen.com/`
	client := &http.Client{}
	var seq int = 1
	fmt.Print("序号,年级,科目,课程名,集名,视频下载链接\n")
	for _, grade := range grades {
		for _, subject := range subjects {

			// 具体科目的各种视频的链接
			courseListLink := fmt.Sprintf("index.php?grade_name=%s&subject=%s", grade, subject)
			kindVideoDoc := getBodyData(client, youxueHome+courseListLink, courseListLink, true)
			kindVideoDiv := kindVideoDoc.Find("body").Find("div.course__list-item")
			for i := 0; i < kindVideoDiv.Length(); i++ {
				attr, err := kindVideoDiv.Eq(i).Attr("onclick")
				if !err {
					panic("视频连接不存在")
				}

				// 具有完整的一套课程播放地址
				lastCourseVideoLink := strings.Split(attr, `'`)[1]
				doc := getBodyData(client, youxueHome+lastCourseVideoLink, lastCourseVideoLink, false)

				// 课程名
				title := doc.Find("body").Find("div.title").Text()

				// 抓取视频链接和课集名
				beat := doc.Find("body").Find("div.details__cnt>ul>li")
				for j := 0; j < beat.Length(); j++ {

					col := beat.Eq(j).Find("div").Text()

					attr, err := beat.Eq(j).Attr("onclick")
					if !err {
						panic("不可能不存在视频链接")
					}
					//javascript:play('7', '387', '7612e23c2b024c66661e.mp4', '7')
					//https://oss.dkclass365.com/video/7/381_e15f16f72d62be50bfe7.mp4
					part := strings.Split(attr, `'`)
					mp4path := fmt.Sprintf("https://oss.dkclass365.com/video/%s/%s_%s", part[1], part[3], part[5])
					fmt.Printf("%d,%s,%s,%s,%s,%s\n", seq, grade, subject, title, col, mp4path)
					seq++

				}
			}

		}
	}
}
func getBodyData(client *http.Client, link, path string, homeLink bool) *goquery.Document {
	//fmt.Println("请求连接：", link)
	r, err := http.NewRequest("GET", link, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("authority", "video.ijiwen.com")
	r.Header.Add("method", "GET")
	r.Header.Add("path", url.QueryEscape(path))
	r.Header.Add("scheme", "https")
	r.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	r.Header.Add("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	if homeLink {
		r.Header.Add("cache-contro", "max-age=0")
	}
	r.Header.Add("cookie", "UID=232992; SESSION=4160693168a46a793b39a98beada4a96; data=%7B%22people_id%22%3A232992%2C%22people_list%22%3A%5B%7B%22school_id%22%3A5642%2C%22school_name%22%3A%22%5Cu5b81%5Cu6ce2%5Cu5e02%5Cu911e%5Cu5dde%5Cu533a%5Cu4e0b%5Cu5e94%5Cu8857%5Cu9053%5Cu4e2d%5Cu6d77%5Cu5c0f%5Cu5b66%22%2C%22class_id%22%3A204024%2C%22class_name%22%3A%22%5Cu4e8c%5Cu5e74%5Cu7ea73%5Cu73ed%22%2C%22child_id%22%3A2547840%2C%22child_name%22%3A%22%5Cu5f20%5Cu67ef%5Cu6db5%22%2C%22grade%22%3A2%7D%2C%7B%22school_id%22%3A5642%2C%22school_name%22%3A%22%5Cu5b81%5Cu6ce2%5Cu5e02%5Cu911e%5Cu5dde%5Cu533a%5Cu4e0b%5Cu5e94%5Cu8857%5Cu9053%5Cu4e2d%5Cu6d77%5Cu5c0f%5Cu5b66%22%2C%22class_id%22%3A204024%2C%22class_name%22%3A%22%5Cu4e8c%5Cu5e74%5Cu7ea73%5Cu73ed%22%2C%22child_id%22%3A232992%2C%22child_name%22%3A%22%5Cu5f20%5Cu67ef%5Cu6db5%5Cu5bb6%5Cu957f%22%2C%22grade%22%3A2%7D%5D%7D; expiry_time=1602986238; subject=%E8%8B%B1%E8%AF%AD; grade_name=%E4%B8%80%E5%B9%B4%E7%BA%A7")
	r.Header.Add("referer", url.QueryEscape(link))
	r.Header.Add("sec-fetch-dest", "document")
	r.Header.Add("sec-fetch-mode", "navigate")
	r.Header.Add("sec-fetch-site", "same-origin")
	r.Header.Add("sec-fetch-user", "?1")
	r.Header.Add("upgrade-insecure-requests", "1")
	r.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Mobile Safari/537.36")

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}
	return doc
}
