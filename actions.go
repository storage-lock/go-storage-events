package storage_events

// 用于统计Storage上的方法调用行为
const (
	ActionStorageGetName = "Storage.GetName"

	ActionStorageInit        = "Storage.Init"
	ActionStorageInitSuccess = "Storage.Init.Success"
	ActionStorageInitError   = "Storage.Init.Error"

	ActionStorageUpdateWithVersion        = "Storage.UpdateWithVersion"
	ActionStorageUpdateWithVersionSuccess = "Storage.UpdateWithVersion.Success"
	ActionStorageUpdateWithVersionError   = "Storage.UpdateWithVersion.Error"
	ActionStorageUpdateWithVersionMiss   = "Storage.UpdateWithVersion.Miss"

	ActionStorageInsertWithVersion        = "Storage.CreateWithVersion"
	ActionStorageInsertWithVersionSuccess = "Storage.CreateWithVersion.Success"
	ActionStorageInsertWithVersionError   = "Storage.CreateWithVersion.Error"
	ActionStorageInsertWithVersionMiss   = "Storage.CreateWithVersion.Miss"

	ActionStorageDeleteWithVersion        = "Storage.DeleteWithVersion"
	ActionStorageDeleteWithVersionSuccess = "Storage.DeleteWithVersion.Success"
	ActionStorageDeleteWithVersionError   = "Storage.DeleteWithVersion.Error"
	ActionStorageDeleteWithVersionMiss   = "Storage.DeleteWithVersion.Miss"

	ActionStorageGetTime        = "Storage.GetTime"
	ActionStorageGetTimeSuccess = "Storage.GetTime.Success"
	ActionStorageGetTimeError   = "Storage.GetTime.Error"

	ActionStorageGet        = "Storage.Get"
	ActionStorageGetSuccess = "Storage.Get.Success"
	ActionStorageGetError   = "Storage.Get.Error"

	ActionStorageClose        = "Storage.Close"
	ActionStorageCloseSuccess = "Storage.Close.Success"
	ActionStorageCloseError   = "Storage.Close.Error"

	ActionStorageList        = "Storage.List"
	ActionStorageListSuccess = "Storage.List.Success"
	ActionStorageListError   = "Storage.List.Error"
)

// action上的payload的名称，用于方便跟处理的时候统一
const (
	PayloadName                      = "name"
	PayloadIterator                  = "iterator"
	PayloadTime                      = "time"
	PayloadLockInformationJsonString = "lockInformationJsonString"
	PayloadLockId                    = "lockId"
	PayloadLockInformation           = "lockInformation"
	PayloadVersion                   = "version"
	PayloadNewVersion                = "newVersion"
	PayloadExceptedVersion           = "exceptedVersion"
)
