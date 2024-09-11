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


type OpIncomeSummaryRes struct {

    /* 批量推送提交IDs (Optional) */
    BatchPushIds []int `json:"batchPushIds"`

    /* 主键id (Optional) */
    Id int `json:"id"`

    /* 是否推计提 (Optional) */
    EbsPush bool `json:"ebsPush"`

    /* 是否推账单 (Optional) */
    BillPush bool `json:"billPush"`

    /* 账单id (Optional) */
    DataSourceId string `json:"dataSourceId"`

    /* 操作类型，1、新增，2、更新 (Optional) */
    OpType int `json:"opType"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 产品serviceCode (Optional) */
    ServiceCode string `json:"serviceCode"`

    /* 计费类型，1、按配置，2、按用量，3、包年包月，4、一次性 (Optional) */
    BillingType int `json:"billingType"`

    /* 计费开始时间（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    StartTime string `json:"startTime"`

    /* 计费结束时间（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    EndTime string `json:"endTime"`

    /* 账期（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    BillTime string `json:"billTime"`

    /* 原价，保留6位小数 (Optional) */
    BillFee string `json:"billFee"`

    /* 实际价格 应付金额，保留2位小数 (Optional) */
    ActualFee string `json:"actualFee"`

    /* 折扣金额，保留2位小数 (Optional) */
    DiscountFee string `json:"discountFee"`

    /* 金额类型，5:电汇 (Optional) */
    AmountType int `json:"amountType"`

    /* 币种 (Optional) */
    PayCurrency string `json:"payCurrency"`

    /* 税率 (Optional) */
    TaxRate string `json:"taxRate"`

    /* 支付状态，0：未支付、1：已支付 (Optional) */
    PayState int `json:"payState"`

    /* 支付时间，必传（格式：yyyy-MM-dd HH:mm:ss） (Optional) */
    PayTime string `json:"payTime"`

    /* 线下签合同的主体id (Optional) */
    OrgId string `json:"orgId"`

    /* 线下签合同的主体名称 (Optional) */
    OrgName string `json:"orgName"`

    /* 1-新购，2-续费，3-配置变更，4-退款，42-降配延时 47-包年包月转换配置、用量 (Optional) */
    OrderType int `json:"orderType"`

    /* 账单类型，1-正常账单，2-退款账单，3-调账账单，4-保底账单 (Optional) */
    BillType int `json:"billType"`

    /* 调整账单需要原始账单ID (Optional) */
    OrigDataSourceId string `json:"origDataSourceId"`

    /* 销售合同编号 (Optional) */
    SalesContractNumber string `json:"salesContractNumber"`

    /* 销售合同状态 1、草稿，2、审批中，3、审批驳回，4、EBS审批中，5、待用印、6、待签订，7、签订中，8、已签订，9、不签订，11、已失效 (Optional) */
    SalesContractStatus int `json:"salesContractStatus"`

    /* 地域，如cn-north-1，不可修改 (Optional) */
    Region string `json:"region"`

    /* 配置描述（备注） (Optional) */
    FormulaDesc string `json:"formulaDesc"`
}
