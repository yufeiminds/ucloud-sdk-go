//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UNet GetEIPUpgradePrice

package unet

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// GetEIPUpgradePriceRequest is request schema for GetEIPUpgradePrice action
type GetEIPUpgradePriceRequest struct {
	request.CommonBase

	// 弹性IP的资源ID
	EIPId *string `required:"true"`

	// 弹性IP的外网带宽, 单位为Mbps, 范围 [1-800]
	Bandwidth *int `required:"true"`
}

// GetEIPUpgradePriceResponse is response schema for GetEIPUpgradePrice action
type GetEIPUpgradePriceResponse struct {
	response.CommonBase

	// 调整带宽后的EIP价格, 单位为"元", 如需退费此处为负值
	Price float64
}

// NewGetEIPUpgradePriceRequest will create request of GetEIPUpgradePrice action.
func (c *UNetClient) NewGetEIPUpgradePriceRequest() *GetEIPUpgradePriceRequest {
	req := &GetEIPUpgradePriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// GetEIPUpgradePrice - 获取弹性IP带宽改动价格
func (c *UNetClient) GetEIPUpgradePrice(req *GetEIPUpgradePriceRequest) (*GetEIPUpgradePriceResponse, error) {
	var err error
	var res GetEIPUpgradePriceResponse

	err = c.client.InvokeAction("GetEIPUpgradePrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
