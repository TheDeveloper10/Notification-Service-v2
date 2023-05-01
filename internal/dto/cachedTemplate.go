package dto

import "time"

type CachedTemplate struct {
	Template   *Template
	ExpiryTime int64
}

func (ct *CachedTemplate) IsExpired() bool {
	return time.Now().Unix() >= ct.ExpiryTime
}
