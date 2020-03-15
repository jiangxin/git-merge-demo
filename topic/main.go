package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Main struct{}

func (v Main) Symbol() string {
	return "◉"
}

func (v Main) Count() int {
	return 3
}

func init() {
	common.Register(&Main{})
}
