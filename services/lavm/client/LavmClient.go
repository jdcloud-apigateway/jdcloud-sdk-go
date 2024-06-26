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

package client

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    lavm "github.com/jdcloud-api/jdcloud-sdk-go/services/lavm/apis"
    "encoding/json"
    "errors"
)

type LavmClient struct {
    core.JDCloudClient
}

func NewLavmClient(credential *core.Credential) *LavmClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("lavm.jdcloud-api.com")

    return &LavmClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "lavm",
            Revision:    "1.3.1",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *LavmClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *LavmClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *LavmClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 启动轻量应用云主机。
 */
func (c *LavmClient) StartInstance(request *lavm.StartInstanceRequest) (*lavm.StartInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.StartInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询域名接口
 */
func (c *LavmClient) DescribeDomains(request *lavm.DescribeDomainsRequest) (*lavm.DescribeDomainsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeDomainsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 
批量查询密钥对。

详细操作说明请参考帮助文档：[密钥概述](https://docs.jdcloud.com/cn/virtual-machines/keypair-overview)

## 接口说明
- 使用 `filters` 过滤器进行条件筛选，每个 `filter` 之间的关系为逻辑与（AND）的关系。
- 单次查询最大可查询100条密钥数据。
 */
func (c *LavmClient) DescribeKeypairs(request *lavm.DescribeKeypairsRequest) (*lavm.DescribeKeypairsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeKeypairsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改轻量应用云主机的属性信息。
 */
func (c *LavmClient) UpdateInstanceAttribute(request *lavm.UpdateInstanceAttributeRequest) (*lavm.UpdateInstanceAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.UpdateInstanceAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 升级轻量应用云主机。
 */
func (c *LavmClient) UpgradeInstance(request *lavm.UpgradeInstanceRequest) (*lavm.UpgradeInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.UpgradeInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 修改镜像属性。
详细操作说明请参考帮助文档：[镜像概述](https://docs.jdcloud.com/cn/virtual-machines/image-overview)
## 接口说明
- 只支持修改镜像名称或描述。
 */
func (c *LavmClient) ModifyImageAttribute(request *lavm.ModifyImageAttributeRequest) (*lavm.ModifyImageAttributeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.ModifyImageAttributeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 为指定的轻量应用云主机创建一条防火墙规则。
 */
func (c *LavmClient) CreateFirewallRule(request *lavm.CreateFirewallRuleRequest) (*lavm.CreateFirewallRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.CreateFirewallRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 停止轻量应用云主机。
 */
func (c *LavmClient) StopInstance(request *lavm.StopInstanceRequest) (*lavm.StopInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.StopInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 为轻量云主机绑定域名。
 */
func (c *LavmClient) AssociateDomains(request *lavm.AssociateDomainsRequest) (*lavm.AssociateDomainsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.AssociateDomainsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 轻量应用云主机重置系统。
 */
func (c *LavmClient) ResetSystem(request *lavm.ResetSystemRequest) (*lavm.ResetSystemResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.ResetSystemResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询一个或多个实例流量包详细信息。
 */
func (c *LavmClient) DescribeInstancesTrafficPackages(request *lavm.DescribeInstancesTrafficPackagesRequest) (*lavm.DescribeInstancesTrafficPackagesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeInstancesTrafficPackagesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 
导入密钥。

与创建密钥不同的是，导入的密钥是由用户生成的。生成之后将公钥部分导入到京东云。

详细操作说明请参考帮助文档：[创建密钥](https://docs.jdcloud.com/cn/virtual-machines/create-keypair)

## 接口说明
- 调用该接口导入由其他工具生成的密钥对的公钥部分。
 */
func (c *LavmClient) ImportKeypair(request *lavm.ImportKeypairRequest) (*lavm.ImportKeypairResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.ImportKeypairResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 为轻量云主机解绑域名。
 */
func (c *LavmClient) DisassociateDomains(request *lavm.DisassociateDomainsRequest) (*lavm.DisassociateDomainsResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DisassociateDomainsResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询轻量应用云主机列表。
 */
func (c *LavmClient) DescribeInstances(request *lavm.DescribeInstancesRequest) (*lavm.DescribeInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取lavm 实例，vnc url
 */
func (c *LavmClient) DescribeInstanceVncUrl(request *lavm.DescribeInstanceVncUrlRequest) (*lavm.DescribeInstanceVncUrlResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeInstanceVncUrlResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询轻量应用云主机详情。
 */
func (c *LavmClient) DescribeInstance(request *lavm.DescribeInstanceRequest) (*lavm.DescribeInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除自定义镜像。
 */
func (c *LavmClient) DeleteCustomImage(request *lavm.DeleteCustomImageRequest) (*lavm.DeleteCustomImageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DeleteCustomImageResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 
创建密钥。

公钥和私钥都由京东云生成，公钥保存在京东云，私钥返回给用户，由用户保存。

详细操作说明请参考帮助文档：[创建密钥](https://docs.jdcloud.com/cn/virtual-machines/create-keypair)

## 接口说明
- 调用该接口创建密钥后，公钥部分存储在京东云，并返回未加密的 `PEM` 编码的 `PKCS#8` 格式私钥，您只有一次机会保存您的私钥。请妥善保管。
 */
func (c *LavmClient) CreateKeypair(request *lavm.CreateKeypairRequest) (*lavm.CreateKeypairResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.CreateKeypairResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询指定地域下轻量应用云主机所有的套餐信息。
 */
func (c *LavmClient) DescribePlans(request *lavm.DescribePlansRequest) (*lavm.DescribePlansResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribePlansResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建自定义镜像。
 */
func (c *LavmClient) CreateCustomImage(request *lavm.CreateCustomImageRequest) (*lavm.CreateCustomImageResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.CreateCustomImageResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询资源配额。 */
func (c *LavmClient) QueryQuota(request *lavm.QueryQuotaRequest) (*lavm.QueryQuotaResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.QueryQuotaResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询指定轻量应用云主机的防火墙规则。
 */
func (c *LavmClient) DescribeFirewallRules(request *lavm.DescribeFirewallRulesRequest) (*lavm.DescribeFirewallRulesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeFirewallRulesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 删除指定轻量应用云主机的一条防火墙规则。
 */
func (c *LavmClient) DeleteFirewallRule(request *lavm.DeleteFirewallRuleRequest) (*lavm.DeleteFirewallRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DeleteFirewallRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建一台或多台轻量应用云主机。
 */
func (c *LavmClient) CreateInstances(request *lavm.CreateInstancesRequest) (*lavm.CreateInstancesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.CreateInstancesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 
删除密钥。

详细操作说明请参考帮助文档：[删除密钥](https://docs.jdcloud.com/cn/virtual-machines/delete-keypair)

## 接口说明
- 密钥删除后，使用该密钥的实例仍可正常使用与之匹配的本地私钥登录，且密钥仍会显示在实例详情中。
- 密钥删除后，与之关联的实例模板将变为不可用，并且与该实例模板关联的高可用组也会变为不可用。
 */
func (c *LavmClient) DeleteKeypair(request *lavm.DeleteKeypairRequest) (*lavm.DeleteKeypairResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DeleteKeypairResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 
为云主机实例解绑密钥。

详细操作说明请参考帮助文档：[绑定密钥](https://docs.jdcloud.com/cn/virtual-machines/bind-keypair)

## 接口说明
- 调用该接口解绑云主机实例中的密钥。
 */
func (c *LavmClient) DetachKeypair(request *lavm.DetachKeypairRequest) (*lavm.DetachKeypairResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DetachKeypairResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 
为云主机实例绑定密钥。

详细操作说明请参考帮助文档：[绑定密钥](https://docs.jdcloud.com/cn/virtual-machines/bind-keypair)

## 接口说明
- 只支持为 linux 云主机实例绑定密钥。
- 每台云主机实例只支持绑定一个密钥。如果云主机绑定的密钥被删除了，那么该云主机还可以再次绑定密钥。
 */
func (c *LavmClient) AttachKeypair(request *lavm.AttachKeypairRequest) (*lavm.AttachKeypairResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.AttachKeypairResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 为指定的轻量应用云主机修改一条防火墙规则。
 */
func (c *LavmClient) ModifyFirewallRule(request *lavm.ModifyFirewallRuleRequest) (*lavm.ModifyFirewallRuleResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.ModifyFirewallRuleResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 重启轻量应用云主机。
 */
func (c *LavmClient) RebootInstance(request *lavm.RebootInstanceRequest) (*lavm.RebootInstanceResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.RebootInstanceResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询指定地域下轻量应用云主机套餐信息是否售罄
 */
func (c *LavmClient) DescribePlansSoldOutStatus(request *lavm.DescribePlansSoldOutStatusRequest) (*lavm.DescribePlansSoldOutStatusResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribePlansSoldOutStatusResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* DescribeImages 查询指定地域下一个或多个镜像信息Image 模型。
 */
func (c *LavmClient) DescribeImages(request *lavm.DescribeImagesRequest) (*lavm.DescribeImagesResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeImagesResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询磁盘信息。
 */
func (c *LavmClient) DescribeDisks(request *lavm.DescribeDisksRequest) (*lavm.DescribeDisksResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribeDisksResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询当前实例可升级套餐列表。
 */
func (c *LavmClient) DescribePlansCanUpgrade(request *lavm.DescribePlansCanUpgradeRequest) (*lavm.DescribePlansCanUpgradeResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &lavm.DescribePlansCanUpgradeResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

