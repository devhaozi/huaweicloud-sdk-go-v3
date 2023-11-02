package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataobjectRelationsResponse Response Object
type ListDataobjectRelationsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 告警总数
	Total *int32 `json:"total,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 告警列表
	Data *[]DataObjectDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDataobjectRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataobjectRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListDataobjectRelationsResponse", string(data)}, " ")
}
