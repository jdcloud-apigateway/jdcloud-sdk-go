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

type K8sconfigNewV2Request struct {

    core.JDCloudRequest

    /* 租户字段 (Optional) */
    TenantName *string `json:"tenantName"`

    /* 集群字段 (Optional) */
    ClusterName *string `json:"clusterName"`

    /* 机房字段 (Optional) */
    ZoneName *string `json:"zoneName"`

    /* node-ip字段 (Optional) */
    NodeIp *string `json:"nodeIp"`

    /* 采集配置字段 (Optional) */
    Conf *string `json:"conf"`

    /* 是否为公有云 (Optional) */
    Cloud *bool `json:"cloud"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewK8sconfigNewV2Request(
) *K8sconfigNewV2Request {

	return &K8sconfigNewV2Request{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/k8sconfigNewV2",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param tenantName: 租户字段 (Optional)
 * param clusterName: 集群字段 (Optional)
 * param zoneName: 机房字段 (Optional)
 * param nodeIp: node-ip字段 (Optional)
 * param conf: 采集配置字段 (Optional)
 * param cloud: 是否为公有云 (Optional)
 */
func NewK8sconfigNewV2RequestWithAllParams(
    tenantName *string,
    clusterName *string,
    zoneName *string,
    nodeIp *string,
    conf *string,
    cloud *bool,
) *K8sconfigNewV2Request {

    return &K8sconfigNewV2Request{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/k8sconfigNewV2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        TenantName: tenantName,
        ClusterName: clusterName,
        ZoneName: zoneName,
        NodeIp: nodeIp,
        Conf: conf,
        Cloud: cloud,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewK8sconfigNewV2RequestWithoutParam() *K8sconfigNewV2Request {

    return &K8sconfigNewV2Request{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/k8sconfigNewV2",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param tenantName: 租户字段(Optional) */
func (r *K8sconfigNewV2Request) SetTenantName(tenantName string) {
    r.TenantName = &tenantName
}
/* param clusterName: 集群字段(Optional) */
func (r *K8sconfigNewV2Request) SetClusterName(clusterName string) {
    r.ClusterName = &clusterName
}
/* param zoneName: 机房字段(Optional) */
func (r *K8sconfigNewV2Request) SetZoneName(zoneName string) {
    r.ZoneName = &zoneName
}
/* param nodeIp: node-ip字段(Optional) */
func (r *K8sconfigNewV2Request) SetNodeIp(nodeIp string) {
    r.NodeIp = &nodeIp
}
/* param conf: 采集配置字段(Optional) */
func (r *K8sconfigNewV2Request) SetConf(conf string) {
    r.Conf = &conf
}
/* param cloud: 是否为公有云(Optional) */
func (r *K8sconfigNewV2Request) SetCloud(cloud bool) {
    r.Cloud = &cloud
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r K8sconfigNewV2Request) GetRegionId() string {
    return ""
}

type K8sconfigNewV2Response struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result K8sconfigNewV2Result `json:"result"`
}

type K8sconfigNewV2Result struct {
    Result bool `json:"result"`
}