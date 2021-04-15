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


type NodeSpec struct {

    /* CPU核数 (Optional) */
    CpuCores int `json:"cpuCores"`

    /* 内存大小（GB） (Optional) */
    MemoryGBs int `json:"memoryGBs"`

    /* 总磁盘大小（GB） (Optional) */
    TotalDiskGBs int `json:"totalDiskGBs"`

    /* 本地磁盘大小（GB） (Optional) */
    LocalDiskGBs int `json:"localDiskGBs"`

    /* 云盘大小（GB） (Optional) */
    CloudDiskGBs int `json:"cloudDiskGBs"`
}