package etcdv3

import (
	"strconv"
	"strings"

	etcdserverpb "github.com/coreos/etcd/etcdserver/etcdserverpb"
	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
)

// Node 需要使用到的模型
type Node struct {
	IsDir   bool   `json:"is_dir"`
	Version int64  `json:"version,string"`
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
		// if len(strs) > 0 {
		// 	for _, val := range strs {
		// 		log.Println(val)
		// 		js, _ := json.Marshal(resp)
		// 		log.Println(string(js))
		// 		if _, ok := resp[val]; ok == false {
		// 			if v.Value == DEFAULT_DIR_VALUE {
		// 				resp[val] = make(map[string]interface{}, 0)
		// 			} else {
		// 				resp[val] = formatValue(v.Value) // 这里应该做个类型预判
		// 			}
		// 		}
		// 		_, ok := resp[val].(map[string]interface{})
		// 		if ok == false {
		// 			break
		// 		}
		// 	}
		// }
		// log.Println("---------------------")
		// log.Println(v.FullDir)
		// log.Println(strs)
		// log.Println(v.Value)

		recursiveJsonMap(strs, v, resp)
		// jjj, _ := json.Marshal(resp)
		// log.Println(string(jjj))

	}
	// jjj, _ := json.Marshal(resp)
	// log.Println(string(jjj))
	return resp, nil
}

// 递归的将一个值赋值到map中
func recursiveJsonMap(strs []string, node *Node, parent map[string]interface{}) interface{} {
	if len(strs) == 0 || strs[0] == "" || node == nil || parent == nil {
		return nil
	}
	if _, ok := parent[strs[0]]; ok == false {
		if node.Value == DEFAULT_DIR_VALUE {
			parent[strs[0]] = make(map[string]interface{}, 0)
		} else {
			parent[strs[0]] = formatValue(node.Value)
		}
	}
	val, ok := parent[strs[0]].(map[string]interface{})
	if ok == false {
		return val
	}
	return recursiveJsonMap(strs[1:], node, val)
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
