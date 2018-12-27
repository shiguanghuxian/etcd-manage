package program

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	gin "github.com/gin-gonic/gin"
	"github.com/shiguanghuxian/etcd-manage/program/config"
	"github.com/shiguanghuxian/etcd-manage/program/etcdv3"
	"github.com/shiguanghuxian/etcd-manage/program/v1"
)

// http 服务

// 启动http服务
func (p *Program) startAPI() {
	router := gin.Default()

	// 设置静态文件目录
	router.GET("/ui/*w", p.handlerStatic)
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/ui")
	})

	// 读取认证列表
	accounts := make(gin.Accounts, 0)
	if p.cfg.Users != nil {
		for _, u := range p.cfg.Users {
			accounts[u.Username] = u.Password
		}
	}

	// v1 api
	apiV1 := router.Group("/v1", gin.BasicAuth(accounts))
	apiV1.Use(p.middleware()) // 只有接口需要中间件
	v1.V1(apiV1)

	addr := fmt.Sprintf("%s:%d", p.cfg.HTTP.Address, p.cfg.HTTP.Port)
	// 监听
	s := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("启动HTTP服务:", addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

// 跨域中间件
func (p *Program) middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取登录用户名，查询角色
		userIn := c.MustGet(gin.AuthUserKey)
		userRole := ""
		if userIn != nil {
			user := userIn.(string)
			if user == "" {
				c.Set("userRole", "")
			} else {
				u := p.cfg.GetUserByUsername(user)
				if u == nil {
					c.Set("userRole", "")
				} else {
					userRole = u.Role
					// 角色和用户信息
					c.Set("userRole", u.Role)
					c.Set("authUser", u)
				}
			}
		}

		// 绑定etcd连接
		etcdServerName := c.GetHeader("EtcdServerName")
		if etcdServerName != "" {
			cli, s, err := getEtcdCli(etcdServerName, userRole)
			if err == nil {
				c.Set("EtcdServer", cli)
				c.Set("EtcdServerCfg", s)
			}
		}

		c.Next()

		// gin设置响应头，设置跨域
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Access-Control-Allow-Origin")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		//放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
		}
	}
}

// 获取etcd 连接
func getEtcdCli(name, role string) (ctl *etcdv3.Etcd3Client, s *config.EtcdServer, err error) {
	s = config.GetEtcdServer(name)
	if s == nil {
		return nil, nil, errors.New("etcd服务不存在")
	}
	if len(s.Roles) > 0 {
		isRole := false
		for _, r := range s.Roles {
			if role == r {
				isRole = true
				break
			}
		}
		if isRole == false {
			return nil, nil, errors.New("无权限访问")
		}
	}
	ctl, err = etcdv3.GetEtcdCli(name, s.Address...)
	return
}
