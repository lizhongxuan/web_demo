package main

import (
	"fmt"
	"ginDemo/router"
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

func initConfig(file string) error {

	//
	// 读取配置文件
	//
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
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
	err := initConfig("./conf/ginDemo.conf")
	if err != nil {
		fmt.Println("init config faild: %v", err.Error())
		os.Exit(2)
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
	http.ListenAndServe(App.Port, g)
}
