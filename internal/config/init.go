package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

func init() {
	file := determineFile()
	loadFromFile(file)
}

func loadFromFile(file string) {
	viper.AddConfigPath("./")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../../configs")
	viper.SetConfigType("yml")

	if strings.ContainsRune(file, filepath.Separator) {
		viper.SetConfigFile(file)
	} else {
		viper.SetConfigName(file)
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("Used config file:", viper.ConfigFileUsed())
	//f, _ := ioutil.ReadFile(viper.ConfigFileUsed())
	//err = yaml.Unmarshal(f, appConf)

	err = viper.Unmarshal(appConf)
	if err != nil {
		panic(err)
	}
	fmt.Println("appConf:", appConf)

}

// viper优先级：配置文件>命令行标志位>环境变量>远程Key/Value存储>默认值
func determineFile() string {
	// flags
	pflag.StringP("env", "e", "dev", "app runtime env")
	pflag.StringP("file", "f", "", "app config file")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}

	// envs
	err = viper.BindEnv("env", "APP_ENV")
	if err != nil {
		panic(err)
	}
	err = viper.BindEnv("file", "APP_CONF_FILE")
	if err != nil {
		panic(err)
	}

	var result string
	file := viper.GetString("file")
	env := viper.GetString("env")
	if strings.TrimSpace(file) != "" {
		fmt.Println("specified config file :", file)
		result = file
	} else {
		fmt.Println("detected environment:", env)
		result = fmt.Sprintf("app-%s.yml", env)
	}
	return result
}
