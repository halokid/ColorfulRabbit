package error

import "fmt"

type ColorError struct {
  Msg string
  File  string
  Line int
}

func (c *ColorError) Error() string {
  return fmt.Sprintf("%s: %d: %s", c.File, c.Line, c.Msg)
}


