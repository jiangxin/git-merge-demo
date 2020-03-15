package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Topic_4 struct{}

func (v Topic_4) Symbol() string {
	return "♦"
}

func (v Topic_4) Count() int {
	return 2
}

func init() {
	common.Register(&Topic_4{})
}
