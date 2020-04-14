package srv6

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/tools"
)

// SIDStructure defines SRv6 SID Structure TLV object
// No RFC yet
type SIDStructure struct {
	LBLength  uint8
	LNLength  uint8
	FunLength uint8
	ArgLength uint8
}

func (st *SIDStructure) String(level ...int) string {
	var s string
	l := 0
	if level != nil {
		l = level[0]
	}
	s += tools.AddLevel(l)
	s += "SRv6 SID Structure TLV:" + "\n"

	s += tools.AddLevel(l + 1)
	s += fmt.Sprintf("Locator Block length: %d\n", st.LBLength)
	s += tools.AddLevel(l + 1)
	s += fmt.Sprintf("Locator Node length: %d\n", st.LNLength)
	s += tools.AddLevel(l + 1)
	s += fmt.Sprintf("SID Function length: %d\n", st.FunLength)
	s += tools.AddLevel(l + 1)
	s += fmt.Sprintf("SID Argument length: %d\n", st.ArgLength)

	return s
}

// UnmarshalSRv6SIDStructureTLV builds SRv6 SID Structure TLV object
func UnmarshalSRv6SIDStructureTLV(b []byte) (*SIDStructure, error) {
	glog.V(6).Infof("SRv6 SID Structure TLV Raw: %s", tools.MessageHex(b))
	st := SIDStructure{}
	p := 0
	st.LBLength = b[p]
	p++
	st.LNLength = b[p]
	p++
	st.FunLength = b[p]
	p++
	st.ArgLength = b[p]

	return &st, nil
}