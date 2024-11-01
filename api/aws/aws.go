// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package aws

import (
	"context"

	"mgpanel/api/aws/v1"
)

type IAwsV1 interface {
	CreateEc2(ctx context.Context, req *v1.CreateEc2Req) (res *v1.CreateEc2Res, err error)
}
