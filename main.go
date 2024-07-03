package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"path"
	"runtime"
	"time"

	"github.com/dazuimao1990/gin-demo/utils"
	"github.com/gin-gonic/gin"
)

func DowFile(c *gin.Context) {
	//通过动态路由方式获取文件名，以实现下载不同文件的功能
	name := c.Param("name")
	//拼接路径,如果没有这一步，则默认在当前路径下寻找
	filename := path.Join("./videos", name)
	//响应一个文件
	c.File(filename)
	// return
}

func main() {

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	IpString := utils.GetEniIP()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/whoami"
		r.HandleContext(c)
	})
	r.GET("/whoami", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time now": time.Now(),
			"message":  "Hello, I am " + IpString,
		})
	})
	r.GET("/sysinfo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"time now":     time.Now(),
			"Cpu Useage: ": fmt.Sprintf("%.2f", utils.GetCpuPercent()) + "%",
			"Mem Useage: ": fmt.Sprintf("%v", utils.GetMemTotal()) + "GB" + fmt.Sprintf("%.2f", utils.GetMemPercent()) + "%",
		})
	})

	// r.GET("/videos", func(c *gin.Context) {
	// 	c.File("./videos/index.html")
	// })

	// r.GET("/videos/:name", DowFile)
	r.Run(utils.GetEnvDefault("ADDRESS", "0.0.0.0") + ":" + utils.GetEnvDefault("PORT", "8080"))
}
