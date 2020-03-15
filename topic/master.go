package topic

import (
	"github.com/jiangxin/git-merge-demo/common"
)

type Master struct{}

func (v Master) Symbol() string {
	return "◉"
}

func (v Master) Count() int {
	return 4
}

func init() {
	common.Register(&Master{})
}
