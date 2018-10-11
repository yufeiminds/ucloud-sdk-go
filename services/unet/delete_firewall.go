//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet DeleteFirewall

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DeleteFirewallRequest is request schema for DeleteFirewall action
type DeleteFirewallRequest struct {
	request.CommonBase

	// 防火墙资源ID
	FWId *string `required:"true"`
}

// DeleteFirewallResponse is response schema for DeleteFirewall action
type DeleteFirewallResponse struct {
	response.CommonBase
}

// NewDeleteFirewallRequest will create request of DeleteFirewall action.
func (c *UNetClient) NewDeleteFirewallRequest() *DeleteFirewallRequest {
	req := &DeleteFirewallRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DeleteFirewall - 删除防火墙
func (c *UNetClient) DeleteFirewall(req *DeleteFirewallRequest) (*DeleteFirewallResponse, error) {
	var err error
	var res DeleteFirewallResponse

	err = c.client.InvokeAction("DeleteFirewall", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}