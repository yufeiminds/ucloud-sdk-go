// Code is generated by ucloud-model, DO NOT EDIT IT.

package cube

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// Cube API Schema

// CreateCubePodRequest is request schema for CreateCubePod action
type CreateCubePodRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"true"`

	// 计费模式。枚举值为： \\ > Year，按年付费； \\ > Month，按月付费；\\ > Postpay， \\ 后付费；默认为后付费
	ChargeType *string `required:"false"`

	// Cpu平台（V6、A2），默认V6
	CpuPlatform *string `required:"false"`

	// pod所在组
	Group *string `required:"false"`

	// pod的名字
	Name *string `required:"false"`

	// base64编码的Pod的yaml
	Pod *string `required:"true"`

	// 购买时长。默认:值 1。 月付时，此参数传0，代表购买至月末。
	Quantity *int `required:"false"`

	// 子网Id
	SubnetId *string `required:"true"`

	// 业务组。默认：Default（Default即为未分组）
	Tag *string `required:"false"`

	// VPCId
	VPCId *string `required:"true"`
}

// CreateCubePodResponse is response schema for CreateCubePod action
type CreateCubePodResponse struct {
	response.CommonBase

	// 操作名称
	Action string

	// cube的资源Id
	CubeId string

	// base64编码的yaml
	Pod string

	// 返回码
	RetCode int
}

// NewCreateCubePodRequest will create request of CreateCubePod action.
func (c *CubeClient) NewCreateCubePodRequest() *CreateCubePodRequest {
	req := &CreateCubePodRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateCubePod

创建Pod
*/
func (c *CubeClient) CreateCubePod(req *CreateCubePodRequest) (*CreateCubePodResponse, error) {
	var err error
	var res CreateCubePodResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateCubePod", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeleteCubePodRequest is request schema for DeleteCubePod action
type DeleteCubePodRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// cubeid和uid任意一个（必须）
	CubeId *string `required:"false"`

	// 删除cube时是否释放绑定的EIP。默认为false。
	ReleaseEIP *bool `required:"false"`

	// cubeid和uid任意一个（必须）
	Uid *string `required:"false"`
}

// DeleteCubePodResponse is response schema for DeleteCubePod action
type DeleteCubePodResponse struct {
	response.CommonBase
}

// NewDeleteCubePodRequest will create request of DeleteCubePod action.
func (c *CubeClient) NewDeleteCubePodRequest() *DeleteCubePodRequest {
	req := &DeleteCubePodRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DeleteCubePod

删除Pod
*/
func (c *CubeClient) DeleteCubePod(req *DeleteCubePodRequest) (*DeleteCubePodResponse, error) {
	var err error
	var res DeleteCubePodResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DeleteCubePod", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetCubeExtendInfoRequest is request schema for GetCubeExtendInfo action
type GetCubeExtendInfoRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// id列表以逗号(,)分割
	CubeIds *string `required:"true"`
}

// GetCubeExtendInfoResponse is response schema for GetCubeExtendInfo action
type GetCubeExtendInfoResponse struct {
	response.CommonBase

	// CubeExtendInfo
	ExtendInfo []CubeExtendInfo
}

// NewGetCubeExtendInfoRequest will create request of GetCubeExtendInfo action.
func (c *CubeClient) NewGetCubeExtendInfoRequest() *GetCubeExtendInfoRequest {
	req := &GetCubeExtendInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetCubeExtendInfo

获取Cube的额外信息
*/
func (c *CubeClient) GetCubeExtendInfo(req *GetCubeExtendInfoRequest) (*GetCubeExtendInfoResponse, error) {
	var err error
	var res GetCubeExtendInfoResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetCubeExtendInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetCubePodRequest is request schema for GetCubePod action
type GetCubePodRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// CubeId和Uid任意一个
	CubeId *string `required:"false"`

	// CubeId和Uid任意一个
	Uid *string `required:"false"`
}

// GetCubePodResponse is response schema for GetCubePod action
type GetCubePodResponse struct {
	response.CommonBase

	// base64编码的pod的yaml
	Pod string
}

// NewGetCubePodRequest will create request of GetCubePod action.
func (c *CubeClient) NewGetCubePodRequest() *GetCubePodRequest {
	req := &GetCubePodRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetCubePod

获取Pod的详细信息
*/
func (c *CubeClient) GetCubePod(req *GetCubePodRequest) (*GetCubePodResponse, error) {
	var err error
	var res GetCubePodResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetCubePod", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ListCubePodRequest is request schema for ListCubePod action
type ListCubePodRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 组名称
	Group *string `required:"false"`

	// 默认20
	Limit *int `required:"false"`

	// 默认0
	Offset *int `required:"false"`

	// 子网Id
	SubnetId *string `required:"false"`

	// VPC的Id
	VPCId *string `required:"false"`
}

// ListCubePodResponse is response schema for ListCubePod action
type ListCubePodResponse struct {
	response.CommonBase

	// Pod列表，每条数据都做了base64编码
	Pods []string

	// Cube的总数
	TotalCount int
}

// NewListCubePodRequest will create request of ListCubePod action.
func (c *CubeClient) NewListCubePodRequest() *ListCubePodRequest {
	req := &ListCubePodRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ListCubePod

获取Pods列表
*/
func (c *CubeClient) ListCubePod(req *ListCubePodRequest) (*ListCubePodResponse, error) {
	var err error
	var res ListCubePodResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ListCubePod", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyCubeExtendInfoRequest is request schema for ModifyCubeExtendInfo action
type ModifyCubeExtendInfoRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// cube的id
	CubeId *string `required:"true"`

	// 修改的名字，规则（^[a-zA-Z0-9-_.\u4e00-\u9fa5]{1,32}）
	Name *string `required:"false"`
}

// ModifyCubeExtendInfoResponse is response schema for ModifyCubeExtendInfo action
type ModifyCubeExtendInfoResponse struct {
	response.CommonBase
}

// NewModifyCubeExtendInfoRequest will create request of ModifyCubeExtendInfo action.
func (c *CubeClient) NewModifyCubeExtendInfoRequest() *ModifyCubeExtendInfoRequest {
	req := &ModifyCubeExtendInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ModifyCubeExtendInfo

修改Cube额外信息
*/
func (c *CubeClient) ModifyCubeExtendInfo(req *ModifyCubeExtendInfoRequest) (*ModifyCubeExtendInfoResponse, error) {
	var err error
	var res ModifyCubeExtendInfoResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyCubeExtendInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// RenewCubePodRequest is request schema for RenewCubePod action
type RenewCubePodRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"true"`

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// 容器Id
	CubeId *string `required:"true"`

	// base64编码的Pod的yaml
	Pod *string `required:"true"`
}

// RenewCubePodResponse is response schema for RenewCubePod action
type RenewCubePodResponse struct {
	response.CommonBase

	// base64编码过的yaml，需要解码获取信息
	Pod string
}

// NewRenewCubePodRequest will create request of RenewCubePod action.
func (c *CubeClient) NewRenewCubePodRequest() *RenewCubePodRequest {
	req := &RenewCubePodRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: RenewCubePod

更新Pod
*/
func (c *CubeClient) RenewCubePod(req *RenewCubePodRequest) (*RenewCubePodResponse, error) {
	var err error
	var res RenewCubePodResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RenewCubePod", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
