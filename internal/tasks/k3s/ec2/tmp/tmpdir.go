package tmp

import (
	"path/filepath"
	"runtime"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

//type Tmpdir struct {
//}

// 获取当前文件所在的目录路径
func GetTmpDir() string {
	// 获取当前文件的绝对路径
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		glog.New().Error(gctx.New(), "无法获取当前文件路径")
		return " "
	}
	// 获取文件所在目录
	dir := filepath.Dir(filename)
	glog.New().Info(gctx.New(), "Tmldir工作路径为：", dir)
	return dir
}
