package storage_events

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-events"
	"github.com/storage-lock/go-storage"
	"time"
)

// WithEventSafeExecutor 带事件观测和recover的Storage执行器
type WithEventSafeExecutor struct {
	storage storage.Storage
}

// NewWithEventSafeExecutor 从给定的Storage创建一个执行器，后面就拿着这个执行器操作而不直接操作Storage了
func NewWithEventSafeExecutor(storage storage.Storage) *WithEventSafeExecutor {
	return &WithEventSafeExecutor{
		storage: storage,
	}
}

func (x *WithEventSafeExecutor) GetName(e *events.Event) (name string) {

	ctx := context.Background()
	getNameAction := events.NewAction(ActionStorageGetName)

	defer func() {
		if r := recover(); r != nil {
			err := x.convertRecoveryToError(r)
			e.AddAction(getNameAction.End().SetErr(err).AddPayload(PayloadName, name)).Publish(ctx)
		}
	}()

	name = x.storage.GetName()
	e.AddAction(getNameAction.End().AddPayload(PayloadName, name)).Publish(ctx)

	return name
}

func (x *WithEventSafeExecutor) Init(ctx context.Context, e *events.Event) (err error) {

	initAction := events.NewAction(ActionStorageInit)

	defer func() {
		if r := recover(); r != nil {
			err := x.convertRecoveryToError(r)
			e.AddAction(initAction.End().SetErr(err)).Publish(ctx)
		}
	}()

	err = x.storage.Init(ctx)
	e.AddAction(initAction.End().SetErr(err)).Publish(ctx)

	return
}

func (x *WithEventSafeExecutor) UpdateWithVersion(ctx context.Context, e *events.Event, lockId string, exceptedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) (err error) {

	// 把相关的上下文能携带的都携带者
	updateAction := events.NewAction(ActionStorageUpdateWithVersion).
		AddPayload(PayloadLockId, lockId).
		AddPayload(PayloadExceptedVersion, exceptedVersion).
		AddPayload(PayloadNewVersion, newVersion).
		AddPayload(PayloadLockInformation, lockInformation)

	defer func() {
		if r := recover(); r != nil {
			err = x.convertRecoveryToError(r)
			e.AddAction(updateAction.End().SetErr(err)).Publish(ctx)
		}
	}()

	err = x.storage.UpdateWithVersion(ctx, lockId, exceptedVersion, newVersion, lockInformation)
	e.AddAction(updateAction.End().SetErr(err)).Publish(ctx)

	return
}

func (x *WithEventSafeExecutor) CreateWithVersion(ctx context.Context, e *events.Event, lockId string, version storage.Version, lockInformation *storage.LockInformation) (err error) {

	insertAction := events.NewAction(ActionStorageInsertWithVersion).
		AddPayload(PayloadLockId, lockId).
		AddPayload(PayloadVersion, version).
		AddPayload(PayloadLockInformation, lockInformation)

	defer func() {
		if r := recover(); r != nil {
			err = x.convertRecoveryToError(r)
			e.AddAction(insertAction.End().SetErr(err)).Publish(ctx)
		}
	}()

	err = x.storage.CreateWithVersion(ctx, lockId, version, lockInformation)
	e.AddAction(insertAction.End().SetErr(err)).Publish(ctx)

	return
}

func (x *WithEventSafeExecutor) DeleteWithVersion(ctx context.Context, e *events.Event, lockId string, exceptedVersion storage.Version, lockInformation *storage.LockInformation) (err error) {

	deleteAction := events.NewAction(ActionStorageDeleteWithVersion).
		AddPayload(PayloadLockId, lockId).
		AddPayload(PayloadExceptedVersion, exceptedVersion).
		AddPayload(PayloadLockInformation, lockInformation)

	defer func() {
		if r := recover(); r != nil {
			err = x.convertRecoveryToError(r)
			e.AddAction(deleteAction.End().SetErr(err)).Publish(ctx)
		}
	}()

	err = x.storage.DeleteWithVersion(ctx, lockId, exceptedVersion, lockInformation)
	e.AddAction(deleteAction.End().SetErr(err)).Publish(ctx)
	return err
}

func (x *WithEventSafeExecutor) Get(ctx context.Context, e *events.Event, lockId string) (lockInformationJsonString string, err error) {

	getAction := events.NewAction(ActionStorageGet).AddPayload(PayloadLockId, lockId)

	defer func() {
		if r := recover(); r != nil {
			err = x.convertRecoveryToError(r)
			// 虽然lockInformationJsonString大概率是空字符串，但是为了防止开发者二笔，这里还是将其也收集一下以免丢失数据
			e.AddAction(getAction.End().SetErr(err).AddPayload(PayloadLockInformationJsonString, lockInformationJsonString)).Publish(ctx)
		}
	}()

	lockInformationJsonString, err = x.storage.Get(ctx, lockId)
	e.AddAction(getAction.End().SetErr(err).AddPayload(PayloadLockInformationJsonString, lockInformationJsonString)).Publish(ctx)

	return
}

func (x *WithEventSafeExecutor) GetTime(ctx context.Context, e *events.Event) (time time.Time, err error) {

	getTimeAction := events.NewAction(ActionStorageGetTime)

	defer func() {
		if r := recover(); r != nil {
			err = x.convertRecoveryToError(r)
			e.AddAction(getTimeAction.End().SetErr(err).AddPayload(PayloadTime, time)).Publish(ctx)
		}
	}()

	time, err = x.storage.GetTime(ctx)
	e.AddAction(getTimeAction.End().SetErr(err).AddPayload(PayloadTime, time)).Publish(ctx)

	return
}

func (x *WithEventSafeExecutor) Close(ctx context.Context, e *events.Event) (err error) {

	closeAction := events.NewAction(ActionStorageClose)

	defer func() {
		if r := recover(); r != nil {
			err = x.convertRecoveryToError(r)
			e.AddAction(closeAction.End().SetErr(err)).Publish(ctx)
		}
	}()

	err = x.storage.Close(ctx)
	e.AddAction(closeAction.End().SetErr(err)).Publish(ctx)

	return
}

func (x *WithEventSafeExecutor) List(ctx context.Context, e *events.Event) (iterator iterator.Iterator[*storage.LockInformation], err error) {

	listAction := events.NewAction(ActionStorageList)

	defer func() {
		if r := recover(); r != nil {
			err = x.convertRecoveryToError(r)
			e.AddAction(listAction.End().SetErr(err).AddPayload(PayloadIterator, iterator)).Publish(ctx)
		}
	}()

	iterator, err = x.storage.List(ctx)
	e.AddAction(listAction.End().SetErr(err).AddPayload(PayloadIterator, iterator)).Publish(ctx)

	return
}

// 把捕捉到的异常信息转换为
func (x *WithEventSafeExecutor) convertRecoveryToError(r any) error {

	// 空值判断
	if r == nil {
		return nil
	}

	// 如果本身就是个错误的话就直接类型转换
	err, ok := r.(error)
	if ok {
		return err
	}

	// TODO 优化转换
	return errors.New(fmt.Sprintf("%#v", r))
}
