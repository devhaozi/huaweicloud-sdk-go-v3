package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RebootCloudTableClusterResponse struct {
	Body           *[]RestartInstanceRsp `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o RebootCloudTableClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootCloudTableClusterResponse struct{}"
	}

	return strings.Join([]string{"RebootCloudTableClusterResponse", string(data)}, " ")
}
