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
    jmr "github.com/jdcloud-api/jdcloud-sdk-go/services/jmr/apis"
    "encoding/json"
    "errors"
)

type JmrClient struct {
    core.JDCloudClient
}

func NewJmrClient(credential *core.Credential) *JmrClient {
    if credential == nil {
        return nil
    }

    config := core.NewConfig()
    config.SetEndpoint("jmr.jdcloud-api.com")

    return &JmrClient{
        core.JDCloudClient{
            Credential:  *credential,
            Config:      *config,
            ServiceName: "jmr",
            Revision:    "1.1.3",
            Logger:      core.NewDefaultLogger(core.LogInfo),
        }}
}

func (c *JmrClient) SetConfig(config *core.Config) {
    c.Config = *config
}

func (c *JmrClient) SetLogger(logger core.Logger) {
    c.Logger = logger
}

func (c *JmrClient) DisableLogger() {
    c.Logger = core.NewDummyLogger()
}

/* 查询JMR的版本信息 */
func (c *JmrClient) GetJmrVersionList(request *jmr.GetJmrVersionListRequest) (*jmr.GetJmrVersionListResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.GetJmrVersionListResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 释放集群
 */
func (c *JmrClient) ReleaseCluster(request *jmr.ReleaseClusterRequest) (*jmr.ReleaseClusterResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.ReleaseClusterResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询集群列表 */
func (c *JmrClient) DescribeClusters(request *jmr.DescribeClustersRequest) (*jmr.DescribeClustersResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.DescribeClustersResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询用户指定clusterId对应的集群列表及相关服务的一些信息 */
func (c *JmrClient) IdataCluster(request *jmr.IdataClusterRequest) (*jmr.IdataClusterResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.IdataClusterResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 获取对应版本的软件清单信息 */
func (c *JmrClient) GetSoftwareInfo(request *jmr.GetSoftwareInfoRequest) (*jmr.GetSoftwareInfoResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.GetSoftwareInfoResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 扩容集群 */
func (c *JmrClient) ClusterExpansion(request *jmr.ClusterExpansionRequest) (*jmr.ClusterExpansionResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.ClusterExpansionResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 创建集群 */
func (c *JmrClient) CreateCluster(request *jmr.CreateClusterRequest) (*jmr.CreateClusterResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.CreateClusterResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

/* 查询指定集群的详细信息
 */
func (c *JmrClient) DescribeCluster(request *jmr.DescribeClusterRequest) (*jmr.DescribeClusterResponse, error) {
    if request == nil {
        return nil, errors.New("Request object is nil. ")
    }
    resp, err := c.Send(request, c.ServiceName)
    if err != nil {
        return nil, err
    }

    jdResp := &jmr.DescribeClusterResponse{}
    err = json.Unmarshal(resp, jdResp)
    if err != nil {
        c.Logger.Log(core.LogError, "Unmarshal json failed, resp: %s", string(resp))
        return nil, err
    }

    return jdResp, err
}

