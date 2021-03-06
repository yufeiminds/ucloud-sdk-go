//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDisk DetachUDisk

package udisk

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DetachUDiskRequest is request schema for DetachUDisk action
type DetachUDiskRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// UHost实例ID
	UHostId *string `required:"true"`

	// 需要卸载的UDisk实例ID
	UDiskId *string `required:"true"`
}

// DetachUDiskResponse is response schema for DetachUDisk action
type DetachUDiskResponse struct {
	response.CommonBase

	// 卸载的UHost实例ID
	UHostId string

	// 卸载的UDisk实例ID
	UDiskId string
}

// NewDetachUDiskRequest will create request of DetachUDisk action.
func (c *UDiskClient) NewDetachUDiskRequest() *DetachUDiskRequest {
	req := &DetachUDiskRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DetachUDisk - 卸载某个已经挂载在指定UHost实例上的UDisk
func (c *UDiskClient) DetachUDisk(req *DetachUDiskRequest) (*DetachUDiskResponse, error) {
	var err error
	var res DetachUDiskResponse

	err = c.client.InvokeAction("DetachUDisk", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
