package store

import (
	"github.com/TigerZhang/ledisdb/store/driver"
)

type Slice interface {
	driver.ISlice
}
