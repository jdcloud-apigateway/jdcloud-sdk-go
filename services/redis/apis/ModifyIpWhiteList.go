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

type ModifyIpWhiteListRequest struct {

    core.JDCloudRequest

    /* 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2  */
    RegionId string `json:"regionId"`

    /* 缓存Redis实例ID，是访问实例的唯一标识  */
    CacheInstanceId string `json:"cacheInstanceId"`

    /* 修改后的IP白名单列表，IP格式为CIDR表示法（x.x.x.x/x），0.0.0.0/0表示任何IP、网络都可以访问  */
    IpWhiteList []string `json:"ipWhiteList"`

}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param ipWhiteList: 修改后的IP白名单列表，IP格式为CIDR表示法（x.x.x.x/x），0.0.0.0/0表示任何IP、网络都可以访问 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewModifyIpWhiteListRequest(
    regionId string,
    cacheInstanceId string,
    ipWhiteList []string,
) *ModifyIpWhiteListRequest {

	return &ModifyIpWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/ipWhiteList",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        IpWhiteList: ipWhiteList,
	}
}

/*
 * param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2 (Required)
 * param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识 (Required)
 * param ipWhiteList: 修改后的IP白名单列表，IP格式为CIDR表示法（x.x.x.x/x），0.0.0.0/0表示任何IP、网络都可以访问 (Required)
 */
func NewModifyIpWhiteListRequestWithAllParams(
    regionId string,
    cacheInstanceId string,
    ipWhiteList []string,
) *ModifyIpWhiteListRequest {

    return &ModifyIpWhiteListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/ipWhiteList",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        CacheInstanceId: cacheInstanceId,
        IpWhiteList: ipWhiteList,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewModifyIpWhiteListRequestWithoutParam() *ModifyIpWhiteListRequest {

    return &ModifyIpWhiteListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/cacheInstance/{cacheInstanceId}/ipWhiteList",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 缓存Redis实例所在区域的Region ID。目前有华北-北京、华南-广州、华东-上海三个区域，Region ID分别为cn-north-1、cn-south-1、cn-east-2(Required) */
func (r *ModifyIpWhiteListRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param cacheInstanceId: 缓存Redis实例ID，是访问实例的唯一标识(Required) */
func (r *ModifyIpWhiteListRequest) SetCacheInstanceId(cacheInstanceId string) {
    r.CacheInstanceId = cacheInstanceId
}

/* param ipWhiteList: 修改后的IP白名单列表，IP格式为CIDR表示法（x.x.x.x/x），0.0.0.0/0表示任何IP、网络都可以访问(Required) */
func (r *ModifyIpWhiteListRequest) SetIpWhiteList(ipWhiteList []string) {
    r.IpWhiteList = ipWhiteList
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ModifyIpWhiteListRequest) GetRegionId() string {
    return r.RegionId
}

type ModifyIpWhiteListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ModifyIpWhiteListResult `json:"result"`
}

type ModifyIpWhiteListResult struct {
}