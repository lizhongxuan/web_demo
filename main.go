package main

import (
	"./router"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/naoina/toml"
	"io/ioutil"
	"net/http"
	"os"
)

var App AppConfig

type AppConfig struct {
	Runmode string
	Port    string
}

// 读取配置文件
func initConfig(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := toml.Unmarshal(buf, &App); err != nil {
		return err
	}
	return nil
}

func main() {
	err := initConfig("./conf/webdemo.conf")
	if err != nil {
		fmt.Println("init config faild: %v", err.Error())
		return
	}

	// Set gin mode.
	gin.SetMode(App.Runmode)

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}
	// Routes.
	router.Load(
		// Cores.
		g,
		// Middlwares.
		middlewares...,
	)
	fmt.Println(http.ListenAndServe(App.Port, g).Error())
}
