//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UDB UploadUDBParamGroup

package udb

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UploadUDBParamGroupRequest is request schema for UploadUDBParamGroup action
type UploadUDBParamGroupRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// DB类型id，DB类型id，mysql/mongodb/postgesql按版本细分 1：mysql-5.1，2：mysql-5.5，3：percona-5.5，4：mysql-5.6，5：percona-5.6，6：mysql-5.7，7：percona-5.7，8：mariadb-10.0，9：mongodb-2.4，10：mongodb-2.6，11：mongodb-3.0，12：mongodb-3.2,13：postgresql-9.4，14：postgresql-9.6
	DBTypeId *string `required:"true"`

	// 配置参数组名称
	GroupName *string `required:"true"`

	// 参数组描述
	Description *string `required:"true"`

	// 配置内容，导入的配置内容采用base64编码
	Content *string `required:"true"`

	// 该配置文件是否是地域级别配置文件，默认是false
	RegionFlag *bool `required:"false"`
}

// UploadUDBParamGroupResponse is response schema for UploadUDBParamGroup action
type UploadUDBParamGroupResponse struct {
	response.CommonBase

	// 配置参数组id
	GroupId int
}

// NewUploadUDBParamGroupRequest will create request of UploadUDBParamGroup action.
func (c *UDBClient) NewUploadUDBParamGroupRequest() *UploadUDBParamGroupRequest {
	req := &UploadUDBParamGroupRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

// UploadUDBParamGroup - 导入UDB配置
func (c *UDBClient) UploadUDBParamGroup(req *UploadUDBParamGroupRequest) (*UploadUDBParamGroupResponse, error) {
	var err error
	var res UploadUDBParamGroupResponse

	err = c.client.InvokeAction("UploadUDBParamGroup", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
