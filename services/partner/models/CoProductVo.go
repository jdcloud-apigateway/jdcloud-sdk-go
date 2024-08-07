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


type CoProductVo struct {

    /* 合作id (Optional) */
    CooperationId string `json:"cooperationId"`

    /* 合作产品名称 (Optional) */
    ProductName string `json:"productName"`

    /* 合作公司名称 (Optional) */
    CompanyName string `json:"companyName"`

    /* 合作名称 (Optional) */
    Name string `json:"name"`

    /* 合同信息列表 (Optional) */
    ContractInfos []Contract `json:"contractInfos"`

    /* 产品类型 (Optional) */
    ProductType int `json:"productType"`

    /* 产品类型名称 (Optional) */
    ProductTypeName string `json:"productTypeName"`

    /* 产品模式 (Optional) */
    ProductMode int `json:"productMode"`

    /* 产品模式显示名称 (Optional) */
    ProductModeName string `json:"productModeName"`

    /* 合作对接人(原合作BD) (Optional) */
    Broker string `json:"broker"`

    /* 结算方式 1固定金额结算，2实际售价固定比例结算，3实际售价浮动比例结算 (Optional) */
    SettlementMode int `json:"settlementMode"`

    /* 结算方式显示名称 (Optional) */
    SettlementModeName string `json:"settlementModeName"`

    /* 结算周期 1周结后付款、2月结后付款、3季结后付款、4年结后付款，5 PO预付款 (Optional) */
    SettlementCyle int `json:"settlementCyle"`

    /* 结算周期显示名称 (Optional) */
    SettlementCyleName string `json:"settlementCyleName"`

    /* 产品状态  1联合开发中；2实施部署中；3测试验证中；4可商用售卖；5可PO下单；6需重新采购；7合作暂停 (Optional) */
    ProductStatus int `json:"productStatus"`

    /* 产品状态显示名称 (Optional) */
    ProductStatusName string `json:"productStatusName"`

    /* 合作公司 (Optional) */
    CooperationCompanys []CooperationCompanyVo `json:"cooperationCompanys"`
}
