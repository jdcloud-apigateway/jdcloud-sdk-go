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


type AuthStatusVo struct {

    /* 银行账号 (Optional) */
    BankAccount string `json:"bankAccount"`

    /* 开户行编号 (Optional) */
    BankCode string `json:"bankCode"`

    /* 是否扫脸实名（已废弃） (Optional) */
    LegalFace int `json:"legalFace"`

    /* 个人实名方式（描述） (Optional) */
    PersonAuthWay string `json:"personAuthWay"`

    /* 是否弱实名0否，1是 (Optional) */
    WeakAuth int `json:"weakAuth"`

    /* 是否需要授权 0否，1是 (Optional) */
    AuthorizeStatus int `json:"authorizeStatus"`

    /* 开户行名称 (Optional) */
    BankName string `json:"bankName"`

    /* 企业实名状态 (Optional) */
    CompanyAuthStatus int `json:"companyAuthStatus"`

    /* 个人实名状态 (Optional) */
    PersonAuthStatus int `json:"personAuthStatus"`

    /* 企业实名方式（描述） (Optional) */
    CompanyAuthWay string `json:"companyAuthWay"`

    /* 认证方式（企业） (Optional) */
    LegalAuthType int `json:"legalAuthType"`

    /* 更新实名认证状态1认证中，0未认证 (Optional) */
    AuthType int `json:"authType"`

    /* 是否允许更新认证false不允许 (Optional) */
    SupportUpdateCompanyAuth bool `json:"supportUpdateCompanyAuth"`

    /* 冻结时间 (Optional) */
    FrozenTime string `json:"frozenTime"`

    /* 是否扫脸实名（已废弃） (Optional) */
    LegalAuth bool `json:"legalAuth"`

    /* 驳回原因 (Optional) */
    RejectReason string `json:"rejectReason"`

    /* 京东金融企业实名状态（-1 为认证 1 审核中 2 被驳回 3 未完成认证 4 认证通过） (Optional) */
    FromJDWallet int `json:"fromJDWallet"`

    /* 京东金融的个人认证状态（-1 未认证 1 已认证） (Optional) */
    FromFinance int `json:"fromFinance"`

    /* 提交方式0：接口提交，1：组件页面 (Optional) */
    IntegrationType int `json:"integrationType"`

    /* 企业实名信息 (Optional) */
    CompanyAuthInfo CompanyAuthInfoVo `json:"companyAuthInfo"`

    /* 个人实名信息 (Optional) */
    PersonAuthInfo PersonAuthInfoVo `json:"personAuthInfo"`
}