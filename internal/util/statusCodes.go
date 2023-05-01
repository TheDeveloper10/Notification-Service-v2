package util

type StatusCode uint8

const (
	StatusSuccess StatusCode = iota
	StatusNotFound
	StatusDuplicate
	StatusInternal
	StatusError
	StatusTooMany
)
