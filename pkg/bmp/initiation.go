package bmp

import (
	"encoding/binary"
	"fmt"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/base"
	"github.com/sbezverk/tools"
)

// InitiationMessage defines BMP Initiation Message per rfc7854
type InitiationMessage struct {
	TLV []InformationalTLV
}

// UnmarshalInitiationMessage processes Initiation Message and returns BMPInitiationMessage object
func UnmarshalInitiationMessage(l3p base.L3Pkt, b []byte) (*InitiationMessage, error) {
	// m
	brt := base.BmpRouter{SrcIpPort: "", SrcIp: "", SrcPort: "", SysFree: "", SysDescr: "", SysName: "", RouterID: "", NodeID: 0}
	brt.SrcIpPort = l3p.SrcIpPort
	brt.SrcIp = l3p.SrcIp
	brt.SrcPort = l3p.SrcPort

	if glog.V(6) {
		glog.Infof("BMP Initiation Message Raw: %s", tools.MessageHex(b))
	}
	im := &InitiationMessage{
		TLV: make([]InformationalTLV, 0),
	}
	for i := 0; i < len(b); {
		// Extracting TLV type 2 bytes
		t := int16(binary.BigEndian.Uint16(b[i : i+2]))
		switch t {
		case 0:
		case 1:
		case 2:
		default:
			return nil, fmt.Errorf("invalid tlv type, expected between 0 and 2 found %d", t)
		}
		// Extracting TLV length
		l := int16(binary.BigEndian.Uint16(b[i+2 : i+4]))
		if l > int16(len(b)-(i+4)) {
			return nil, fmt.Errorf("invalid tlv length %d", l)
		}
		v := b[i+4 : i+4+int(l)]
		im.TLV = append(im.TLV, InformationalTLV{
			InformationType:   t,
			InformationLength: l,
			Information:       v,
		})
		i += 4 + int(l)
		// m
		switch t {
		case 0:
			brt.SysFree = string(v)
		case 1:
			brt.SysDescr = string(v)
		case 2:
			brt.SysName = string(v)
		}
	}
	brt.RouterID = brt.SrcIp + ":" + brt.SysName
	base.BmpRtrM[l3p.SrcIpPort] = brt

	return im, nil
}
