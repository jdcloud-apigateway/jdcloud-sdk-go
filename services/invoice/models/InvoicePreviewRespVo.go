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


type InvoicePreviewRespVo struct {

    /* 单位名称/用户的发票抬头信息 (Optional) */
    CompanyName string `json:"companyName"`

    /* 纳税人识别号;个人类型时为空 (Optional) */
    CompanyTaxCode string `json:"companyTaxCode"`

    /* 公司地址;个人类型时为空 (Optional) */
    CompanyAddress string `json:"companyAddress"`

    /* 公司电话/手机号 (Optional) */
    CompanyTelphone string `json:"companyTelphone"`

    /* 开户行;个人类型时为空 (Optional) */
    CompanyBank string `json:"companyBank"`

    /* 开户行的银行账户;个人类型时为空 (Optional) */
    CompanyBankAccount string `json:"companyBankAccount"`

    /* 货物或应税劳务 、服务名称 (Optional) */
    ServiceName string `json:"serviceName"`

    /* 规格 (Optional) */
    Specification string `json:"specification"`

    /* 发票总金额 (Optional) */
    TotalAmount int `json:"totalAmount"`

    /* 发票不含税金额 (Optional) */
    Amount int `json:"amount"`

    /* 税率 (Optional) */
    TaxRate int `json:"taxRate"`

    /* 税额 (Optional) */
    TaxAmount int `json:"taxAmount"`

    /* 税额中文 (Optional) */
    TaxAmountCnStr string `json:"taxAmountCnStr"`

    /* 销售主体 (Optional) */
    ChargeSubject string `json:"chargeSubject"`

    /* 销售主体识别码 (Optional) */
    ChargeSubjectTaxCode string `json:"chargeSubjectTaxCode"`

    /* 销售主体地址 (Optional) */
    ChargeSubjectAddress string `json:"chargeSubjectAddress"`

    /* 销售主体电话 (Optional) */
    ChargeSubjectTelephone string `json:"chargeSubjectTelephone"`

    /* 销售主体开户行 (Optional) */
    ChargeSubjectBank string `json:"chargeSubjectBank"`

    /* 销售主体开户行账号 (Optional) */
    ChargeSubjectBankAccount string `json:"chargeSubjectBankAccount"`

    /* 内容 (Optional) */
    Content string `json:"content"`

    /* 预览信息的唯一标识 (Optional) */
    Uuid string `json:"uuid"`
}
