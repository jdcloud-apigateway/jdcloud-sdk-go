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
    waf "github.com/jdcloud-api/jdcloud-sdk-go/services/waf/models"
)

type EnableCname2RsExternalRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 实例Id  */
    WafInstanceId string `json:"wafInstanceId"`

    /* 请求  */
    Req *waf.ActiveFuncReq `json:"req"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例Id (Required)
 * param req: 请求 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEnableCname2RsExternalRequest(
    regionId string,
    wafInstanceId string,
    req *waf.ActiveFuncReq,
) *EnableCname2RsExternalRequest {

	return &EnableCname2RsExternalRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/domain:enableCname2RsExternal",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        WafInstanceId: wafInstanceId,
        Req: req,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param wafInstanceId: 实例Id (Required)
 * param req: 请求 (Required)
 */
func NewEnableCname2RsExternalRequestWithAllParams(
    regionId string,
    wafInstanceId string,
    req *waf.ActiveFuncReq,
) *EnableCname2RsExternalRequest {

    return &EnableCname2RsExternalRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/domain:enableCname2RsExternal",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        WafInstanceId: wafInstanceId,
        Req: req,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEnableCname2RsExternalRequestWithoutParam() *EnableCname2RsExternalRequest {

    return &EnableCname2RsExternalRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/wafInstanceIds/{wafInstanceId}/domain:enableCname2RsExternal",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *EnableCname2RsExternalRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param wafInstanceId: 实例Id(Required) */
func (r *EnableCname2RsExternalRequest) SetWafInstanceId(wafInstanceId string) {
    r.WafInstanceId = wafInstanceId
}

/* param req: 请求(Required) */
func (r *EnableCname2RsExternalRequest) SetReq(req *waf.ActiveFuncReq) {
    r.Req = req
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EnableCname2RsExternalRequest) GetRegionId() string {
    return r.RegionId
}

type EnableCname2RsExternalResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EnableCname2RsExternalResult `json:"result"`
}

type EnableCname2RsExternalResult struct {
}