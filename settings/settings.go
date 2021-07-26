package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 初始化 AppConfig 结构体
var Conf = new(AppConfig)

// AppConfig 配置结构体
type AppConfig struct {
	Name      string `mapstructure:"name"`  // 注意 : 左右不能有空格
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	// 继承父结构体 mysql、redis、log 配置
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

// LogConfig
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// MySQLConfig
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisConfig
type RedisConfig struct {
	Host         string `mapstructure:"host"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"Min_Idle_conns"`
}

// Init 初始化加载配置信息
func Init(filePath string) (err error) {
	// 1.使用viper热加载配置文件,指定yaml格式的文件
	// yiper.SetConfigType("json") // 指定json格式
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err: %v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("Viper.Unmarshal failed, err: %v\n", err)
	}

	viper.WatchConfig() // 监听配置
	// 检测如果配置文件若修改了，则重新反序列化更新配置
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了，马上重新加载配置信息...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("Viper.Unmarshal failed, err: %v\n", err)
		}
	})
	return
}
