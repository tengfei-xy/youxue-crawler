# 优学课堂的批量抓取视频
作为第三方应用的链接：https://video.ijiwen.com/index.php?f=nb&code=12cf388c6e809509&state=5642#auth
# 说明
1. 以输出log为准,总计输出与通过迅雷发现的连接数并不一定一致，因为有重复。
2. mac的迅雷一次性只能接受1k条链接，windows至少3k条。
# golang依赖
`go get github.com/PuerkitoBio/goquery`
# 运行
`go run main.go`
# 更新情况
- 2020/08/13 创建并完成爬虫项目
- 2020/08/17 补充一点内容和上次log输出