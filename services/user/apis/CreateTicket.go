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

type CreateTicketRequest struct {

    core.JDCloudRequest

    /* 有效期（默认24，最小1，最大24，单位小时） (Optional) */
    Expire *int64 `json:"expire"`
}

/*
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTicketRequest(
) *CreateTicketRequest {

	return &CreateTicketRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/user:createTicket",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
	}
}

/*
 * param expire: 有效期（默认24，最小1，最大24，单位小时） (Optional)
 */
func NewCreateTicketRequestWithAllParams(
    expire *int64,
) *CreateTicketRequest {

    return &CreateTicketRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/user:createTicket",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        Expire: expire,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTicketRequestWithoutParam() *CreateTicketRequest {

    return &CreateTicketRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/user:createTicket",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param expire: 有效期（默认24，最小1，最大24，单位小时）(Optional) */
func (r *CreateTicketRequest) SetExpire(expire int64) {
    r.Expire = &expire
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTicketRequest) GetRegionId() string {
    return ""
}

type CreateTicketResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTicketResult `json:"result"`
}

type CreateTicketResult struct {
    Ticket string `json:"ticket"`
}