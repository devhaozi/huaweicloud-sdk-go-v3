package model

import (
	"github.com/devhaozi/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTaskRequest Request Object
type CreateTaskRequest struct {
	Body *PostTask `json:"body,omitempty"`
}

func (o CreateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateTaskRequest", string(data)}, " ")
}
