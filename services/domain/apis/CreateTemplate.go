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
)

type CreateTemplateRequest struct {

    core.JDCloudRequest

    /* 实例所属的地域ID  */
    RegionId string `json:"regionId"`

    /* 联系人姓名(中文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional) */
    UserNameCh *string `json:"userNameCh"`

    /* 联系人姓(英文),必填,必须含有英文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional) */
    UserNameEn1 *string `json:"userNameEn1"`

    /* 联系人名(英文),必填,必须含有英文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional) */
    UserNameEn2 *string `json:"userNameEn2"`

    /* 域名所有者或所有者单位名称(中文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional) */
    OwnerNameCh *string `json:"ownerNameCh"`

    /* 域名所有者或所有者单位名称(英文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional) */
    OwnerNameEn *string `json:"ownerNameEn"`

    /* 国家及地区（中文） (Optional) */
    NationCodeCh *string `json:"nationCodeCh"`

    /* 国家及地区（英文） (Optional) */
    NationCodeEn *string `json:"nationCodeEn"`

    /* 省份（中文） (Optional) */
    ProvinceCodeCh *string `json:"provinceCodeCh"`

    /* 省份（英文） (Optional) */
    ProvinceCodeEn *string `json:"provinceCodeEn"`

    /* 城市（中文） (Optional) */
    CityCodeCh *string `json:"cityCodeCh"`

    /* 城市（英文） (Optional) */
    CityCodeEn *string `json:"cityCodeEn"`

    /* 通信地址（中文） (Optional) */
    AddressCh *string `json:"addressCh"`

    /* 通信地址（英文） (Optional) */
    AddressEn *string `json:"addressEn"`

    /* 邮编 6位数字 (Optional) */
    ZipCode *string `json:"zipCode"`

    /* 联系电话，国家区号-地区区号(或手机号码前3位)-电话号码（或手机号码后8位) 例:86-138-12345678 (Optional) */
    Phone *string `json:"phone"`

    /* 传真，国家区号-地区区号(或手机号码前3位)-电话号码（或手机号码后8位) 例:86-138-12345678 (Optional) */
    Fax *string `json:"fax"`

    /* 邮件 (Optional) */
    Email *string `json:"email"`

    /* 所有者类型  1个人 2企业 (Optional) */
    OwnerType *int `json:"ownerType"`
}

/*
 * param regionId: 实例所属的地域ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewCreateTemplateRequest(
    regionId string,
) *CreateTemplateRequest {

	return &CreateTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/template",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
	}
}

/*
 * param regionId: 实例所属的地域ID (Required)
 * param userNameCh: 联系人姓名(中文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional)
 * param userNameEn1: 联系人姓(英文),必填,必须含有英文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional)
 * param userNameEn2: 联系人名(英文),必填,必须含有英文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional)
 * param ownerNameCh: 域名所有者或所有者单位名称(中文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional)
 * param ownerNameEn: 域名所有者或所有者单位名称(英文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符 (Optional)
 * param nationCodeCh: 国家及地区（中文） (Optional)
 * param nationCodeEn: 国家及地区（英文） (Optional)
 * param provinceCodeCh: 省份（中文） (Optional)
 * param provinceCodeEn: 省份（英文） (Optional)
 * param cityCodeCh: 城市（中文） (Optional)
 * param cityCodeEn: 城市（英文） (Optional)
 * param addressCh: 通信地址（中文） (Optional)
 * param addressEn: 通信地址（英文） (Optional)
 * param zipCode: 邮编 6位数字 (Optional)
 * param phone: 联系电话，国家区号-地区区号(或手机号码前3位)-电话号码（或手机号码后8位) 例:86-138-12345678 (Optional)
 * param fax: 传真，国家区号-地区区号(或手机号码前3位)-电话号码（或手机号码后8位) 例:86-138-12345678 (Optional)
 * param email: 邮件 (Optional)
 * param ownerType: 所有者类型  1个人 2企业 (Optional)
 */
func NewCreateTemplateRequestWithAllParams(
    regionId string,
    userNameCh *string,
    userNameEn1 *string,
    userNameEn2 *string,
    ownerNameCh *string,
    ownerNameEn *string,
    nationCodeCh *string,
    nationCodeEn *string,
    provinceCodeCh *string,
    provinceCodeEn *string,
    cityCodeCh *string,
    cityCodeEn *string,
    addressCh *string,
    addressEn *string,
    zipCode *string,
    phone *string,
    fax *string,
    email *string,
    ownerType *int,
) *CreateTemplateRequest {

    return &CreateTemplateRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        UserNameCh: userNameCh,
        UserNameEn1: userNameEn1,
        UserNameEn2: userNameEn2,
        OwnerNameCh: ownerNameCh,
        OwnerNameEn: ownerNameEn,
        NationCodeCh: nationCodeCh,
        NationCodeEn: nationCodeEn,
        ProvinceCodeCh: provinceCodeCh,
        ProvinceCodeEn: provinceCodeEn,
        CityCodeCh: cityCodeCh,
        CityCodeEn: cityCodeEn,
        AddressCh: addressCh,
        AddressEn: addressEn,
        ZipCode: zipCode,
        Phone: phone,
        Fax: fax,
        Email: email,
        OwnerType: ownerType,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewCreateTemplateRequestWithoutParam() *CreateTemplateRequest {

    return &CreateTemplateRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/template",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 实例所属的地域ID(Required) */
func (r *CreateTemplateRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param userNameCh: 联系人姓名(中文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符(Optional) */
func (r *CreateTemplateRequest) SetUserNameCh(userNameCh string) {
    r.UserNameCh = &userNameCh
}

/* param userNameEn1: 联系人姓(英文),必填,必须含有英文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符(Optional) */
func (r *CreateTemplateRequest) SetUserNameEn1(userNameEn1 string) {
    r.UserNameEn1 = &userNameEn1
}

/* param userNameEn2: 联系人名(英文),必填,必须含有英文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符(Optional) */
func (r *CreateTemplateRequest) SetUserNameEn2(userNameEn2 string) {
    r.UserNameEn2 = &userNameEn2
}

/* param ownerNameCh: 域名所有者或所有者单位名称(中文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符(Optional) */
func (r *CreateTemplateRequest) SetOwnerNameCh(ownerNameCh string) {
    r.OwnerNameCh = &ownerNameCh
}

/* param ownerNameEn: 域名所有者或所有者单位名称(英文),必填,必须含有中文,只允许输入特殊字符(.,、·()()-""“”/\'),最多可输入64 字符(Optional) */
func (r *CreateTemplateRequest) SetOwnerNameEn(ownerNameEn string) {
    r.OwnerNameEn = &ownerNameEn
}

/* param nationCodeCh: 国家及地区（中文）(Optional) */
func (r *CreateTemplateRequest) SetNationCodeCh(nationCodeCh string) {
    r.NationCodeCh = &nationCodeCh
}

/* param nationCodeEn: 国家及地区（英文）(Optional) */
func (r *CreateTemplateRequest) SetNationCodeEn(nationCodeEn string) {
    r.NationCodeEn = &nationCodeEn
}

/* param provinceCodeCh: 省份（中文）(Optional) */
func (r *CreateTemplateRequest) SetProvinceCodeCh(provinceCodeCh string) {
    r.ProvinceCodeCh = &provinceCodeCh
}

/* param provinceCodeEn: 省份（英文）(Optional) */
func (r *CreateTemplateRequest) SetProvinceCodeEn(provinceCodeEn string) {
    r.ProvinceCodeEn = &provinceCodeEn
}

/* param cityCodeCh: 城市（中文）(Optional) */
func (r *CreateTemplateRequest) SetCityCodeCh(cityCodeCh string) {
    r.CityCodeCh = &cityCodeCh
}

/* param cityCodeEn: 城市（英文）(Optional) */
func (r *CreateTemplateRequest) SetCityCodeEn(cityCodeEn string) {
    r.CityCodeEn = &cityCodeEn
}

/* param addressCh: 通信地址（中文）(Optional) */
func (r *CreateTemplateRequest) SetAddressCh(addressCh string) {
    r.AddressCh = &addressCh
}

/* param addressEn: 通信地址（英文）(Optional) */
func (r *CreateTemplateRequest) SetAddressEn(addressEn string) {
    r.AddressEn = &addressEn
}

/* param zipCode: 邮编 6位数字(Optional) */
func (r *CreateTemplateRequest) SetZipCode(zipCode string) {
    r.ZipCode = &zipCode
}

/* param phone: 联系电话，国家区号-地区区号(或手机号码前3位)-电话号码（或手机号码后8位) 例:86-138-12345678(Optional) */
func (r *CreateTemplateRequest) SetPhone(phone string) {
    r.Phone = &phone
}

/* param fax: 传真，国家区号-地区区号(或手机号码前3位)-电话号码（或手机号码后8位) 例:86-138-12345678(Optional) */
func (r *CreateTemplateRequest) SetFax(fax string) {
    r.Fax = &fax
}

/* param email: 邮件(Optional) */
func (r *CreateTemplateRequest) SetEmail(email string) {
    r.Email = &email
}

/* param ownerType: 所有者类型  1个人 2企业(Optional) */
func (r *CreateTemplateRequest) SetOwnerType(ownerType int) {
    r.OwnerType = &ownerType
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r CreateTemplateRequest) GetRegionId() string {
    return r.RegionId
}

type CreateTemplateResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result CreateTemplateResult `json:"result"`
}

type CreateTemplateResult struct {
    TemplateId int64 `json:"templateId"`
}