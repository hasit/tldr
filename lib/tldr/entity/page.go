package entity

import "io"

type Page interface {
	Reader() io.Reader
	Close() error
}
