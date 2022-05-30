package crypto

import (
	"github.com/elancom/go-util/rand"
	"github.com/elancom/go-util/str"
	"github.com/google/uuid"
	"time"
)

func NewUUID() string {
	u := uuid.New()
	return u.String()
}

func NewId32() string {
	return md6(str.String(time.Now().UnixNano()) + NewUUID() + rand.RandomStr(32))
}

func NewId64() string {
	return NewId32() + NewId32()
}
