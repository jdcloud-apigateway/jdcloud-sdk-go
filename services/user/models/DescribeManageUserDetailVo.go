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


type DescribeManageUserDetailVo struct {

    /* 销售erp (Optional) */
    Seller string `json:"seller"`

    /* 登录态认证类型 (Optional) */
    ActiveDeviceType string `json:"activeDeviceType"`

    /* 计费状态 (Optional) */
    CostDay int `json:"costDay"`

    /* jdbu (Optional) */
    Jdbu string `json:"jdbu"`

    /* 网络专区 (Optional) */
    NetworkZone string `json:"networkZone"`

    /* 门槛价格 (Optional) */
    Threshold string `json:"threshold"`

    /* vip标签 (Optional) */
    VipInfo string `json:"vipInfo"`

    /* 商户号 (Optional) */
    MerchantId string `json:"merchantId"`

    /* 用户分组和实名名称 (Optional) */
    UserTypeAndName UserTypeAndNameVo `json:"userTypeAndName"`

    /* 企业账号信息 (Optional) */
    CloudUser CloudUserVo `json:"cloudUser"`

    /* 用户信息 (Optional) */
    UserInfo UserInfoVo `json:"userInfo"`

    /* 用户信息 (Optional) */
    BlacklistRemark BlacklistRemarkVo `json:"blacklistRemark"`

    /* 用户信息 (Optional) */
    ContactPersons ContactPersonsVo `json:"contactPersons"`
}
