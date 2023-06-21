package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePortInfoRequest struct {

	// 通道号ID。
	PortId string `json:"port_id"`
}

func (o DeletePortInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortInfoRequest struct{}"
	}

	return strings.Join([]string{"DeletePortInfoRequest", string(data)}, " ")
}
