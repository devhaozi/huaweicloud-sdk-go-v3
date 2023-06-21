package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PushPortalInfoRequest struct {

	// 主页ID。
	PortalId string `json:"portal_id"`
}

func (o PushPortalInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushPortalInfoRequest struct{}"
	}

	return strings.Join([]string{"PushPortalInfoRequest", string(data)}, " ")
}
