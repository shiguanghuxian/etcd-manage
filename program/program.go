package program

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
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

	// jj, _ := json.Marshal(cfg)
	// fmt.Println(string(jj))

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
		openURL(fmt.Sprintf("http://127.0.0.1:%d/ui/", p.cfg.HTTP.Port))
	}()

	return nil
}

// Stop 停止服务
func (p *Program) Stop() {
	if p.s != nil {
		p.s.Close()
	}
}

// 打开url
func openURL(urlAddr string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", " /c start "+urlAddr)
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", urlAddr)
	} else {
		return
	}
	err := cmd.Start()
	if err != nil {
		logger.Log.Errorw("Error opening browser", "err", err)
	}
}
