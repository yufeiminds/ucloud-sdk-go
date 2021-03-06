//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDisk CheckUDiskAllowance

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CheckUDiskAllowanceRequest is request schema for CheckUDiskAllowance action
type CheckUDiskAllowanceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 购买UDisk大小,单位:GB,范围[1~2000], 权限位控制可达8T,若需要请申请开通相关权限。
	Size *int `required:"true"`

	// 资源申请个数,默认为一个
	Count *int `required:"false"`
}

// CheckUDiskAllowanceResponse is response schema for CheckUDiskAllowance action
type CheckUDiskAllowanceResponse struct {
	response.CommonBase

	// 资源核查部分成功情况下，成功个数
	Count int
}

// NewCheckUDiskAllowanceRequest will create request of CheckUDiskAllowance action.
func (c *UDiskClient) NewCheckUDiskAllowanceRequest() *CheckUDiskAllowanceRequest {
	req := &CheckUDiskAllowanceRequest{}
	c.client.SetupRequest(req)
	return req
}

// CheckUDiskAllowance - 检查UDisk资源余量
func (c *UDiskClient) CheckUDiskAllowance(req *CheckUDiskAllowanceRequest) (*CheckUDiskAllowanceResponse, error) {
	var err error
	var res CheckUDiskAllowanceResponse

	err = c.client.InvokeAction("CheckUDiskAllowance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
