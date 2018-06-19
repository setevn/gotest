package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])

	}
	return
}
func SpiderOneJoy(url string) (title, content string, err error) {
	result, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	//取关键信息
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	tmpTitle := re1.FindAllStringSubmatch(result, 1)
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1)
		break
	}
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err2")
		return
	}
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		break
	}
	return
}
func StoreJoyTofile(i int, fileTile, fileContent []string) {
	//新建文件
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("%os.create err = ", err)
		return
	}
	defer f.Close()

	n := len(fileTile)
	for i := 0; i < n; i++ {
		f.WriteString(fileTile[i] + "\r\n")
		f.WriteString(fileContent[i] + "\r\n")
		f.WriteString("\r\n=======================\r\n")
	}
}
func SpiderPage(i int, page chan int) {
	url := "https://www.pengfu.com/index_" + strconv.Itoa(i) + ".html"
	fmt.Println("正在爬取第%d页网页%s:", i, url)
	//2爬
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("httpget err = ", err)
		return
	}
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println(" regexp.MustCompile err ")
		return
	}
	joyUrls := re.FindAllStringSubmatch(result, -1)
	//fmt.Println("joyUrls = ", joyUrls)
	fileTile := make([]string, 0)
	fileContent := make([]string, 0)
	for _, data := range joyUrls {
		//fmt.Println("url = ", data[1])
		//开始爬取每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println(" SpiderOneJoy err =", err)
			continue
		}
		//fmt.Printf("title = #%v#", title)
		//fmt.Printf("content = #%v#", content)
		fileTile = append(fileTile, title)
		fileContent = append(fileContent, content)
	}
	//把内容写入文件
	StoreJoyTofile(i, fileTile, fileContent)
	page <- i
}
func Dowork(start, end int) {
	fmt.Printf("正在爬取%d 到 %d 的页面\n", start, end)

	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}
func main() {
	var start, end int
	fmt.Printf("请输入起始页（>=1）:")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页（>=起始页）")
	fmt.Scan(&end)

	Dowork(start, end)
}
