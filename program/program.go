package program

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/shiguanghuxian/etcd-manage/program/config"
	"github.com/shiguanghuxian/etcd-manage/program/logger"
)

// Program 主程序
type Program struct {
	cfg *config.Config
	s   *http.Server
}

// New 创建主程序
func New() (*Program, error) {
	// 配置文件
	cfg, err := config.LoadConfig("")
	if err != nil {
		return nil, err
	}

	// 日志对象
	_, err = logger.InitLogger(cfg.LogPath, cfg.Debug)
	if err != nil {
		return nil, err
	}

	return &Program{
		cfg: cfg,
	}, nil
}

// Run 启动程序
func (p *Program) Run() error {
	// 启动http服务
	go p.startAPI()

	// 打开浏览器
	go func() {
		time.Sleep(100 * time.Millisecond)
		cmd := exec.Command("open", fmt.Sprintf("http://127.0.0.1:%d/ui/", p.cfg.HTTP.Port))
		cmd.Run()
	}()

	return nil
}

// Stop 停止服务
func (p *Program) Stop() {
	if p.s != nil {
		p.s.Close()
	}
}

// 重新加载配置文件
// 使用go-cache
