package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

//配置文件中字母要小写，结构体属性首字母要大写

type Myconf struct {
	Ipport             string
	StartSendTime      string
	SendMaxCountPerDay int
	Devices            []Device
	WarnFrequency      int
	SendFrequency      int
}
type Device struct {
	DevId string
	Nodes []Node
}
type Node struct {
	PkId     string
	BkId     string
	Index    string
	MinValue float32
	MaxValue float32
	DataType string
}

func TestConfigYamlAll(t *testing.T) {
	data, _ := ioutil.ReadFile("config.yaml")
	fmt.Println(string(data))
	c := Myconf{}
	//把yaml形式的字符串解析成struct类型
	yaml.Unmarshal(data, &c)
	fmt.Println("初始数据", c)
	if c.Ipport == "" {
		fmt.Println("配置文件设置错误")
		return
	}
	d, _ := yaml.Marshal(&c)
	fmt.Println("看看 :", string(d))
	fmt.Println(c.Devices[1].DevId)
}
