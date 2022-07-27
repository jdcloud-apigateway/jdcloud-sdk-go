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

type CreateAccountRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 账号名称  */
    AccountName string `json:"accountName"`

    /* 账号密码  */
    AccountPassword string `json:"accountPassword"`

    /* 账号权限，默认为读写权限。支持RoleReadOnly（只读权限）、RoleReadWrite（读写权限） (Optional) */
    AccountPrivilege *string `json:"accountPrivilege"`

    /* 账号描述 (Optional) */
    AccountDescription *string `json:"accountDescription"`

}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param accountName: 账号名称 (Required)
 * param accountPassword: 账号密码 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateAccountRequest(
    regionId string,
    cacheInstanceId string,
    accountName string,
    accountPassword string,
) *CreateAccountRequest {

	return &CreateAccountRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/account",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        AccountName: accountName,
        AccountPassword: accountPassword,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param accountName: 账号名称 (Required)
 * param accountPassword: 账号密码 (Required)
 * param accountPrivilege: 账号权限，默认为读写权限。支持RoleReadOnly（只读权限）、RoleReadWrite（读写权限） (Optional)
 * param accountDescription: 账号描述 (Optional)
 */
func NewCreateAccountRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    accountName string,
    accountPassword string,
    accountPrivilege *string,
    accountDescription *string,
) *CreateAccountRequest {

    return &CreateAccountRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/account",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        AccountName: accountName,
        AccountPassword: accountPassword,
        AccountPrivilege: accountPrivilege,
        AccountDescription: accountDescription,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateAccountRequestWithoutParam() *CreateAccountRequest {

    return &CreateAccountRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/account",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *CreateAccountRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *CreateAccountRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param accountName: 账号名称(Required) */
func (r *CreateAccountRequest) SetAccountName(accountName string) {
    r.AccountName = accountName
}

/* param accountPassword: 账号密码(Required) */
func (r *CreateAccountRequest) SetAccountPassword(accountPassword string) {
    r.AccountPassword = accountPassword
}

/* param accountPrivilege: 账号权限，默认为读写权限。支持RoleReadOnly（只读权限）、RoleReadWrite（读写权限）(Optional) */
func (r *CreateAccountRequest) SetAccountPrivilege(accountPrivilege string) {
    r.AccountPrivilege = &accountPrivilege
}

/* param accountDescription: 账号描述(Optional) */
func (r *CreateAccountRequest) SetAccountDescription(accountDescription string) {
    r.AccountDescription = &accountDescription
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateAccountRequest) GetRegionId() string {
    return r.RegionId
}

type CreateAccountResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateAccountResult `json:"result"`
}

type CreateAccountResult struct {
}