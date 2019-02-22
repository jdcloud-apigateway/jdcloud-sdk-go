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

type EncryptRequest struct {

    core.JDCloudRequest

    /* 密钥ID  */
    KeyId string `json:"keyId"`

    /* 明文数据 Base64-encoded binary data object (Optional) */
    Plaintext *string `json:"plaintext"`
}

/*
 * param keyId: 密钥ID (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewEncryptRequest(
    keyId string,
) *EncryptRequest {

	return &EncryptRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/key/{keyId}:Encrypt",
			Method:  "POST",
			Header:  nil,
			Version: "v1",
		},
        KeyId: keyId,
	}
}

/*
 * param keyId: 密钥ID (Required)
 * param plaintext: 明文数据 Base64-encoded binary data object (Optional)
 */
func NewEncryptRequestWithAllParams(
    keyId string,
    plaintext *string,
) *EncryptRequest {

    return &EncryptRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:Encrypt",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
        KeyId: keyId,
        Plaintext: plaintext,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewEncryptRequestWithoutParam() *EncryptRequest {

    return &EncryptRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/key/{keyId}:Encrypt",
            Method:  "POST",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param keyId: 密钥ID(Required) */
func (r *EncryptRequest) SetKeyId(keyId string) {
    r.KeyId = keyId
}

/* param plaintext: 明文数据 Base64-encoded binary data object(Optional) */
func (r *EncryptRequest) SetPlaintext(plaintext string) {
    r.Plaintext = &plaintext
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r EncryptRequest) GetRegionId() string {
    return ""
}

type EncryptResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result EncryptResult `json:"result"`
}

type EncryptResult struct {
    CiphertextBlob string `json:"ciphertextBlob"`
}