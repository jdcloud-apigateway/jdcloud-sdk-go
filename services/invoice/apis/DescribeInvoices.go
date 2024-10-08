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

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    invoice "github.com/jdcloud-api/jdcloud-sdk-go/services/invoice/models"
)

type DescribeInvoicesRequest struct {

    core.JDCloudRequest

    /* 地域编码，参考OpenAPI公共说明  */
    RegionId string `json:"regionId"`

    /* 发票状态[已申请-applied ，处理中-processing ，已开票-invoiced ，已邮寄-mailed ，已驳回-dismissed ，已作废-obsolete ，已取消-cancelled，退票中-refund，已退票-refunded，退票驳回-refund_rejected] (Optional) */
    Status *string `json:"status"`

    /* 开始时间 (Optional) */
    SearchStartDate *string `json:"searchStartDate"`

    /* 结束时间 (Optional) */
    SearchEndDate *string `json:"searchEndDate"`

    /* 页码 (Optional) */
    PageNumber *int `json:"pageNumber"`

    /* 页大小 (Optional) */
    PageSize *int `json:"pageSize"`

    /* 按明细开票 1 按月结算单开票 2 按指定金额开票 3 (Optional) */
    InvoiceType *int `json:"invoiceType"`
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewDescribeInvoicesRequest(
    regionId string,
) *DescribeInvoicesRequest {

	return &DescribeInvoicesRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/invoices",
			Method:  "POST",
			Header:  nil,
			Version: "v2",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 地域编码，参考OpenAPI公共说明 (Required)
 * param status: 发票状态[已申请-applied ，处理中-processing ，已开票-invoiced ，已邮寄-mailed ，已驳回-dismissed ，已作废-obsolete ，已取消-cancelled，退票中-refund，已退票-refunded，退票驳回-refund_rejected] (Optional)
 * param searchStartDate: 开始时间 (Optional)
 * param searchEndDate: 结束时间 (Optional)
 * param pageNumber: 页码 (Optional)
 * param pageSize: 页大小 (Optional)
 * param invoiceType: 按明细开票 1 按月结算单开票 2 按指定金额开票 3 (Optional)
 */
func NewDescribeInvoicesRequestWithAllParams(
    regionId string,
    status *string,
    searchStartDate *string,
    searchEndDate *string,
    pageNumber *int,
    pageSize *int,
    invoiceType *int,
) *DescribeInvoicesRequest {

    return &DescribeInvoicesRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoices",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
        RegionId: regionId,
        Status: status,
        SearchStartDate: searchStartDate,
        SearchEndDate: searchEndDate,
        PageNumber: pageNumber,
        PageSize: pageSize,
        InvoiceType: invoiceType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewDescribeInvoicesRequestWithoutParam() *DescribeInvoicesRequest {

    return &DescribeInvoicesRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/invoices",
            Method:  "POST",
            Header:  nil,
            Version: "v2",
        },
    }
}

/* param regionId: 地域编码，参考OpenAPI公共说明(Required) */
func (r *DescribeInvoicesRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}
/* param status: 发票状态[已申请-applied ，处理中-processing ，已开票-invoiced ，已邮寄-mailed ，已驳回-dismissed ，已作废-obsolete ，已取消-cancelled，退票中-refund，已退票-refunded，退票驳回-refund_rejected](Optional) */
func (r *DescribeInvoicesRequest) SetStatus(status string) {
    r.Status = &status
}
/* param searchStartDate: 开始时间(Optional) */
func (r *DescribeInvoicesRequest) SetSearchStartDate(searchStartDate string) {
    r.SearchStartDate = &searchStartDate
}
/* param searchEndDate: 结束时间(Optional) */
func (r *DescribeInvoicesRequest) SetSearchEndDate(searchEndDate string) {
    r.SearchEndDate = &searchEndDate
}
/* param pageNumber: 页码(Optional) */
func (r *DescribeInvoicesRequest) SetPageNumber(pageNumber int) {
    r.PageNumber = &pageNumber
}
/* param pageSize: 页大小(Optional) */
func (r *DescribeInvoicesRequest) SetPageSize(pageSize int) {
    r.PageSize = &pageSize
}
/* param invoiceType: 按明细开票 1 按月结算单开票 2 按指定金额开票 3(Optional) */
func (r *DescribeInvoicesRequest) SetInvoiceType(invoiceType int) {
    r.InvoiceType = &invoiceType
}


// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r DescribeInvoicesRequest) GetRegionId() string {
    return r.RegionId
}

type DescribeInvoicesResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result DescribeInvoicesResult `json:"result"`
}

type DescribeInvoicesResult struct {
    EnableInvoiceFee int `json:"enableInvoiceFee"`
    InvoicedFee int `json:"invoicedFee"`
    IsSetInvoiceMSGTemplate bool `json:"isSetInvoiceMSGTemplate"`
    InvoiceList invoice.InvoiceList `json:"invoiceList"`
    IsApplyInvoice bool `json:"isApplyInvoice"`
    InvoiceSearch invoice.Invoices `json:"invoiceSearch"`
}