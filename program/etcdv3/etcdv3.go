package etcdv3

import (
	"errors"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Etcd3Client etcd v3客户端
type Etcd3Client struct {
	*clientv3.Client
}

const (
	DEFAULT_DIR_VALUE = "etcdv3_dir_$2H#%gRe3*t"
)

var (
	// EtcdClis etcd连接对象
	etcdClis *sync.Map
)

func init() {
	etcdClis = new(sync.Map)

}

// NewEtcdCli 创建一个etcd客户端
func NewEtcdCli(name string, addr ...string) (*Etcd3Client, error) {
	if len(addr) == 0 {
		return nil, errors.New("Etcd connection address cannot be empty")
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	etcdClis.Store(name, cli)
	return &Etcd3Client{
		Client: cli,
	}, nil
}

// GetEtcdCli 获取一个etcd cli对象
func GetEtcdCli(name string, addr ...string) (*Etcd3Client, error) {
	val, ok := etcdClis.Load(name)
	if ok == false {
		if len(addr) > 0 {
			cli, err := NewEtcdCli(name, addr...)
			if err != nil {
				return nil, err
			}
			return cli, nil
		}
		return nil, errors.New("Getting etcd client error")
	}
	return &Etcd3Client{
		Client: val.(*clientv3.Client),
	}, nil
}
