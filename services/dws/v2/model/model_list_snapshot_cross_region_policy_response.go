package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotCrossRegionPolicyResponse Response Object
type ListSnapshotCrossRegionPolicyResponse struct {

	// 跨区域配置信息
	CrossRegionConfigs *[]CrossRegionSnapshotConfig `json:"cross_region_configs,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSnapshotCrossRegionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotCrossRegionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotCrossRegionPolicyResponse", string(data)}, " ")
}
