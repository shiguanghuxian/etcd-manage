package v1

import (
	"errors"
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
	"github.com/shiguanghuxian/etcd-manage/program/config"
	"github.com/shiguanghuxian/etcd-manage/program/etcdv3"
)

// V1 v1 版接口
func V1(v1 *gin.RouterGroup) {
	v1.GET("/members", getEtcdMembers) // 获取节点列表
	v1.GET("/list", getEtcdKeyList)    // 获取目录下列表
	v1.GET("/key", getEtcdKeyValue)    // 获取一个key的具体值
	v1.POST("/key", postEtcdKey)       // 添加key
	v1.PUT("/key", putEtcdKey)         // 修改key
	v1.DELETE("/key", delEtcdKey)      // 删除key

	v1.GET("/server", getEtcdServerList) // 获取etcd服务列表

}

// 获取etcd值
func getEtcdKeyList(c *gin.Context) {
	key := c.Query("key")

	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()

	// log.Println(key)
	etcdCli, exists := c.Get("EtcdServer")
	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Etcd client is empty",
		})
		return
	}
	cli := etcdCli.(*etcdv3.Etcd3Client)

	resp, err := cli.List(key)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, resp)
}

// 获取key的值
func getEtcdKeyValue(c *gin.Context) {
	key := c.Query("key")
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()

	etcdCli, exists := c.Get("EtcdServer")
	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Etcd client is empty",
		})
		return
	}
	cli := etcdCli.(*etcdv3.Etcd3Client)

	val, err := cli.Value(key)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, val)
}

// 获取服务节点
func getEtcdMembers(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()

	etcdCli, exists := c.Get("EtcdServer")
	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Etcd client is empty",
		})
		return
	}
	cli := etcdCli.(*etcdv3.Etcd3Client)

	members, err := cli.Members()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, members)
}

// PostReq 添加和修改时的body
type PostReq struct {
	*etcdv3.Node
	EtcdName string `json:"etcd_name"`
}

// 添加key
func postEtcdKey(c *gin.Context) {
	saveEtcdKey(c, false)
}

// 修改key
func putEtcdKey(c *gin.Context) {
	saveEtcdKey(c, true)
}

// 删除key
func delEtcdKey(c *gin.Context) {
	key := c.Query("key")

	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()

	etcdCli, exists := c.Get("EtcdServer")
	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Etcd client is empty",
		})
		return
	}
	cli := etcdCli.(*etcdv3.Etcd3Client)

	err = cli.Delete(key)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// 保存key
func saveEtcdKey(c *gin.Context, isPut bool) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()

	req := new(PostReq)
	err = c.Bind(req)
	if err != nil {
		return
	}
	if req.FullDir == "" {
		err = errors.New("参数错误")
		return
	}

	etcdCli, exists := c.Get("EtcdServer")
	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Etcd client is empty",
		})
		return
	}
	cli := etcdCli.(*etcdv3.Etcd3Client)

	if req.IsDir == true {
		if isPut == true {
			err = errors.New("目录不能修改")
		} else {
			err = cli.Put(req.FullDir, etcdv3.DEFAULT_DIR_VALUE, !isPut)
		}
	} else {
		err = cli.Put(req.FullDir, req.Value, !isPut)
	}

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, "ok")
}

// 获取etcd服务列表
func getEtcdServerList(c *gin.Context) {
	cfg := config.GetCfg()
	if cfg == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "配置为nil",
		})
		return
	}
	list := cfg.Server
	if list == nil {
		list = make([]*config.EtcdServer, 0)
		c.JSON(http.StatusOK, list)
		return
	}
	// 当前用户角色
	userRole := ""
	userRoleIn, exists := c.Get("userRole")
	if exists == true {
		userRole = userRoleIn.(string)
	}

	log.Println(userRole)

	// 只返回有权限服务列表
	list1 := make([]*config.EtcdServer, 0)
	for _, s := range list {
		if s.Roles == nil || len(s.Roles) == 0 {
			list1 = append(list1, s)
		} else {
			for _, r := range s.Roles {
				if r == userRole {
					list1 = append(list1, s)
					break
				}
			}
		}
	}

	log.Println(list1)

	c.JSON(http.StatusOK, list1)
}
