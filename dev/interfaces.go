package dev

type Info interface {
	GetName() string
	GetAddress() Address
	GetStatus() Status
}

type Action interface {
	Reboot() error
	Reset() error
}

type Address interface {
	stringCompatible
}
type Status interface {
	stringCompatible
}

type stringCompatible interface {
	ToString() string
}