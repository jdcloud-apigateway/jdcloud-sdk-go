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


type InvoiceOrder struct {

    /* id (Optional) */
    Id int `json:"id"`

    /* 用户pin (Optional) */
    Pin string `json:"pin"`

    /* 单据编号 (Optional) */
    OrderId string `json:"orderId"`

    /* 单据类型[订单-1,账单-2] (Optional) */
    ReceiptType int `json:"receiptType"`

    /* 单据类型[订单-order,账单-bill] (Optional) */
    TypeDesc string `json:"typeDesc"`

    /* 产品名称 (Optional) */
    Name string `json:"name"`

    /* 实付金额 (Optional) */
    Amount int `json:"amount"`

    /* invoicedAmount (Optional) */
    InvoicedAmount int `json:"invoicedAmount"`

    /* 可开票金额 (Optional) */
    EnableAmount int `json:"enableAmount"`

    /* 已开票金额 (Optional) */
    UseAmount int `json:"useAmount"`

    /* 退款金额 (Optional) */
    RefundAmount int `json:"refundAmount"`

    /* 申请状态 (Optional) */
    Status int `json:"status"`

    /* 成单日期 (Optional) */
    FinishTime string `json:"finishTime"`

    /* 创建日期 (Optional) */
    CreateTime string `json:"createTime"`

    /* 更新日期 (Optional) */
    UpdateTime string `json:"updateTime"`

    /* chargeSubject (Optional) */
    ChargeSubject string `json:"chargeSubject"`

    /* batch (Optional) */
    Batch string `json:"batch"`

    /* 查询 开始日期 (Optional) */
    SearchStartTime string `json:"searchStartTime"`

    /* 查询 结束日期 (Optional) */
    SearchEndTime string `json:"searchEndTime"`

    /* 使用人 (Optional) */
    Owner string `json:"owner"`

    /* 交易(支付)类型 1-代付 2-自付 (Optional) */
    PayType int `json:"payType"`

    /* payType描述 1-代付 2-自付 (Optional) */
    PayTypeDesc string `json:"payTypeDesc"`

    /* 如果是月结订单，返回具体月份 例202301 (Optional) */
    MonthGroup string `json:"monthGroup"`
}
