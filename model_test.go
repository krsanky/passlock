package passlock

import (
	"testing"
	"time"
)

func TestCreateSave(t *testing.T) {
	pl := Create(1, "test title 123 ...", PasslockPassword(), time.Now())

	err := pl.Save()
	if err != nil {
		t.Errorf("err:%s\n", err.Error())
	}

	pl.Delete()
}
