package etcdv3

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

// Node 需要使用到的模型
type Node struct {
	IsDir   bool   `json:"is_dir"`
	Version int64  `json:"version"`
	Value   string `json:"value"`
	FullDir string `json:"full_dir"`
}

const (
	ROLE_LEADER   = "leader"
	ROLE_FOLLOWER = "follower"

	STATUS_HEALTHY   = "healthy"
	STATUS_UNHEALTHY = "unhealthy"
)

// Member 节点信息
type Member struct {
	*etcdserverpb.Member
	Role   string `json:"role"`
	Status string `json:"status"`
	DbSize int64  `json:"db_size"`
}

// NewNode 创建节点
func NewNode(dir string, kv *mvccpb.KeyValue) *Node {
	return &Node{
		IsDir:   string(kv.Value) == DEFAULT_DIR_VALUE,
		Value:   strings.TrimPrefix(string(kv.Key), dir),
		FullDir: string(kv.Key),
		Version: kv.Version,
	}
}

// node 列表格式化成json
func NodeJsonFormat(prefix string, list []*Node) (interface{}, error) {
	resp := make(map[string]interface{}, 0)
	if len(list) == 0 {
		return resp, nil
	}
	for _, v := range list {
		key := strings.TrimPrefix(v.FullDir, prefix)
		key = strings.TrimLeft(key, "/")
		strs := strings.Split(key, "/")
		if len(strs) > 0 {
			mp := resp
			for _, val := range strs {
				log.Println(val)
				js, _ := json.Marshal(mp)
				log.Println(string(js))
				if _, ok := mp[val]; ok == false {
					if v.Value == DEFAULT_DIR_VALUE {
						mp[val] = make(map[string]interface{}, 0)
					} else {
						mp[val] = formatValue(v.Value) // 这里应该做个类型预判
					}
				}
				var ok bool
				subMp, ok := mp[val].(map[string]interface{})
				if ok == false {
					break
				}
				mp = subMp
			}
			resp = mp
		}
	}
	return resp, nil
}

// Format 时获取值，转为指定类型
func formatValue(v string) interface{} {
	if v == "true" {
		return true
	} else if v == "false" {
		return false
	}
	// 尝试转浮点数
	vf, err := strconv.ParseFloat(v, 64)
	if err == nil {
		return vf
	}
	// 尝试转整数
	vi, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		return vi
	}
	return v
}
