package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.zhipin.com/")
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
	fmt.Printf("%s", all) // %s输出为字符串格式

}
