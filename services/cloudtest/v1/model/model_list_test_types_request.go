package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestTypesRequest Request Object
type ListTestTypesRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`
}

func (o ListTestTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestTypesRequest struct{}"
	}

	return strings.Join([]string{"ListTestTypesRequest", string(data)}, " ")
}
