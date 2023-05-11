package hash

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hash struct {
	salt string
}

func NewMd5Hash(sait string) *md5Hash {
	return &md5Hash{salt: sait}
}

func (h *md5Hash) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (h *md5Hash) HashSliceByte(data []byte) string {
	hasher := md5.New()
	hasher.Write(data)
	hasher.Write([]byte(h.salt))
	return hex.EncodeToString(hasher.Sum(nil))
}
