package config

import (
	"os"

	"github.com/naoina/toml"
	"github.com/shiguanghuxian/etcd-manage/program/common"
)

// Config 配置
type Config struct {
	Debug   bool          `toml:"debug"`
	LogPath string        `toml:"log_path"`
	HTTP    *HTTP         `toml:"http"`
	Server  []*EtcdServer `toml:"server"`
	Users   []*User       `toml:"user"`
}

// GetUserByUsername 根据用户名获取用户信息
func (c *Config) GetUserByUsername(username string) *User {
	if c.Users != nil && len(c.Users) > 0 {
		for _, u := range c.Users {
			if u.Username == username {
				return u
			}
		}
	}
	return nil
}

// HTTP http 件套配置
type HTTP struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

// EtcdServer etcd 服务
type EtcdServer struct {
	Title     string   `toml:"title"`
	Name      string   `toml:"name"`
	Address   []string `toml:"address"`
	KeyPrefix string   `toml:"key_prefix"`
	Desc      string   `toml:"desc"`
	Roles     []string `toml:"roles"` // 可访问此etcd服务的角色列表
}

// User 用户
type User struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Role     string `toml:"role"`
}

// GetEtcdServer 根据服务标识查找服务
func GetEtcdServer(name string) *EtcdServer {
	if cfg == nil {
		return nil
	}
	for _, v := range cfg.Server {
		if v.Name == name {
			return v
		}
	}
	return nil
}

var (
	cfg *Config
)

// GetCfg 获取配置
func GetCfg() *Config {
	if cfg == nil {
		LoadConfig("")
	}
	return cfg
}

// LoadConfig 读取配置
func LoadConfig(cfgPath string) (*Config, error) {
	cfgPath = getCfgPath(cfgPath)
	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cfg = new(Config)
	if err := toml.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func getCfgPath(cfgPath string) string {
	if cfgPath == "" {
		cfgPath = common.GetRootDir() + "config/cfg.toml"
	}
	return cfgPath
}
