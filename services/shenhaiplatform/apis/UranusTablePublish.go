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
    shenhaiplatform "github.com/jdcloud-api/jdcloud-sdk-go/services/shenhaiplatform/models"
)

type UranusTablePublishRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 应用名称  */
    AppName string `json:"appName"`

    /* 表英文名称  */
    TableEnName string `json:"tableEnName"`

    /* 表中文名 (Optional) */
    TableCnName *string `json:"tableCnName"`

    /* 负责人 (Optional) */
    Owner *string `json:"owner"`

    /* 协助人 (Optional) */
    Manager *string `json:"manager"`

    /* 表的普通字段信息  */
    Columns []shenhaiplatform.UranusColumnSaveOrUpdate `json:"columns"`

    /* 表的分区字段信息 (Optional) */
    Partitions []shenhaiplatform.UranusColumnSaveOrUpdate `json:"partitions"`

    /* 0 非分区表 1 分区表 (Optional) */
    IsPartition *int `json:"isPartition"`

    /* 0 新建表 1 修改表 (Optional) */
    UpdateTable *int `json:"updateTable"`

    /* 表存储格式 (Optional) */
    StorageType *string `json:"storageType"`

    /* 表字段分隔符（只有TEXTFILE类型表需要该字段） (Optional) */
    FieldDelim *string `json:"fieldDelim"`

    /* 表行分隔符（只有TEXTFILE类型表需要该字段） (Optional) */
    LineDelim *string `json:"lineDelim"`

    /* 分桶字段 (Optional) */
    BucketCols []string `json:"bucketCols"`

    /* 分桶字段 (Optional) */
    BucketNum *int `json:"bucketNum"`
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param tableEnName: 表英文名称 (Required)
 * param columns: 表的普通字段信息 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUranusTablePublishRequest(
    regionId string,
    appName string,
    tableEnName string,
    columns []shenhaiplatform.UranusColumnSaveOrUpdate,
) *UranusTablePublishRequest {

	return &UranusTablePublishRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/apps/{appName}/uranusTablePublish",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
        AppName: appName,
        TableEnName: tableEnName,
        Columns: columns,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param appName: 应用名称 (Required)
 * param tableEnName: 表英文名称 (Required)
 * param tableCnName: 表中文名 (Optional)
 * param owner: 负责人 (Optional)
 * param manager: 协助人 (Optional)
 * param columns: 表的普通字段信息 (Required)
 * param partitions: 表的分区字段信息 (Optional)
 * param isPartition: 0 非分区表 1 分区表 (Optional)
 * param updateTable: 0 新建表 1 修改表 (Optional)
 * param storageType: 表存储格式 (Optional)
 * param fieldDelim: 表字段分隔符（只有TEXTFILE类型表需要该字段） (Optional)
 * param lineDelim: 表行分隔符（只有TEXTFILE类型表需要该字段） (Optional)
 * param bucketCols: 分桶字段 (Optional)
 * param bucketNum: 分桶字段 (Optional)
 */
func NewUranusTablePublishRequestWithAllParams(
    regionId string,
    appName string,
    tableEnName string,
    tableCnName *string,
    owner *string,
    manager *string,
    columns []shenhaiplatform.UranusColumnSaveOrUpdate,
    partitions []shenhaiplatform.UranusColumnSaveOrUpdate,
    isPartition *int,
    updateTable *int,
    storageType *string,
    fieldDelim *string,
    lineDelim *string,
    bucketCols []string,
    bucketNum *int,
) *UranusTablePublishRequest {

    return &UranusTablePublishRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTablePublish",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        AppName: appName,
        TableEnName: tableEnName,
        TableCnName: tableCnName,
        Owner: owner,
        Manager: manager,
        Columns: columns,
        Partitions: partitions,
        IsPartition: isPartition,
        UpdateTable: updateTable,
        StorageType: storageType,
        FieldDelim: fieldDelim,
        LineDelim: lineDelim,
        BucketCols: bucketCols,
        BucketNum: bucketNum,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUranusTablePublishRequestWithoutParam() *UranusTablePublishRequest {

    return &UranusTablePublishRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/apps/{appName}/uranusTablePublish",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UranusTablePublishRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param appName: 应用名称(Required) */
func (r *UranusTablePublishRequest) SetAppName(appName string) {
    r.AppName = appName
}
/* param tableEnName: 表英文名称(Required) */
func (r *UranusTablePublishRequest) SetTableEnName(tableEnName string) {
    r.TableEnName = tableEnName
}
/* param tableCnName: 表中文名(Optional) */
func (r *UranusTablePublishRequest) SetTableCnName(tableCnName string) {
    r.TableCnName = &tableCnName
}
/* param owner: 负责人(Optional) */
func (r *UranusTablePublishRequest) SetOwner(owner string) {
    r.Owner = &owner
}
/* param manager: 协助人(Optional) */
func (r *UranusTablePublishRequest) SetManager(manager string) {
    r.Manager = &manager
}
/* param columns: 表的普通字段信息(Required) */
func (r *UranusTablePublishRequest) SetColumns(columns []shenhaiplatform.UranusColumnSaveOrUpdate) {
    r.Columns = columns
}
/* param partitions: 表的分区字段信息(Optional) */
func (r *UranusTablePublishRequest) SetPartitions(partitions []shenhaiplatform.UranusColumnSaveOrUpdate) {
    r.Partitions = partitions
}
/* param isPartition: 0 非分区表 1 分区表(Optional) */
func (r *UranusTablePublishRequest) SetIsPartition(isPartition int) {
    r.IsPartition = &isPartition
}
/* param updateTable: 0 新建表 1 修改表(Optional) */
func (r *UranusTablePublishRequest) SetUpdateTable(updateTable int) {
    r.UpdateTable = &updateTable
}
/* param storageType: 表存储格式(Optional) */
func (r *UranusTablePublishRequest) SetStorageType(storageType string) {
    r.StorageType = &storageType
}
/* param fieldDelim: 表字段分隔符（只有TEXTFILE类型表需要该字段）(Optional) */
func (r *UranusTablePublishRequest) SetFieldDelim(fieldDelim string) {
    r.FieldDelim = &fieldDelim
}
/* param lineDelim: 表行分隔符（只有TEXTFILE类型表需要该字段）(Optional) */
func (r *UranusTablePublishRequest) SetLineDelim(lineDelim string) {
    r.LineDelim = &lineDelim
}
/* param bucketCols: 分桶字段(Optional) */
func (r *UranusTablePublishRequest) SetBucketCols(bucketCols []string) {
    r.BucketCols = bucketCols
}
/* param bucketNum: 分桶字段(Optional) */
func (r *UranusTablePublishRequest) SetBucketNum(bucketNum int) {
    r.BucketNum = &bucketNum
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UranusTablePublishRequest) GetRegionId() string {
    return r.RegionId
}

type UranusTablePublishResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UranusTablePublishResult `json:"result"`
}

type UranusTablePublishResult struct {
    Code string `json:"code"`
    ErrorTitle string `json:"errorTitle"`
    ErrorMsg string `json:"errorMsg"`
    Result interface{} `json:"result"`
    SubCode string `json:"subCode"`
    Successed bool `json:"successed"`
}