package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Next struct{}

func (v Next) Symbol() string {
	return "❍"
}

func (v Next) Count() int {
	return 4
}

func init() {
	common.Register(&Next{})
}
