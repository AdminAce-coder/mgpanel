// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package k8s

import (
	"context"

	"mgpanel/api/k8s/v1"
)

type IK8SV1 interface {
	CreateDp(ctx context.Context, req *v1.CreateDpReq) (res *v1.CreateDpRes, err error)
}
