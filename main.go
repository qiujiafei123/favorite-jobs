package main

import (
	"favorite-jobs/cache"
	"favorite-jobs/config"
	"favorite-jobs/crawler"
	"favorite-jobs/log"
	"favorite-jobs/model"
	"favorite-jobs/parser"
	"fmt"
	"sync"
)

func main() {
	Init()
	// 开始抓取
	group := sync.WaitGroup{}
	group.Add(1)
	go func() {
		for i := 1; i <= 10; i++ {
			url := fmt.Sprintf("https://www.lagou.com/zhaopin/PHP/%d/?filterOption=%d&sid=503571df8f5d42b488326d4945572410", i, i)
			fmt.Printf("正在访问: %s\n", url)
			crawler.Visit(url, "#s_position_list ul .con_list_item", parser.JobIndex)
		}
		group.Done()
	}()
	// 启动 model 写入协程
	go cache.Revice(&group)
	group.Wait()
}

// 启动前需要初始化的func
func Init() {
	config.LoadConfig()
	log.InitLog()
	cache.InitCache()
	cache.InitModelChan()
	model.InitDB()
}