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


type FormulaBillSummary struct {

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 资源id (Optional) */
    ResourceId string `json:"resourceId"`

    /* 产品线代码 (Optional) */
    AppCode string `json:"appCode"`

    /* 产品线代码名称 (Optional) */
    AppCodeName string `json:"appCodeName"`

    /* 产品代码 (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 产品代码名称 (Optional) */
    ServiceCodeName string `json:"serviceCodeName"`

    /* 区域 (Optional) */
    Region string `json:"region"`

    /* 区域名称 (Optional) */
    RegionName string `json:"regionName"`

    /* 币种 (Optional) */
    Currency string `json:"currency"`

    /* 计费类型 (Optional) */
    BillingType int `json:"billingType"`

    /* 计费类型描述 (Optional) */
    BillingTypeName string `json:"billingTypeName"`

    /* 网络类型 (Optional) */
    NetworkOperator string `json:"networkOperator"`

    /* 属性 (Optional) */
    Property string `json:"property"`

    /* 属性名称 (Optional) */
    PropertyName string `json:"propertyName"`

    /* 属性单位 (Optional) */
    PropertyUnit string `json:"propertyUnit"`

    /* 计费开始时间 (Optional) */
    StartTime string `json:"startTime"`

    /* 计费结束时间 (Optional) */
    EndTime string `json:"endTime"`

    /* 出账月份 (Optional) */
    BillDate string `json:"billDate"`

    /* 实际用量 (Optional) */
    ActualUsage float64 `json:"actualUsage"`

    /* 抵扣用量 (Optional) */
    DeductUsage float64 `json:"deductUsage"`

    /* 计费用量 (Optional) */
    BillingUsage float64 `json:"billingUsage"`

    /* 优惠前费用 (Optional) */
    BillFee float64 `json:"billFee"`

    /* 优惠金额 (Optional) */
    DiscountFee float64 `json:"discountFee"`

    /* 应付金额 (Optional) */
    ActualFee float64 `json:"actualFee"`

    /* 代金券支付金额 (Optional) */
    CashCouponPayFee float64 `json:"cashCouponPayFee"`

    /* 余额支付金额 (Optional) */
    BalancePayFee float64 `json:"balancePayFee"`

    /* 现金支付金额 (Optional) */
    CashPayFee float64 `json:"cashPayFee"`
}