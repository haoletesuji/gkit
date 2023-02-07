package pkg

import (
	gkit "gkit"
)

type Closer struct {
}

func NewCloser() gkit.Closer {
	return &Closer{}
}

func (c *Closer) Close() error {
	return nil
}
