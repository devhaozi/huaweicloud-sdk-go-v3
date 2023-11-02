package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckpointRequest Request Object
type ListCheckpointRequest struct {

	// 数据作业id
	DataJobId string `json:"data_job_id"`

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`
}

func (o ListCheckpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckpointRequest struct{}"
	}

	return strings.Join([]string{"ListCheckpointRequest", string(data)}, " ")
}
