package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAlarmRequest struct {

	// 告警规则的ID。
	AlarmId string `json:"alarm_id"`
}

func (o ShowAlarmRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmRequest struct{}"
	}

	return strings.Join([]string{"ShowAlarmRequest", string(data)}, " ")
}
