package databases

import (
	"fmt"
	"sync"
)

type Driver interface {
	Open(conn, prefix string) error
	Up(source string) error
	Down(source string) error
}

var (
	driversMu sync.RWMutex
	drivers   map[string]Driver = make(map[string]Driver)
)

func GetDriver(scheme string) (Driver, error) {
	driversMu.RLock()
	d, ok := drivers[scheme]
	driversMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("unknown driver '%s'", scheme)
	}

	return d, nil
}

func Register(scheme string, driver Driver) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if driver == nil {
		panic("trying to register nil driver")
	}

	if _, ok := drivers[scheme]; ok {
		panic("trying to register the same driver " + scheme)
	}

	drivers[scheme] = driver
}
