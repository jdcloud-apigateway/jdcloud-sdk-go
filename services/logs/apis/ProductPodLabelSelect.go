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
)

type ProductPodLabelSelectRequest struct {

    core.JDCloudRequest

    /* 操作类型 (Optional) */
    Operate *string `json:"operate"`

    /* pod详情 (Optional) */
    PodInfo *string `json:"podInfo"`

    /* 租户 (Optional) */
    Tenant *string `json:"tenant"`

    /* pod详情数组 (Optional) */
    PodListStr *string `json:"podListStr"`

    /* 产品名 (Optional) */
    ProductName *string `json:"productName"`

    /* 是否所有租户可见 (Optional) */
    AllTenantDistributability *bool `json:"allTenantDistributability"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewProductPodLabelSelectRequest(
) *ProductPodLabelSelectRequest {

	return &ProductPodLabelSelectRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/productPodLabelSelect",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param operate: 操作类型 (Optional)
 * param podInfo: pod详情 (Optional)
 * param tenant: 租户 (Optional)
 * param podListStr: pod详情数组 (Optional)
 * param productName: 产品名 (Optional)
 * param allTenantDistributability: 是否所有租户可见 (Optional)
 */
func NewProductPodLabelSelectRequestWithAllParams(
    operate *string,
    podInfo *string,
    tenant *string,
    podListStr *string,
    productName *string,
    allTenantDistributability *bool,
) *ProductPodLabelSelectRequest {

    return &ProductPodLabelSelectRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/productPodLabelSelect",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Operate: operate,
        PodInfo: podInfo,
        Tenant: tenant,
        PodListStr: podListStr,
        ProductName: productName,
        AllTenantDistributability: allTenantDistributability,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewProductPodLabelSelectRequestWithoutParam() *ProductPodLabelSelectRequest {

    return &ProductPodLabelSelectRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/productPodLabelSelect",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param operate: 操作类型(Optional) */
func (r *ProductPodLabelSelectRequest) SetOperate(operate string) {
    r.Operate = &operate
}
/* param podInfo: pod详情(Optional) */
func (r *ProductPodLabelSelectRequest) SetPodInfo(podInfo string) {
    r.PodInfo = &podInfo
}
/* param tenant: 租户(Optional) */
func (r *ProductPodLabelSelectRequest) SetTenant(tenant string) {
    r.Tenant = &tenant
}
/* param podListStr: pod详情数组(Optional) */
func (r *ProductPodLabelSelectRequest) SetPodListStr(podListStr string) {
    r.PodListStr = &podListStr
}
/* param productName: 产品名(Optional) */
func (r *ProductPodLabelSelectRequest) SetProductName(productName string) {
    r.ProductName = &productName
}
/* param allTenantDistributability: 是否所有租户可见(Optional) */
func (r *ProductPodLabelSelectRequest) SetAllTenantDistributability(allTenantDistributability bool) {
    r.AllTenantDistributability = &allTenantDistributability
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ProductPodLabelSelectRequest) GetRegionId() string {
    return ""
}

type ProductPodLabelSelectResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ProductPodLabelSelectResult `json:"result"`
}

type ProductPodLabelSelectResult struct {
    Result bool `json:"result"`
}