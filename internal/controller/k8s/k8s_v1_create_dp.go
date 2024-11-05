package k8s

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"mgpanel/api/k8s/v1"
)

func (c *ControllerV1) CreateDp(ctx context.Context, req *v1.CreateDpReq) (res *v1.CreateDpRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
