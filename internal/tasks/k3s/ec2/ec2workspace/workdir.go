package ec2workspace

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

type CreateTf struct {
	CreatefilePath string
}

// 获取当前文件所在的目录路径
func (c *CreateTf) GetWorkspaceAbs() *CreateTf {
	// 获取当前文件的绝对路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		glog.New().Error(gctx.New(), "无法获取当前文件路径")
		return c
	}
	// 获取文件所在目录
	dir := filepath.Dir(filename)
	c.CreatefilePath = dir
	glog.New().Info(gctx.New(), "工作路径为：", c.CreatefilePath)
	return c
}

// 创建文件
func (c *CreateTf) Creatfile(filename string) *os.File {
	ctx := gctx.New()

	c.CreatefilePath = filepath.Join(c.CreatefilePath, filename)
	glog.New().Info(ctx, "模板输出文件的完整路径为：", c.CreatefilePath)
	// 尝试创建文件
	f, err := os.Create(c.CreatefilePath)
	if err != nil {
		// 如果错误不是因为文件已存在，则记录错误
		if !os.IsExist(err) {
			glog.New().Error(ctx, "文件创建失败:", err)
			return nil
		} else {
			glog.New().Error(ctx, "文件已存在")
			return nil
		}
	}
	glog.New().Info(ctx, "文件创建成功", c.CreatefilePath)
	return f
}
