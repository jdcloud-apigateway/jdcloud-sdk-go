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

package models


type ResourceRefundReq struct {

    /* 退款类型,退资源：REFUND_RESOURCE;退续费订单：REFUND_RENEW  */
    RefundType string `json:"refundType"`

    /* 退款请求的唯一标示  */
    RefId string `json:"refId"`

    /* 操作者类型  0:用户 1:运营人员  */
    OperatorType int `json:"operatorType"`

    /* 操作人，用户时传pin; 运营人员传操作人erp;  */
    OperatorName string `json:"operatorName"`

    /* 退款资源列表  */
    ResourceList []ResourceInfo `json:"resourceList"`
}
