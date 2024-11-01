package file

import (
	"os"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

// Creatfile 创建文件
func Creatfile(name string) *os.File {
	ctx := gctx.New()

	// 尝试创建文件
	f, err := os.Create(name)
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

	return f // 返回成功创建的文件句柄
}
