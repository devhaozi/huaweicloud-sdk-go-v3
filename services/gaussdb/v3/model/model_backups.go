package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Backups struct {

	// 备份ID。
	Id *string `json:"id,omitempty"`

	// 备份名称。
	Name *string `json:"name,omitempty"`

	// 备份开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	BeginTime *string `json:"begin_time,omitempty"`

	// 备份结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	EndTime *string `json:"end_time,omitempty"`

	// 备份状态，取值： - BUILDING：备份中。 - COMPLETED：备份完成。 - FAILED：备份失败。 - AVAILABLE：备份可用。
	Status *string `json:"status,omitempty"`

	// 备份花费时间（单位：minutes）
	TakeUpTime *int32 `json:"take_up_time,omitempty"`

	// 备份类型，取值：  - auto：自动全量备份。 - manual：手动全量备份。
	Type *string `json:"type,omitempty"`

	// 备份大小（单位：MB）。
	Size *int64 `json:"size,omitempty"`

	Datastore *MysqlDatastoreInBackup `json:"datastore,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 备份级别。当开启一级备份开关时，返回该参数。
	BackupLevel *BackupsBackupLevel `json:"backup_level,omitempty"`

	// 备份文件描述信息。
	Description *string `json:"description,omitempty"`
}

func (o Backups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Backups struct{}"
	}

	return strings.Join([]string{"Backups", string(data)}, " ")
}

type BackupsBackupLevel struct {
	value string
}

type BackupsBackupLevelEnum struct {
	E_0 BackupsBackupLevel
	E_1 BackupsBackupLevel
	E_2 BackupsBackupLevel
}

func GetBackupsBackupLevelEnum() BackupsBackupLevelEnum {
	return BackupsBackupLevelEnum{
		E_0: BackupsBackupLevel{
			value: "0",
		},
		E_1: BackupsBackupLevel{
			value: "1",
		},
		E_2: BackupsBackupLevel{
			value: "2",
		},
	}
}

func (c BackupsBackupLevel) Value() string {
	return c.value
}

func (c BackupsBackupLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupsBackupLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
