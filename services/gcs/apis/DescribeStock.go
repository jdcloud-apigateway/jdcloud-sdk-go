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
    gcs "github.com/jdcloud-api/jdcloud-sdk-go/services/gcs/models"
)

type DescribeStockRequest struct {

    core.JDCloudRequest

    /* 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域  */
    RegionId string `json:"regionId"`

    /* sku id (Optional) */
    SkuId *string `json:"skuId"`
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeStockRequest(
    regionId string,
) *DescribeStockRequest {

	return &DescribeStockRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/describeStock",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域 (Required)
 * param skuId: sku id (Optional)
 */
func NewDescribeStockRequestWithAllParams(
    regionId string,
    skuId *string,
) *DescribeStockRequest {

    return &DescribeStockRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeStock",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        SkuId: skuId,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeStockRequestWithoutParam() *DescribeStockRequest {

    return &DescribeStockRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/describeStock",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID，可调用接口（describeRegiones）获取云物理服务器支持的地域(Required) */
func (r *DescribeStockRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param skuId: sku id(Optional) */
func (r *DescribeStockRequest) SetSkuId(skuId string) {
    r.SkuId = &skuId
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeStockRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeStockResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeStockResult `json:"result"`
}

type DescribeStockResult struct {
    Stock gcs.SkuStock `json:"stock"`
}