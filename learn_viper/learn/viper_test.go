package learn

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"testing"
)

type config struct {
	Language string
	Host     string
	Port     int
	People   Person `mapstructure:"Person"`
}

type Person struct {
	Name   string
	Age    int
	Height int `mapstructure:"h"`
}

type Family struct {
	Address string
}

type Social struct {
	Wife string
	Know []string
}

var v = viper.New()

func TestReadConfig(t *testing.T) {
	// 设置配置文件名，不需要带扩展名，便于在不修改代码的情况下替换配置文件的类型
	v.SetConfigName("configfile")
	// 设置配置文件类型(可选，viper 会自动判断)
	v.SetConfigType("toml")
	// 可以添加多个配置文件的路径
	v.AddConfigPath("../")
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		fill()
	})
	fill()
}

func fill() {
	// 读取配置文件，会根据不同的文件类型调用不同的解析库进行解析
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("Can not find config file"))
		} else {
			panic(fmt.Errorf("Fatal error config file: %s", err))
		}
	}
	// 读取配置文件
	language := v.GetString("language")
	age := v.GetInt("person.age")
	fmt.Printf("language = %d, language = %s\n", age, language)
	// 将从配置文件中的读取到的配置信息填充到结构体中(含内嵌的结构体！)
	var conf config
	if err := v.Unmarshal(&conf); err != nil {
		panic("error 1")
	}
	fmt.Printf("%v\n", conf)
	var family Family
	if err := v.UnmarshalKey("family", &family); err != nil {
		panic("error 2")
	}
	fmt.Printf("%v\n", family)
}

// 填充至多个结构体
func TestStructures(t *testing.T) {
	v.AddConfigPath("../")
	v.SetConfigName("info")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	var person Person
	var social Social
	if err := v.UnmarshalKey("person", &person); err != nil {
		panic("error 1")
	}
	t.Logf("%v", person)
	if err := v.UnmarshalKey("social", &social); err != nil {
		panic("error 2")
	}
	t.Logf("%v", social)
}
