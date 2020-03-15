package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Topic_1_1 struct {
}

func (Topic_1_1) Symbol() string {
	return "▲"
}

func (Topic_1_1) Count() int {
	return 1
}

func init() {
	common.Register(&Topic_1_1{})
}
