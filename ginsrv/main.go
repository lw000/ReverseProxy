// test project main.go
package main

import (
	"demo/ReverseProxy/ginsrv/config"
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	// "github.com/gin-contrib/cors"
	// "github.com/gin-contrib/cache"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods",
				"POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers",
				"Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers",
				"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")          // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")             // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523433"},
}

func runSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-c

		signal.Stop(c)

		close(c)

		os.Exit(1)
	}()
}

func main() {
	runSignal()

	engine := gin.Default()
	engine.Use(Cors())

	limiter := tollbooth.NewLimiter(1000, nil)
	limiter.SetMethods([]string{"GET", "POST"})
	limiter.SetMessage("休息下吧")

	engine.GET("/tollbooth", tollbooth_gin.LimitHandler(limiter), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": 1})
	})

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": 1})
	})

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"timestamp": time.Now().Unix()})
	})

	engine.GET("/add", func(c *gin.Context) {
		a := c.Query("a")
		b := c.Query("b")

		ia, err := strconv.Atoi(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		ib, err := strconv.Atoi(b)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": ia + ib})
	})

	engine.POST("/add", func(c *gin.Context) {
		a := c.PostForm("a")
		b := c.PostForm("b")

		ia, err := strconv.Atoi(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		ib, err := strconv.Atoi(b)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": ia + ib})
	})

	engine.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	v1 := engine.Group("/v1")
	v1.GET("/sub", func(c *gin.Context) {
		a := c.DefaultQuery("a", "")
		b := c.DefaultQuery("b", "")

		ia, err := strconv.Atoi(a)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		ib, err := strconv.Atoi(b)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": ia - ib})
	})

	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foot":   "bar",
		"austin": "123456",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "no secret :("})
		}
	})

	cfg := config.NewConfig()
	if er := cfg.Load("conf/conf.json"); er != nil {
		return
	}

	log.Panic(engine.Run(fmt.Sprintf(":%d", cfg.Port)))
}
