package storage_events

// 用于统计Storage上的方法调用行为
const (
	ActionStorageGetName           = "Storage.GetName"
	ActionStorageInit              = "Storage.Init"
	ActionStorageUpdateWithVersion = "Storage.UpdateWithVersion"
	ActionStorageInsertWithVersion = "Storage.CreateWithVersion"
	ActionStorageDeleteWithVersion = "Storage.DeleteWithVersion"
	ActionStorageGetTime           = "Storage.GetTime"
	ActionStorageGet               = "Storage.Get"
	ActionStorageClose             = "Storage.Close"
	ActionStorageList              = "Storage.List"
)
