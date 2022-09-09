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

type UpdateBastionDescRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 堡垒机id  */
    Bid string `json:"bid"`

    /* 描述信息  */
    Desc string `json:"desc"`
}

/*
 * param regionId: 地域ID (Required)
 * param bid: 堡垒机id (Required)
 * param desc: 描述信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateBastionDescRequest(
    regionId string,
    bid string,
    desc string,
) *UpdateBastionDescRequest {

	return &UpdateBastionDescRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/bastion/{bid}",
			Method:  "PATCH",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        Bid: bid,
        Desc: desc,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param bid: 堡垒机id (Required)
 * param desc: 描述信息 (Required)
 */
func NewUpdateBastionDescRequestWithAllParams(
    regionId string,
    bid string,
    desc string,
) *UpdateBastionDescRequest {

    return &UpdateBastionDescRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bastion/{bid}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Bid: bid,
        Desc: desc,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateBastionDescRequestWithoutParam() *UpdateBastionDescRequest {

    return &UpdateBastionDescRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/bastion/{bid}",
            Method:  "PATCH",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UpdateBastionDescRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param bid: 堡垒机id(Required) */
func (r *UpdateBastionDescRequest) SetBid(bid string) {
    r.Bid = bid
}
/* param desc: 描述信息(Required) */
func (r *UpdateBastionDescRequest) SetDesc(desc string) {
    r.Desc = desc
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateBastionDescRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateBastionDescResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateBastionDescResult `json:"result"`
}

type UpdateBastionDescResult struct {
    Result bool `json:"result"`
}