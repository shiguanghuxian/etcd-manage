package etcdv3

import (
	"errors"
	"sync"
	"time"

	"github.com/shiguanghuxian/etcd-manage/program/config"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/pkg/transport"
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
func NewEtcdCli(etcdCfg *config.EtcdServer) (*Etcd3Client, error) {
	if etcdCfg == nil {
		return nil, errors.New("etcdCfg is nil")
	}
	if etcdCfg.TLSEnable == true && etcdCfg.TLSConfig == nil {
		return nil, errors.New("TLSConfig is nil")
	}
	if len(etcdCfg.Address) == 0 {
		return nil, errors.New("Etcd connection address cannot be empty")
	}

	var cli *clientv3.Client
	var err error

	if etcdCfg.TLSEnable == true {
		// tls 配置
		tlsInfo := transport.TLSInfo{
			CertFile:      etcdCfg.TLSConfig.CertFile,
			KeyFile:       etcdCfg.TLSConfig.KeyFile,
			TrustedCAFile: etcdCfg.TLSConfig.CAFile,
		}
		tlsConfig, err := tlsInfo.ClientConfig()
		if err != nil {
			return nil, err
		}

		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   etcdCfg.Address,
			DialTimeout: 10 * time.Second,
			TLS:         tlsConfig,
			Username:    etcdCfg.Username,
			Password:    etcdCfg.Password,
		})
	} else {
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   etcdCfg.Address,
			DialTimeout: 10 * time.Second,
			Username:    etcdCfg.Username,
			Password:    etcdCfg.Password,
		})
	}

	if err != nil {
		return nil, err
	}
	etcdClis.Store(etcdCfg.Name, cli)
	return &Etcd3Client{
		Client: cli,
	}, nil
}

// Close 关闭连接
func (c *Etcd3Client) Close() error {
	return c.Client.Close()
}

// GetEtcdCli 获取一个etcd cli对象
func GetEtcdCli(etcdCfg *config.EtcdServer) (*Etcd3Client, error) {
	if etcdCfg == nil {
		return nil, errors.New("etcdCfg is nil")
	}
	if len(etcdCfg.Address) > 0 {
		cli, err := NewEtcdCli(etcdCfg)
		if err != nil {
			return nil, err
		}
		return cli, nil
	}
	return nil, errors.New("Getting etcd client error")

	// val, ok := etcdClis.Load(etcdCfg.Name)
	// if ok == false {
	// 	if len(etcdCfg.Address) > 0 {
	// 		cli, err := NewEtcdCli(etcdCfg)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return cli, nil
	// 	}
	// 	return nil, errors.New("Getting etcd client error")
	// }
	// return &Etcd3Client{
	// 	Client: val.(*clientv3.Client),
	// }, nil
}
