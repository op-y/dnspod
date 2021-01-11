package dnspod

import (
	"fmt"
	"strings"
)

type Param struct {
	KV map[string]string
	S  []string
}

func NewParam() *Param {
	return &Param{KV: make(map[string]string), S: make([]string, 0)}
}

func (p *Param) Add(key, value string) {
	p.KV[key] = value
}

func (p *Param) ToSlice() {
	for k, v := range p.KV {
		item := fmt.Sprintf("%s=%s", k, v)
		p.S = append(p.S, item)
	}
}

func (p *Param) ToString() string {
	return strings.Join(p.S, "&")
}
