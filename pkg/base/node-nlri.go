package base

import (
	"encoding/binary"
	"fmt"

	"github.com/sbezverk/gobmp/pkg/internal"
)

// NodeNLRI defines Node NLRI onject
// https://tools.ietf.org/html/rfc7752#section-3.2
type NodeNLRI struct {
	ProtocolID uint8
	Identifier uint64
	LocalNode  *NodeDescriptor
}

func (n *NodeNLRI) String() string {
	var s string
	s += fmt.Sprintf("Protocol ID: %s\n", internal.ProtocolIDString(n.ProtocolID))
	s += fmt.Sprintf("Identifier: %d\n", n.Identifier)
	s += n.LocalNode.String()

	return s
}

// UnmarshalNodeNLRI builds Node NLRI object
func UnmarshalNodeNLRI(b []byte) (*NodeNLRI, error) {
	n := NodeNLRI{}
	p := 0
	n.ProtocolID = b[p]
	p++

	n.Identifier = binary.BigEndian.Uint64(b[p : p+8])
	p += 8
	// Local Node Descriptor
	// Get Node Descriptor's length, skip Node Descriptor Type
	ndl := binary.BigEndian.Uint16(b[p+2 : p+4])
	ln, err := UnmarshalNodeDescriptor(b[p : p+int(ndl)])
	if err != nil {
		return nil, err
	}
	n.LocalNode = ln

	return &n, nil
}