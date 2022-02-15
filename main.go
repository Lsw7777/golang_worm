package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func printCity(contents []byte) { // []byte就是字符串
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`) //  Compile用来解析正则表达式是否合法，如果合法，则返回一个Regexp对象
	matches := re.FindAllSubmatch(contents, -1)                                                       // 返回一个切片，第一个参数是要进行正则处理的数据，第二个参数小于0，表示返回所有匹配的内容
	fmt.Println(matches)
	for _, v := range matches { // 遍历打印所有结果
		fmt.Printf("city:%s,URL:%s\n", v[2], v[1]) // 从html元素中截取city和URL，其它的内容忽略
	}
	fmt.Printf("%d\n", len(matches)) // %d使用十进制格式显示
}

func main() {
	res, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // res要在检查错误后再使用，如果err报错那res就无法获取到Body，也就无法使用了

	if res.StatusCode != http.StatusOK {
		fmt.Println("状态码出错")
		return
	}
	all, err := ioutil.ReadAll(res.Body) // ioutil.ReadAll从一个io.Reader接口参数中一次性读取所有数据
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%s", all) // %s输出为字符串格式
	printCity(all)

}
