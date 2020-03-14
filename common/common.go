package common

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Topic interface {
	Symbol() string
	Count() int
}

var (
	topics = []Topic{}
)

func Register(t Topic) {
	topics = append(topics, t)
}

type sortByType []Topic

func (v sortByType) Len() int {
	return len(v)
}

func (v sortByType) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v sortByType) TypeName(i int) string {
	items := strings.Split(reflect.TypeOf(v[i]).String(), ".")
	if len(items) == 0 {
		return ""
	}
	return items[len(items)-1]
}

func (v sortByType) Less(i, j int) bool {
	return v.TypeName(i) < v.TypeName(j)
}

func Show() {
	if len(topics) == 0 {
		fmt.Println("no topic")
		return
	}
	fmt.Println("Topics:")
	sort.Sort(sortByType(topics))
	for _, t := range topics {
		fmt.Printf("  %s\n", strings.Repeat(t.Symbol()+" ", t.Count()))
	}
}
