package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthPackageId 云连接实例ID。
type BandwidthPackageId struct {

	// 资源ID标识符。
	BandwidthPackageId string `json:"bandwidth_package_id"`
}

func (o BandwidthPackageId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPackageId struct{}"
	}

	return strings.Join([]string{"BandwidthPackageId", string(data)}, " ")
}
