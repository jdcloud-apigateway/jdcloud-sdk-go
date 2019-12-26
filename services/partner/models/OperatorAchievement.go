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


type OperatorAchievement struct {

    /* pin (Optional) */
    Pin string `json:"pin"`

    /* 名称 (Optional) */
    Name string `json:"name"`

    /* 真实名称 (Optional) */
    RealName string `json:"realName"`

    /* 一级服务商名称 (Optional) */
    OneLevelDistributorName string `json:"oneLevelDistributorName"`

    /* 二级服务商名称 (Optional) */
    SecondLevelDistributorName string `json:"secondLevelDistributorName"`

    /* 部门 (Optional) */
    Dept int `json:"dept"`

    /* 部门名称 (Optional) */
    DeptName string `json:"deptName"`

    /* 服务商类型 (Optional) */
    DistributorType string `json:"distributorType"`

    /* 服务商类型名称 (Optional) */
    DistributorTypeName string `json:"distributorTypeName"`

    /* 上级服务商名称 (Optional) */
    SuperDistributorName string `json:"superDistributorName"`

    /* 上级服务商pin (Optional) */
    SuperDistributorPin string `json:"superDistributorPin"`

    /* serviceCode (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* serviceCodeName (Optional) */
    ServiceCodeName string `json:"serviceCodeName"`

    /* 帐户名 (Optional) */
    LoginName string `json:"loginName"`

    /* 时间 (Optional) */
    ConsumeDate string `json:"consumeDate"`

    /* 优惠前金额 (Optional) */
    ConsumeCount int `json:"consumeCount"`

    /* 现金支付 (Optional) */
    CashPayFeeCount int `json:"cashPayFeeCount"`

    /* 业绩金额 (Optional) */
    AchievmentCount int `json:"achievmentCount"`

    /* 付费代金卷金额 (Optional) */
    PayCouponFeeCount int `json:"payCouponFeeCount"`

    /* 免费代金卷金额 (Optional) */
    FreeCouponFeeCount int `json:"freeCouponFeeCount"`
}
