package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccountRequest Request Object
type ShowAccountRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 是否返回账号邮箱、手机号信息。若此参数为True，Limit最多200。
	WithRegisterContactInfo *bool `json:"with_register_contact_info,omitempty"`

	// 账号的唯一标识符（ID）。
	AccountId string `json:"account_id"`
}

func (o ShowAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowAccountRequest", string(data)}, " ")
}
