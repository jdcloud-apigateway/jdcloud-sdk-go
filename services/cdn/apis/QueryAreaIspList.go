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
    cdn "github.com/jdcloud-api/jdcloud-sdk-go/services/cdn/models"
)

type QueryAreaIspListRequest struct {

    core.JDCloudRequest
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewQueryAreaIspListRequest(
) *QueryAreaIspListRequest {

	return &QueryAreaIspListRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/console:areaIspList",
			Method:  "GET",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 */
func NewQueryAreaIspListRequestWithAllParams(
) *QueryAreaIspListRequest {

    return &QueryAreaIspListRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:areaIspList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewQueryAreaIspListRequestWithoutParam() *QueryAreaIspListRequest {

    return &QueryAreaIspListRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/console:areaIspList",
            Method:  "GET",
            Header:  nil,
            Version: "v1",
        },
    }
}



// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r QueryAreaIspListRequest) GetRegionId() string {
    return ""
}

type QueryAreaIspListResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result QueryAreaIspListResult `json:"result"`
}

type QueryAreaIspListResult struct {
    MainLand []cdn.AreaIspItem `json:"mainLand"`
    Overseas []cdn.AreaIspItem `json:"overseas"`
    Isp []cdn.AreaIspItem `json:"isp"`
    Gangaotai []cdn.AreaIspItem `json:"gangaotai"`
    Oceanica []cdn.AreaIspItem `json:"oceanica"`
    SouthAmerica []cdn.AreaIspItem `json:"southAmerica"`
    NorthAmerica []cdn.AreaIspItem `json:"northAmerica"`
    Asia []cdn.AreaIspItem `json:"asia"`
    Europe []cdn.AreaIspItem `json:"europe"`
}