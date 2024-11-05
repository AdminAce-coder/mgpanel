package file

import (
	"mgpanel/internal/consts"
	w "mgpanel/internal/tasks/k3s/ec2/ec2workspace"
	"testing"
)

func TestOs(t *testing.T) {
	cf := w.CreateTf{}
	cf.GetWorkspaceAbs().Creatfile(consts.AwsEC2createOut)
}
