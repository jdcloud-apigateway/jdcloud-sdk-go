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
    fc "github.com/jdcloud-api/jdcloud-sdk-go/services/fc/models"
)

type CreateServiceRequest struct {

    core.JDCloudRequest

    /* Region ID  */
    RegionId string `json:"regionId"`

    /* 服务 创建参数  */
    ServiceSpec *fc.ServiceSpec `json:"serviceSpec"`

    /* 保证请求幂等性的字符串；最大长度64个ASCII字符 (Optional) */
    ClientToken *string `json:"clientToken"`
}

/*
 * param regionId: Region ID (Required)
 * param serviceSpec: 服务 创建参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateServiceRequest(
    regionId string,
    serviceSpec *fc.ServiceSpec,
) *CreateServiceRequest {

	return &CreateServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/services",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        ServiceSpec: serviceSpec,
	}
}

/*
 * param regionId: Region ID (Required)
 * param serviceSpec: 服务 创建参数 (Required)
 * param clientToken: 保证请求幂等性的字符串；最大长度64个ASCII字符 (Optional)
 */
func NewCreateServiceRequestWithAllParams(
    regionId string,
    serviceSpec *fc.ServiceSpec,
    clientToken *string,
) *CreateServiceRequest {

    return &CreateServiceRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        ServiceSpec: serviceSpec,
        ClientToken: clientToken,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateServiceRequestWithoutParam() *CreateServiceRequest {

    return &CreateServiceRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/services",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: Region ID(Required) */
func (r *CreateServiceRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param serviceSpec: 服务 创建参数(Required) */
func (r *CreateServiceRequest) SetServiceSpec(serviceSpec *fc.ServiceSpec) {
    r.ServiceSpec = serviceSpec
}
/* param clientToken: 保证请求幂等性的字符串；最大长度64个ASCII字符(Optional) */
func (r *CreateServiceRequest) SetClientToken(clientToken string) {
    r.ClientToken = &clientToken
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateServiceRequest) GetRegionId() string {
    return r.RegionId
}

type CreateServiceResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateServiceResult `json:"result"`
}

type CreateServiceResult struct {
    ServiceName string `json:"serviceName"`
}