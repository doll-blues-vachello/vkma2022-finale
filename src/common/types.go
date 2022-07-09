package types

type Id uint64

type AccessLevel uint8

const (
	AccessLevel READ   = 1
	AccessLevel CREATE = 2
	AccessLevel DELETE = 4
)
