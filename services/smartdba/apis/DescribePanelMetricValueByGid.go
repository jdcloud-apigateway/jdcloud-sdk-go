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
    smartdba "github.com/jdcloud-api/jdcloud-sdk-go/services/smartdba/models"
)

type DescribePanelMetricValueByGidRequest struct {

    core.JDCloudRequest

    /* 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)  */
    RegionId string `json:"regionId"`

    /* 监控大盘id  */
    PanelGid string `json:"panelGid"`

    /* 监控指标，如： tps、qps 等，在supportMetrics接口有返回  */
    Metric string `json:"metric"`

    /* 查询起始时间，格式为：2021-11-11T15:04:05Z  */
    StartTime string `json:"startTime"`

    /* 查询截止时间，格式为：2021-11-11T15:04:05Z  */
    EndTime string `json:"endTime"`

    /* 数据库类型,默认MySQL (Optional) */
    DbType *string `json:"dbType"`
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param panelGid: 监控大盘id (Required)
 * param metric: 监控指标，如： tps、qps 等，在supportMetrics接口有返回 (Required)
 * param startTime: 查询起始时间，格式为：2021-11-11T15:04:05Z (Required)
 * param endTime: 查询截止时间，格式为：2021-11-11T15:04:05Z (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribePanelMetricValueByGidRequest(
    regionId string,
    panelGid string,
    metric string,
    startTime string,
    endTime string,
) *DescribePanelMetricValueByGidRequest {

	return &DescribePanelMetricValueByGidRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/panels/{panelGid}/metric/{metric}",
			Method:  "GET",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        PanelGid: panelGid,
        Metric: metric,
        StartTime: startTime,
        EndTime: endTime,
	}
}

/*
 * param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md) (Required)
 * param panelGid: 监控大盘id (Required)
 * param metric: 监控指标，如： tps、qps 等，在supportMetrics接口有返回 (Required)
 * param startTime: 查询起始时间，格式为：2021-11-11T15:04:05Z (Required)
 * param endTime: 查询截止时间，格式为：2021-11-11T15:04:05Z (Required)
 * param dbType: 数据库类型,默认MySQL (Optional)
 */
func NewDescribePanelMetricValueByGidRequestWithAllParams(
    regionId string,
    panelGid string,
    metric string,
    startTime string,
    endTime string,
    dbType *string,
) *DescribePanelMetricValueByGidRequest {

    return &DescribePanelMetricValueByGidRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/panels/{panelGid}/metric/{metric}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        PanelGid: panelGid,
        Metric: metric,
        StartTime: startTime,
        EndTime: endTime,
        DbType: dbType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribePanelMetricValueByGidRequestWithoutParam() *DescribePanelMetricValueByGidRequest {

    return &DescribePanelMetricValueByGidRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/panels/{panelGid}/metric/{metric}",
            Method:  "GET",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域代码，取值范围参见[《各地域及可用区对照表》](../Enum-Definitions/Regions-AZ.md)(Required) */
func (r *DescribePanelMetricValueByGidRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param panelGid: 监控大盘id(Required) */
func (r *DescribePanelMetricValueByGidRequest) SetPanelGid(panelGid string) {
    r.PanelGid = panelGid
}

/* param metric: 监控指标，如： tps、qps 等，在supportMetrics接口有返回(Required) */
func (r *DescribePanelMetricValueByGidRequest) SetMetric(metric string) {
    r.Metric = metric
}

/* param startTime: 查询起始时间，格式为：2021-11-11T15:04:05Z(Required) */
func (r *DescribePanelMetricValueByGidRequest) SetStartTime(startTime string) {
    r.StartTime = startTime
}

/* param endTime: 查询截止时间，格式为：2021-11-11T15:04:05Z(Required) */
func (r *DescribePanelMetricValueByGidRequest) SetEndTime(endTime string) {
    r.EndTime = endTime
}

/* param dbType: 数据库类型,默认MySQL(Optional) */
func (r *DescribePanelMetricValueByGidRequest) SetDbType(dbType string) {
    r.DbType = &dbType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribePanelMetricValueByGidRequest) GetRegionId() string {
    return r.RegionId
}

type DescribePanelMetricValueByGidResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribePanelMetricValueByGidResult `json:"result"`
}

type DescribePanelMetricValueByGidResult struct {
    MetricDatas []smartdba.MetricData `json:"metricDatas"`
}