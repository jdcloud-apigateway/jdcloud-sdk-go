// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    billing "github.com/jdcloud-api/jdcloud-sdk-go/services/billing/models"
)

type GetExpiringOrdersRequest struct {

    core.JDCloudRequest

    /*   */
    RegionId string `json:"regionId"`

    /* 主键id (Optional) */
    Id *int `json:"id"`

    /* 用户pin (Optional) */
    Pin *string `json:"pin"`

    /* 资源id (Optional) */
    ResourceId *string `json:"resourceId"`

    /* appCode (Optional) */
    AppCode *string `json:"appCode"`

    /* 产品码列表 (Optional) */
    ServiceCodeList []string `json:"serviceCodeList"`

    /* 地域 (Optional) */
    Region *string `json:"region"`

    /* 网络类型 0: non-BGP, 1: BGP (Optional) */
    NetworkOperator *int `json:"networkOperator"`

    /* 计费类型 1:按配置 2:按用量 3:包年包月 (Optional) */
    BillingType *int `json:"billingType"`

    /* resourceIdList (Optional) */
    ResourceIdList []string `json:"resourceIdList"`

    /* >0: 订单还有几天到期; ==0: 订单已经到期; <0: 不管是否到期 (Optional) */
    ExpireInDays *int `json:"expireInDays"`

    /* isOnTrial (Optional) */
    IsOnTrial *int `json:"isOnTrial"`

    /* 站点信息 0:中国 1:国际 (Optional) */
    Site *int `json:"site"`

    /* 资源状态 1:正常 2:停服 3:删除 (Optional) */
    Status *int `json:"status"`

    /* 计费状态 0:停止计费 1:计费中 (Optional) */
    BillingStatus *int `json:"billingStatus"`

    /* 1、内部计算，使用getExpiringOrderCount时使用，不用传值 默认为-1mybatis筛选时不会关注 2、selectResourceOrderForTask定时任务查询列表时使用，停服天数，必须大于0 (Optional) */
    ExpiringInDays *int `json:"expiringInDays"`

    /* billingTypeList (Optional) */
    BillingTypeList []int `json:"billingTypeList"`

    /* 交易单号列表 (Optional) */
    TransactionNos []string `json:"transactionNos"`

    /*  (Optional) */
    OpTypes []int `json:"opTypes"`

    /* 开始时间 (Optional) */
    StartTime *string `json:"startTime"`

    /* 结束时间 (Optional) */
    EndTime *string `json:"endTime"`

    /* 服务编码 (Optional) */
    ServiceCode *string `json:"serviceCode"`

    /* statusList (Optional) */
    StatusList []int `json:"statusList"`

    /* excludeResources (Optional) */
    ExcludeResources []string `json:"excludeResources"`

    /* orderByClaus (Optional) */
    OrderByClaus *string `json:"orderByClaus"`

    /* 专有云节点的code（节点云查询使用） (Optional) */
    Node *string `json:"node"`

    /* 部门（节点云查询使用） (Optional) */
    DepartmentId *int `json:"departmentId"`

    /*  (Optional) */
    PinList []string `json:"pinList"`

    /* 是否是专有云 1:是  其他不是 (Optional) */
    IsSpecial *int `json:"isSpecial"`

    /* 节点信息 (Optional) */
    NodeCode *string `json:"nodeCode"`

    /* 超时时间，暂时不用 (Optional) */
    Timeout *int `json:"timeout"`

    /* 当前页序号 (Optional) */
    PageIndex *int `json:"pageIndex"`

    /* 每页结果数量 (Optional) */
    PageSize *int `json:"pageSize"`

    /*  (Optional) */
    Offset *int `json:"offset"`
}

/*
 * param regionId:  (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewGetExpiringOrdersRequest(
    regionId string,
) *GetExpiringOrdersRequest {

	return &GetExpiringOrdersRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/getExpiringOrders",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId:  (Required)
 * param id: 主键id (Optional)
 * param pin: 用户pin (Optional)
 * param resourceId: 资源id (Optional)
 * param appCode: appCode (Optional)
 * param serviceCodeList: 产品码列表 (Optional)
 * param region: 地域 (Optional)
 * param networkOperator: 网络类型 0: non-BGP, 1: BGP (Optional)
 * param billingType: 计费类型 1:按配置 2:按用量 3:包年包月 (Optional)
 * param resourceIdList: resourceIdList (Optional)
 * param expireInDays: >0: 订单还有几天到期; ==0: 订单已经到期; <0: 不管是否到期 (Optional)
 * param isOnTrial: isOnTrial (Optional)
 * param site: 站点信息 0:中国 1:国际 (Optional)
 * param status: 资源状态 1:正常 2:停服 3:删除 (Optional)
 * param billingStatus: 计费状态 0:停止计费 1:计费中 (Optional)
 * param expiringInDays: 1、内部计算，使用getExpiringOrderCount时使用，不用传值 默认为-1mybatis筛选时不会关注 2、selectResourceOrderForTask定时任务查询列表时使用，停服天数，必须大于0 (Optional)
 * param billingTypeList: billingTypeList (Optional)
 * param transactionNos: 交易单号列表 (Optional)
 * param opTypes:  (Optional)
 * param startTime: 开始时间 (Optional)
 * param endTime: 结束时间 (Optional)
 * param serviceCode: 服务编码 (Optional)
 * param statusList: statusList (Optional)
 * param excludeResources: excludeResources (Optional)
 * param orderByClaus: orderByClaus (Optional)
 * param node: 专有云节点的code（节点云查询使用） (Optional)
 * param departmentId: 部门（节点云查询使用） (Optional)
 * param pinList:  (Optional)
 * param isSpecial: 是否是专有云 1:是  其他不是 (Optional)
 * param nodeCode: 节点信息 (Optional)
 * param timeout: 超时时间，暂时不用 (Optional)
 * param pageIndex: 当前页序号 (Optional)
 * param pageSize: 每页结果数量 (Optional)
 * param offset:  (Optional)
 */
func NewGetExpiringOrdersRequestWithAllParams(
    regionId string,
    id *int,
    pin *string,
    resourceId *string,
    appCode *string,
    serviceCodeList []string,
    region *string,
    networkOperator *int,
    billingType *int,
    resourceIdList []string,
    expireInDays *int,
    isOnTrial *int,
    site *int,
    status *int,
    billingStatus *int,
    expiringInDays *int,
    billingTypeList []int,
    transactionNos []string,
    opTypes []int,
    startTime *string,
    endTime *string,
    serviceCode *string,
    statusList []int,
    excludeResources []string,
    orderByClaus *string,
    node *string,
    departmentId *int,
    pinList []string,
    isSpecial *int,
    nodeCode *string,
    timeout *int,
    pageIndex *int,
    pageSize *int,
    offset *int,
) *GetExpiringOrdersRequest {

    return &GetExpiringOrdersRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/getExpiringOrders",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Id: id,
        Pin: pin,
        ResourceId: resourceId,
        AppCode: appCode,
        ServiceCodeList: serviceCodeList,
        Region: region,
        NetworkOperator: networkOperator,
        BillingType: billingType,
        ResourceIdList: resourceIdList,
        ExpireInDays: expireInDays,
        IsOnTrial: isOnTrial,
        Site: site,
        Status: status,
        BillingStatus: billingStatus,
        ExpiringInDays: expiringInDays,
        BillingTypeList: billingTypeList,
        TransactionNos: transactionNos,
        OpTypes: opTypes,
        StartTime: startTime,
        EndTime: endTime,
        ServiceCode: serviceCode,
        StatusList: statusList,
        ExcludeResources: excludeResources,
        OrderByClaus: orderByClaus,
        Node: node,
        DepartmentId: departmentId,
        PinList: pinList,
        IsSpecial: isSpecial,
        NodeCode: nodeCode,
        Timeout: timeout,
        PageIndex: pageIndex,
        PageSize: pageSize,
        Offset: offset,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewGetExpiringOrdersRequestWithoutParam() *GetExpiringOrdersRequest {

    return &GetExpiringOrdersRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/getExpiringOrders",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: (Required) */
func (r *GetExpiringOrdersRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param id: 主键id(Optional) */
func (r *GetExpiringOrdersRequest) SetId(id int) {
    r.Id = &id
}

/* param pin: 用户pin(Optional) */
func (r *GetExpiringOrdersRequest) SetPin(pin string) {
    r.Pin = &pin
}

/* param resourceId: 资源id(Optional) */
func (r *GetExpiringOrdersRequest) SetResourceId(resourceId string) {
    r.ResourceId = &resourceId
}

/* param appCode: appCode(Optional) */
func (r *GetExpiringOrdersRequest) SetAppCode(appCode string) {
    r.AppCode = &appCode
}

/* param serviceCodeList: 产品码列表(Optional) */
func (r *GetExpiringOrdersRequest) SetServiceCodeList(serviceCodeList []string) {
    r.ServiceCodeList = serviceCodeList
}

/* param region: 地域(Optional) */
func (r *GetExpiringOrdersRequest) SetRegion(region string) {
    r.Region = &region
}

/* param networkOperator: 网络类型 0: non-BGP, 1: BGP(Optional) */
func (r *GetExpiringOrdersRequest) SetNetworkOperator(networkOperator int) {
    r.NetworkOperator = &networkOperator
}

/* param billingType: 计费类型 1:按配置 2:按用量 3:包年包月(Optional) */
func (r *GetExpiringOrdersRequest) SetBillingType(billingType int) {
    r.BillingType = &billingType
}

/* param resourceIdList: resourceIdList(Optional) */
func (r *GetExpiringOrdersRequest) SetResourceIdList(resourceIdList []string) {
    r.ResourceIdList = resourceIdList
}

/* param expireInDays: >0: 订单还有几天到期; ==0: 订单已经到期; <0: 不管是否到期(Optional) */
func (r *GetExpiringOrdersRequest) SetExpireInDays(expireInDays int) {
    r.ExpireInDays = &expireInDays
}

/* param isOnTrial: isOnTrial(Optional) */
func (r *GetExpiringOrdersRequest) SetIsOnTrial(isOnTrial int) {
    r.IsOnTrial = &isOnTrial
}

/* param site: 站点信息 0:中国 1:国际(Optional) */
func (r *GetExpiringOrdersRequest) SetSite(site int) {
    r.Site = &site
}

/* param status: 资源状态 1:正常 2:停服 3:删除(Optional) */
func (r *GetExpiringOrdersRequest) SetStatus(status int) {
    r.Status = &status
}

/* param billingStatus: 计费状态 0:停止计费 1:计费中(Optional) */
func (r *GetExpiringOrdersRequest) SetBillingStatus(billingStatus int) {
    r.BillingStatus = &billingStatus
}

/* param expiringInDays: 1、内部计算，使用getExpiringOrderCount时使用，不用传值 默认为-1mybatis筛选时不会关注 2、selectResourceOrderForTask定时任务查询列表时使用，停服天数，必须大于0(Optional) */
func (r *GetExpiringOrdersRequest) SetExpiringInDays(expiringInDays int) {
    r.ExpiringInDays = &expiringInDays
}

/* param billingTypeList: billingTypeList(Optional) */
func (r *GetExpiringOrdersRequest) SetBillingTypeList(billingTypeList []int) {
    r.BillingTypeList = billingTypeList
}

/* param transactionNos: 交易单号列表(Optional) */
func (r *GetExpiringOrdersRequest) SetTransactionNos(transactionNos []string) {
    r.TransactionNos = transactionNos
}

/* param opTypes: (Optional) */
func (r *GetExpiringOrdersRequest) SetOpTypes(opTypes []int) {
    r.OpTypes = opTypes
}

/* param startTime: 开始时间(Optional) */
func (r *GetExpiringOrdersRequest) SetStartTime(startTime string) {
    r.StartTime = &startTime
}

/* param endTime: 结束时间(Optional) */
func (r *GetExpiringOrdersRequest) SetEndTime(endTime string) {
    r.EndTime = &endTime
}

/* param serviceCode: 服务编码(Optional) */
func (r *GetExpiringOrdersRequest) SetServiceCode(serviceCode string) {
    r.ServiceCode = &serviceCode
}

/* param statusList: statusList(Optional) */
func (r *GetExpiringOrdersRequest) SetStatusList(statusList []int) {
    r.StatusList = statusList
}

/* param excludeResources: excludeResources(Optional) */
func (r *GetExpiringOrdersRequest) SetExcludeResources(excludeResources []string) {
    r.ExcludeResources = excludeResources
}

/* param orderByClaus: orderByClaus(Optional) */
func (r *GetExpiringOrdersRequest) SetOrderByClaus(orderByClaus string) {
    r.OrderByClaus = &orderByClaus
}

/* param node: 专有云节点的code（节点云查询使用）(Optional) */
func (r *GetExpiringOrdersRequest) SetNode(node string) {
    r.Node = &node
}

/* param departmentId: 部门（节点云查询使用）(Optional) */
func (r *GetExpiringOrdersRequest) SetDepartmentId(departmentId int) {
    r.DepartmentId = &departmentId
}

/* param pinList: (Optional) */
func (r *GetExpiringOrdersRequest) SetPinList(pinList []string) {
    r.PinList = pinList
}

/* param isSpecial: 是否是专有云 1:是  其他不是(Optional) */
func (r *GetExpiringOrdersRequest) SetIsSpecial(isSpecial int) {
    r.IsSpecial = &isSpecial
}

/* param nodeCode: 节点信息(Optional) */
func (r *GetExpiringOrdersRequest) SetNodeCode(nodeCode string) {
    r.NodeCode = &nodeCode
}

/* param timeout: 超时时间，暂时不用(Optional) */
func (r *GetExpiringOrdersRequest) SetTimeout(timeout int) {
    r.Timeout = &timeout
}

/* param pageIndex: 当前页序号(Optional) */
func (r *GetExpiringOrdersRequest) SetPageIndex(pageIndex int) {
    r.PageIndex = &pageIndex
}

/* param pageSize: 每页结果数量(Optional) */
func (r *GetExpiringOrdersRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}

/* param offset: (Optional) */
func (r *GetExpiringOrdersRequest) SetOffset(offset int) {
    r.Offset = &offset
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r GetExpiringOrdersRequest) GetRegionId() string {
    return r.RegionId
}

type GetExpiringOrdersResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result GetExpiringOrdersResult `json:"result"`
}

type GetExpiringOrdersResult struct {
    Pagination billing.Pagination `json:"pagination"`
    Result []billing.ResourceOrderVo `json:"result"`
}