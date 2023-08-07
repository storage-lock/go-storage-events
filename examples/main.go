package main

import (
	"context"
	"fmt"
	"github.com/golang-infrastructure/go-iterator"
	"github.com/storage-lock/go-events"
	"github.com/storage-lock/go-storage"
	storage_events "github.com/storage-lock/go-storage-events"
	"time"
)

type FooStorage struct {
}

var _ storage.Storage = &FooStorage{}

func (x *FooStorage) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) Init(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) UpdateWithVersion(ctx context.Context, lockId string, exceptedVersion, newVersion storage.Version, lockInformation *storage.LockInformation) error {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) CreateWithVersion(ctx context.Context, lockId string, version storage.Version, lockInformation *storage.LockInformation) error {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) DeleteWithVersion(ctx context.Context, lockId string, exceptedVersion storage.Version, lockInformation *storage.LockInformation) error {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) Get(ctx context.Context, lockId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) GetTime(ctx context.Context) (time.Time, error) {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (x *FooStorage) List(ctx context.Context) (iterator.Iterator[*storage.LockInformation], error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	storage := &FooStorage{}
	executor := storage_events.NewWithEventSafeExecutor(storage)

	event := events.NewEvent("foo").AddListeners(events.NewListenerWrapper("", func(ctx context.Context, e *events.Event) {
		fmt.Print(e.ToJsonString())
	}))
	getTime, err := executor.GetTime(context.Background(), event)
	if err != nil {
		panic(err)
	}
	fmt.Println(getTime)
	// Output:
	// {"id":"storage-lock-event-eb8bffdcbf69414bb34719bfbda0147c","root_id":"storage-lock-event-eb8bffdcbf69414bb34719bfbda0147c","parent_id":"","lock_id":"foo","owner_id":"","storage_name":"","start_time":"2023-08-07T01:45:13.2932725+08:00","end_time":"2023-08-07T01:45:13.2932725+08:00","event_type":0,"actions":[{"start_time":"2023-08-07T01:45:13.293272 5+08:00","end_time":"2023-08-07T01:45:13.2932725+08:00","name":"Storage.GetTime","err":{},"payload_map":{"time":"0001-01-01T00:00:00Z"}}],"watch_dog_id":"","lock_information":null,"err":null}
	//
	// panic: implement me goroutine 1 [running]:
	// main.main()
	//     D:/workspace/go-storage-events/examples/main.go:72 +0x1c8
}
