package etcdv3

import (
	"strings"

	"go.etcd.io/etcd/etcdserver/etcdserverpb"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

// Node 需要使用到的模型
type Node struct {
	IsDir   bool   `json:"is_dir"`
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
	}
}
