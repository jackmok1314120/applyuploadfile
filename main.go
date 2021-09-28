package main

import (
	"applyUpLoadFile/config"
	"applyUpLoadFile/middleware/watch"
	"applyUpLoadFile/routers"
	_ "applyUpLoadFile/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	cfg := config.Cfg
	router.MaxMultipartMemory = cfg.UpLoad.MaxFile << 20 // 8 MiB
	r := gin.Default()
	r.Use(routers.Cors())
	routers.InitRouter(r, watch.WatchTag)
	r.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
}
