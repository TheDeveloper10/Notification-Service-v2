package util

type StatusCode int8

const (
	StatusSuccess StatusCode = iota
	StatusNotFound
	StatusDuplicate
	StatusError
)
