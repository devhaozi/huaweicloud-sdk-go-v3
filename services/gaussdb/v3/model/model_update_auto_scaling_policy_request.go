package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoScalingPolicyRequest Request Object
type UpdateAutoScalingPolicyRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 内容类型。  取值：application/json。
	ContentType string `json:"Content-Type"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *UpdateAutoScalingPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutoScalingPolicyRequest", string(data)}, " ")
}
