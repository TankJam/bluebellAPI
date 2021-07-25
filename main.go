package main

import (
	"bluebellAPI/controller"
	"bluebellAPI/dao/mysql"
	"bluebellAPI/dao/redis"
	"bluebellAPI/logger"
	"bluebellAPI/pkg/snowflake"
	"bluebellAPI/router"
	"bluebellAPI/settings"
	"fmt"
	"os"
)

func main() {
	// 启动程序时，必须带上配置文件参数
	if len(os.Args) < 2 {
		fmt.Println("need config file (bluebell config file)")
		return
	}

	// 1、加载配置
	/// 1) 初始化全局配置 settings.Init(os.Args[1])
	if err := settings.Init(os.Args[1]); err != nil {
		fmt.Println("load config file failed, err: ", err)
		return
	}

	/// 2) 初始化日志配置
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Println("init logger failed, err: ", err)
		return
	}

	/// 3) 初始化mysql配置
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	/// 4) 初始化redis配置
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis falied, err: %v\n", err)
		return
	}
	defer redis.Close()

	/// 5) 初始化雪花算法ID
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err: %v\n", err)
		return
	}

	/// 6) 初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err: %v\n", err)
		return
	}

	/// 7) 注册路由
	r := router.SetupRouter(settings.Conf.Mode)
	/// 8) 启动gin服务
	err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err: %v\n", err)
		return
	}

}
