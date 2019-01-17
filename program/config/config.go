package config

import (
	"errors"
	"os"
	"regexp"

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
	Address               string   `toml:"address"`
	Port                  int      `toml:"port"`
	TLSEnable             bool     `toml:"tls_enable"`               // 是否启用tls连接
	TLSConfig             *HTTPTls `toml:"tls_config"`               // 启用tls时必须配置此内容
	TLSEncryptEnable      bool     `toml:"tls_encrypt_enable"`       // 是否启用 Let's Encrypt tls
	TLSEncryptDomainNames []string `toml:"tls_encrypt_domain_names"` // 启用 Let's Encrypt 时的域名列表
}

// HTTPTls http tls配置
type HTTPTls struct {
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
}

// EtcdServer etcd 服务
type EtcdServer struct {
	Title     string         `toml:"title"`
	Name      string         `toml:"name"`
	Address   []string       `toml:"address"`
	Username  string         `toml:"username"`
	Password  string         `toml:"password"`
	KeyPrefix string         `toml:"key_prefix"`
	Desc      string         `toml:"desc"`
	TLSEnable bool           `toml:"tls_enable"` // 是否启用tls连接
	TLSConfig *EtcdTLSConfig `toml:"tls_config"` // 启用tls时必须配置此内容
	Roles     []string       `toml:"roles"`      // 可访问此etcd服务的角色列表
}

// EtcdTLSConfig etcd tls配置
type EtcdTLSConfig struct {
	CertFile string `toml:"cert_file"`
	KeyFile  string `toml:"key_file"`
	CAFile   string `toml:"ca_file"`
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

	// 验证服务name是否非字母和数字
	for _, v := range cfg.Server {
		if chackEtcdServerName(v.Name) == false {
			return nil, errors.New("etcd server name can only be letters or numbers or '_'")
		}
	}

	return cfg, nil
}

func getCfgPath(cfgPath string) string {
	if cfgPath == "" {
		cfgPath = common.GetRootDir() + "config/cfg.toml"
	}
	return cfgPath
}

// 判断etcd服务名是否包含非字母和数字
func chackEtcdServerName(name string) bool {
	if name == "" {
		return false
	}
	reg := regexp.MustCompile("[^0-9A-Za-z_]+")
	return !reg.MatchString(name)
}
