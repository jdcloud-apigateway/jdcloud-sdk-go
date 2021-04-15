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
    disk "github.com/jdcloud-api/jdcloud-sdk-go/services/disk/models"
)

type ApplySnapshotPoliciesRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 绑定/解绑操作 (Optional) */
    Relations []disk.PolicyDiskRelationOp `json:"relations"`
}

/*
 * param regionId: 地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewApplySnapshotPoliciesRequest(
    regionId string,
) *ApplySnapshotPoliciesRequest {

	return &ApplySnapshotPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/snapshotPolicies:apply",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param relations: 绑定/解绑操作 (Optional)
 */
func NewApplySnapshotPoliciesRequestWithAllParams(
    regionId string,
    relations []disk.PolicyDiskRelationOp,
) *ApplySnapshotPoliciesRequest {

    return &ApplySnapshotPoliciesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshotPolicies:apply",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        Relations: relations,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewApplySnapshotPoliciesRequestWithoutParam() *ApplySnapshotPoliciesRequest {

    return &ApplySnapshotPoliciesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/snapshotPolicies:apply",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *ApplySnapshotPoliciesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param relations: 绑定/解绑操作(Optional) */
func (r *ApplySnapshotPoliciesRequest) SetRelations(relations []disk.PolicyDiskRelationOp) {
    r.Relations = relations
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r ApplySnapshotPoliciesRequest) GetRegionId() string {
    return r.RegionId
}

type ApplySnapshotPoliciesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result ApplySnapshotPoliciesResult `json:"result"`
}

type ApplySnapshotPoliciesResult struct {
    OpResults []disk.PolicyDiskRelationOpResult `json:"opResults"`
}