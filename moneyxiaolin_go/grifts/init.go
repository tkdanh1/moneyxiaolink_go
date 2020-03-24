package grifts

import (
  "github.com/gobuffalo/buffalo"
	"moneyxiaolin_go/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
