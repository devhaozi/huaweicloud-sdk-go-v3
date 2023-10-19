package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InterRegionBandwidth struct {

	// 资源ID标识符。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属帐号ID。
	DomainId string `json:"domain_id"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 资源ID标识符。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 资源ID标识符。
	BandwidthPackageId string `json:"bandwidth_package_id"`

	// 域间实例信息。
	InterRegions *[]InterRegion `json:"inter_regions,omitempty"`

	// 域间带宽的值。
	Bandwidth *int32 `json:"bandwidth,omitempty"`
}

func (o InterRegionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterRegionBandwidth struct{}"
	}

	return strings.Join([]string{"InterRegionBandwidth", string(data)}, " ")
}
