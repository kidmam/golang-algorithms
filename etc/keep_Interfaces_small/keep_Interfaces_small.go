package main

import (
	"io"
	"os"
)

// https://dev.to/wagslane/best-practices-for-writing-clean-interfaces-in-go-5c2j?utm_source=dormosheio&utm_campaign=dormosheio
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]os.FileInfo, error)
	Stat() (os.FileInfo, error)
}

type car interface {
	GetColor() string
	GetSpeed() int
	IsFiretruck() bool
}

type firetruck interface {
	car
	HoseLength() int
}
