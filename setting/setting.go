package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var GlobalConfig = new(Config)

func Init() (err error) {

	// 配置文件名称
	viper.SetConfigName("config")
	// 配置文件扩展名
	viper.SetConfigType("yaml")
	// 配置文件所在路径
	viper.AddConfigPath("./config")
	// 查找并读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		// 处理读取配置文件的错误
		return
	}

	// 配置信息绑定到结构体变量
	err = viper.Unmarshal(GlobalConfig)
	if err != nil {
		fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(c fsnotify.Event) {
		err = viper.ReadInConfig()
		if err != nil {
			// 处理读取配置文件的错误
			fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
			return
		}

		// 配置信息绑定到结构体变量
		err = viper.Unmarshal(GlobalConfig)
		if err != nil {
			fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
		}

		fmt.Println("检测到配置文件有变动,已实时加载")
	})
	return
}
