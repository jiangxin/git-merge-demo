package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Topic_2 struct{}

func (v Topic_2) Symbol() string {
	return "♥"
}

func (v Topic_2) Count() int {
	return 1
}

func init() {
	common.Register(&Topic_2{})
}
