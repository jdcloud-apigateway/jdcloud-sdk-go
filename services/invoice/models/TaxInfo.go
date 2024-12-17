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


type TaxInfo struct {

    /* 主键 (Optional) */
    Id int `json:"id"`

    /* 产品线 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 产品线名称 (Optional) */
    ServiceName string `json:"serviceName"`

    /* 税率(小数 如 0.06) (Optional) */
    TaxRa int `json:"taxRa"`

    /* 税号 (Optional) */
    TaxCode string `json:"taxCode"`

    /* 开票内容 (Optional) */
    Content string `json:"content"`

    /* 收费主体 (Optional) */
    ChargeSubject string `json:"chargeSubject"`
}