package device

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// Interface defines the API for the plugin package
type Interface interface {
	Start(ctx context.Context) error
	Stop() error
}

func Add(d *Device) error {
	if d == nil {
		return fmt.Errorf("no device specified")
	}

	if d.handler == nil {
		return fmt.Errorf("no handler specified")
	}

	if d.ID == "" {
		d.ID = uuid.NewString()
	}

	if d.stopCh == nil {
		d.stopCh = make(chan struct{})
	}

	if d.workers == 0 {
		d.workers = 1
	}

	lock.Lock()
	defer lock.Unlock()

	if _, ok := devs[d.ID]; ok {
		return fmt.Errorf("already exists")
	}

	devs[d.ID] = d

	return nil
}

func Del(id string) error {
	lock.Lock()
	defer lock.Unlock()

	if _, ok := devs[id]; !ok {
		return fmt.Errorf("%s no exist", id)
	}

	delete(devs, id)

	return nil
}

func GetAllDevices() []Interface {
	lock.RLock()
	defer lock.RUnlock()

	results := []Interface{}
	for _, v := range devs {
		results = append(results, v)
	}

	return results
}
