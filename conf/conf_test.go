package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

//配置文件中字母要小写，结构体属性首字母要大写

func TestConfigYamlAll(t *testing.T) {
	data, _ := ioutil.ReadFile("config.yaml")
	fmt.Println(string(data))
	c := Config{}
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &c)
	fmt.Println("初始数据", c)

	d, _ := yaml.Marshal(&c)
	fmt.Println("看看 :", string(d))
	fmt.Println(c.Server.AllowIp)
}
