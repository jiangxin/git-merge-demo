package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Topic_1 struct{}

func (v Topic_1) Symbol() string {
	return "♠"
}

func (v Topic_1) Count() int {
	return 2
}

func init() {
	common.Register(&Topic_1{})
}
