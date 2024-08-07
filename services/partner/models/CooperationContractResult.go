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


type CooperationContractResult struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 合作ID (Optional) */
    CooperationId string `json:"cooperationId"`

    /* UUID (Optional) */
    Uuid string `json:"uuid"`

    /* 合作名称 (Optional) */
    Name string `json:"name"`

    /* CRM合同编号 (Optional) */
    CrmContractNo string `json:"crmContractNo"`

    /* CRM合同名称 (Optional) */
    CrmContractName string `json:"crmContractName"`

    /* 协议类型编码 (Optional) */
    TypeCode string `json:"typeCode"`

    /* 协议类型名称 (Optional) */
    TypeName string `json:"typeName"`

    /* EBS合同编号 (Optional) */
    EbsContractNo string `json:"ebsContractNo"`

    /* EBS主合同编号 (Optional) */
    EbsMainNo string `json:"ebsMainNo"`

    /* 合同签署日期 (Optional) */
    SignedDate string `json:"signedDate"`

    /* 合同生效时间 (Optional) */
    EffectiveDate string `json:"effectiveDate"`

    /* 合同失效时间 (Optional) */
    ExpirationDate string `json:"expirationDate"`

    /* 合同提交人 (Optional) */
    Submitter string `json:"submitter"`

    /* 是否补充协议，值为：Y/N (Optional) */
    Bcagreement string `json:"bcagreement"`

    /* 合同性质：1、收款，2、付款 (Optional) */
    Xz string `json:"xz"`

    /* 合同总额 (Optional) */
    Amount int `json:"amount"`

    /* 备注 (Optional) */
    Remark string `json:"remark"`

    /* createTime (Optional) */
    CreateTime string `json:"createTime"`

    /* 公司名称 (Optional) */
    CompanyName string `json:"companyName"`

    /* 合作类型 (Optional) */
    PartnerType int `json:"partnerType"`

    /* 伙伴类型名称 (Optional) */
    PartnerTypeName string `json:"partnerTypeName"`

    /* 合作模式 (Optional) */
    CooperationModel int `json:"cooperationModel"`

    /* 合作模式名称 (Optional) */
    CooperationModelName string `json:"cooperationModelName"`

    /* 合作部门 (Optional) */
    OrgName string `json:"orgName"`

    /* 负责人erp (Optional) */
    Erp string `json:"erp"`

    /* 有效期 (Optional) */
    ContractDateName string `json:"contractDateName"`

    /* 协议状态 (Optional) */
    StatusName string `json:"statusName"`

    /* 是否补充协议，是或者否 (Optional) */
    BcContract string `json:"bcContract"`

    /* 交易方 (Optional) */
    PartiesInfoList []PartiesInfoDTO `json:"partiesInfoList"`

    /* 协议文档 (Optional) */
    CooperationContractPdfList []CooperationContractPdfDTO `json:"cooperationContractPdfList"`
}
