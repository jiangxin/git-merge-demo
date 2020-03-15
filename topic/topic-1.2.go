package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Topic_1_2 struct {
}

func (Topic_1_2) Symbol() string {
	return "△"
}

func (Topic_1_2) Count() int {
	return 3
}

func init() {
	common.Register(&Topic_1_2{})
}
