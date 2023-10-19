package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnterpriseRouterRegionId ER路由器的region ID。
type EnterpriseRouterRegionId struct {

	// RegionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`
}

func (o EnterpriseRouterRegionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseRouterRegionId struct{}"
	}

	return strings.Join([]string{"EnterpriseRouterRegionId", string(data)}, " ")
}
