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

type ZoneSecurityReportRequest struct {

    core.JDCloudRequest

    /* 域名标识  */
    Zone_identifier string `json:"zone_identifier"`

    /* 域名  */
    ZoneName string `json:"zoneName"`

    /* 开始时间  */
    Since string `json:"since"`

    /* 结束时间  */
    Until string `json:"until"`
}

/*
 * param zone_identifier: 域名标识 (Required)
 * param zoneName: 域名 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewZoneSecurityReportRequest(
    zone_identifier string,
    zoneName string,
    since string,
    until string,
) *ZoneSecurityReportRequest {

	return &ZoneSecurityReportRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/zones/{zone_identifier}/zoneSecurityReport",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        Zone_identifier: zone_identifier,
        ZoneName: zoneName,
        Since: since,
        Until: until,
	}
}

/*
 * param zone_identifier: 域名标识 (Required)
 * param zoneName: 域名 (Required)
 * param since: 开始时间 (Required)
 * param until: 结束时间 (Required)
 */
func NewZoneSecurityReportRequestWithAllParams(
    zone_identifier string,
    zoneName string,
    since string,
    until string,
) *ZoneSecurityReportRequest {

    return &ZoneSecurityReportRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/zoneSecurityReport",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Zone_identifier: zone_identifier,
        ZoneName: zoneName,
        Since: since,
        Until: until,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewZoneSecurityReportRequestWithoutParam() *ZoneSecurityReportRequest {

    return &ZoneSecurityReportRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/zones/{zone_identifier}/zoneSecurityReport",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param zone_identifier: 域名标识(Required) */
func (r *ZoneSecurityReportRequest) SetZone_identifier(zone_identifier string) {
    r.Zone_identifier = zone_identifier
}
/* param zoneName: 域名(Required) */
func (r *ZoneSecurityReportRequest) SetZoneName(zoneName string) {
    r.ZoneName = zoneName
}
/* param since: 开始时间(Required) */
func (r *ZoneSecurityReportRequest) SetSince(since string) {
    r.Since = since
}
/* param until: 结束时间(Required) */
func (r *ZoneSecurityReportRequest) SetUntil(until string) {
    r.Until = until
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ZoneSecurityReportRequest) GetRegionId() string {
    return ""
}

type ZoneSecurityReportResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ZoneSecurityReportResult `json:"result"`
}

type ZoneSecurityReportResult struct {
    ZoneName string `json:"zoneName"`
    QueryDate string `json:"queryDate"`
    RequestSum float64 `json:"requestSum"`
    TrafficSum float64 `json:"trafficSum"`
    CleanTrafficSum float64 `json:"cleanTrafficSum"`
    CachedPercentage string `json:"cachedPercentage"`
    AttackSum float64 `json:"attackSum"`
    AttackPercentage string `json:"attackPercentage"`
    WebAttackDefenseTrends []starshield.WebAttackDefenseTrends `json:"webAttackDefenseTrends"`
    AttackInfo []starshield.AttackInfo `json:"attackInfo"`
}