package storage_events

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-events"
	"github.com/storage-lock/go-storage"
)

type scriptStorage struct {
	err    error
	panic  any
	called string
}

func (x *scriptStorage) touch(name string) {
	x.called = name
	if x.panic != nil {
		panic(x.panic)
	}
}

func (x *scriptStorage) GetName() string {
	x.touch("name")
	return "script"
}
func (x *scriptStorage) Capabilities() []storage.StorageCapability { return nil }
func (x *scriptStorage) Init(context.Context) error {
	x.touch("init")
	return x.err
}
func (x *scriptStorage) UpdateWithVersion(context.Context, string, storage.Version, storage.Version, *storage.LockInformation) error {
	x.touch("update")
	return x.err
}
func (x *scriptStorage) CreateWithVersion(context.Context, string, storage.Version, *storage.LockInformation) error {
	x.touch("create")
	return x.err
}
func (x *scriptStorage) DeleteWithVersion(context.Context, string, storage.Version, *storage.LockInformation) error {
	x.touch("delete")
	return x.err
}
func (x *scriptStorage) Get(context.Context, string) (string, error) {
	x.touch("get")
	return "value", x.err
}
func (x *scriptStorage) GetTime(context.Context) (time.Time, error) {
	x.touch("time")
	return time.Unix(1, 0), x.err
}
func (x *scriptStorage) Close(context.Context) error {
	x.touch("close")
	return x.err
}
func (x *scriptStorage) List(context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	x.touch("list")
	return nil, x.err
}

func TestExecutorSuccessAndErrors(t *testing.T) {
	ctx := context.Background()
	store := &scriptStorage{err: errors.New("storage failed")}
	executor := NewWithEventSafeExecutor(store)
	event := events.NewEvent("lock")
	info := &storage.LockInformation{LockId: "lock", Version: 1}

	if name, err := executor.GetName(event); name != "script" || err != nil {
		t.Fatal(name, err)
	}
	if err := executor.Init(ctx, event); !errors.Is(err, store.err) {
		t.Fatal(err)
	}
	if err := executor.UpdateWithVersion(ctx, event, "lock", 1, 2, info); !errors.Is(err, store.err) {
		t.Fatal(err)
	}
	if err := executor.CreateWithVersion(ctx, event, "lock", 1, info); !errors.Is(err, store.err) {
		t.Fatal(err)
	}
	if err := executor.DeleteWithVersion(ctx, event, "lock", 1, info); !errors.Is(err, store.err) {
		t.Fatal(err)
	}
	if value, err := executor.Get(ctx, event, "lock"); value != "value" || !errors.Is(err, store.err) {
		t.Fatal(value, err)
	}
	if got, err := executor.GetTime(ctx, event); got.IsZero() || !errors.Is(err, store.err) {
		t.Fatal(got, err)
	}
	if err := executor.Close(ctx, event); !errors.Is(err, store.err) {
		t.Fatal(err)
	}
	if _, err := executor.List(ctx, event); !errors.Is(err, store.err) {
		t.Fatal(err)
	}
	if len(event.Actions) == 0 {
		t.Fatal("events were not recorded")
	}
}

func TestExecutorRecoversPanics(t *testing.T) {
	ctx := context.Background()
	panicErr := errors.New("exploded")
	store := &scriptStorage{panic: panicErr}
	executor := NewWithEventSafeExecutor(store)
	event := events.NewEvent("lock")
	info := &storage.LockInformation{LockId: "lock"}
	calls := []func() error{
		func() error { _, err := executor.GetName(event); return err },
		func() error { return executor.Init(ctx, event) },
		func() error { return executor.UpdateWithVersion(ctx, event, "lock", 1, 2, info) },
		func() error { return executor.CreateWithVersion(ctx, event, "lock", 1, info) },
		func() error { return executor.DeleteWithVersion(ctx, event, "lock", 1, info) },
		func() error { _, err := executor.Get(ctx, event, "lock"); return err },
		func() error { _, err := executor.GetTime(ctx, event); return err },
		func() error { return executor.Close(ctx, event) },
		func() error { _, err := executor.List(ctx, event); return err },
	}
	for _, call := range calls {
		if err := call(); err != nil && !errors.Is(err, panicErr) {
			t.Fatal(err)
		}
	}
	store.panic = "plain panic"
	if err := executor.Init(ctx, event); err == nil || err.Error() == "" {
		t.Fatal(err)
	}
	if err := executor.convertRecoveryToError(nil); err != nil {
		t.Fatal(err)
	}
}
