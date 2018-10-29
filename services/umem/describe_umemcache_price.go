//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem DescribeUMemcachePrice

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUMemcachePriceRequest is request schema for DescribeUMemcachePrice action
type DescribeUMemcachePriceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 容量大小,单位:GB 取值范围[1-32]
	Size *int `required:"true"`

	// 计费模式，Year， Month， Dynamic，默认: Dynamic 默认: 获取所有计费模式的价格
	ChargeType *string `required:"false"`

	// 购买umemcache的时长，默认值为1
	Quantity *int `required:"false"`
}

// DescribeUMemcachePriceResponse is response schema for DescribeUMemcachePrice action
type DescribeUMemcachePriceResponse struct {
	response.CommonBase

	// 价格列表, 参见 UMemcachePriceSet
	DataSet []UMemcachePriceSet
}

// NewDescribeUMemcachePriceRequest will create request of DescribeUMemcachePrice action.
func (c *UMemClient) NewDescribeUMemcachePriceRequest() *DescribeUMemcachePriceRequest {
	req := &DescribeUMemcachePriceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUMemcachePrice - 获取umemcache组价格信息
func (c *UMemClient) DescribeUMemcachePrice(req *DescribeUMemcachePriceRequest) (*DescribeUMemcachePriceResponse, error) {
	var err error
	var res DescribeUMemcachePriceResponse

	err = c.client.InvokeAction("DescribeUMemcachePrice", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}