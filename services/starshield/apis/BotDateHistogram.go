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
    starshield "github.com/jdcloud-api/jdcloud-sdk-go/services/starshield/models"
)

type BotDateHistogramRequest struct {

    core.JDCloudRequest

    /* 域名标识  */
    ZoneId string `json:"zoneId"`

    /* 开始时间  */
    Since string `json:"since"`

    /* 结束时间  */
    Until string `json:"until"`

    /*  (Optional) */
    Filters []starshield.BotFilter `json:"filters"`
}

/*
 * param zoneId: 域名标识 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewBotDateHistogramRequest(
    zoneId string,
    since string,
    until string,
) *BotDateHistogramRequest {

	return &BotDateHistogramRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zoneId}/analyticsBotDateHistogram",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
        ZoneId: zoneId,
        Since: since,
        Until: until,
	}
}

/*
 * param zoneId: 域名标识 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 * param filters:  (Optional)
 */
func NewBotDateHistogramRequestWithAllParams(
    zoneId string,
    since string,
    until string,
    filters []starshield.BotFilter,
) *BotDateHistogramRequest {

    return &BotDateHistogramRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zoneId}/analyticsBotDateHistogram",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
        ZoneId: zoneId,
        Since: since,
        Until: until,
        Filters: filters,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewBotDateHistogramRequestWithoutParam() *BotDateHistogramRequest {

    return &BotDateHistogramRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zoneId}/analyticsBotDateHistogram",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zoneId: 域名标识(Required) */
func (r *BotDateHistogramRequest) SetZoneId(zoneId string) {
    r.ZoneId = zoneId
}
/* param since: 开始时间(Required) */
func (r *BotDateHistogramRequest) SetSince(since string) {
    r.Since = since
}
/* param until: 结束时间(Required) */
func (r *BotDateHistogramRequest) SetUntil(until string) {
    r.Until = until
}
/* param filters: (Optional) */
func (r *BotDateHistogramRequest) SetFilters(filters []starshield.BotFilter) {
    r.Filters = filters
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r BotDateHistogramRequest) GetRegionId() string {
    return ""
}

type BotDateHistogramResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result BotDateHistogramResult `json:"result"`
}

type BotDateHistogramResult struct {
    DataSeries []starshield.RequestBotGroup `json:"dataSeries"`
    TimeSeries []int `json:"timeSeries"`
}