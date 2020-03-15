package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Topic_3 struct{}

func (v Topic_3) Symbol() string {
	return "♦"
}

func (v Topic_3) Count() int {
	return 3
}

func init() {
	common.Register(&Topic_3{})
}
