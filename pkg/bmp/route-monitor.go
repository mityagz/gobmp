package bmp

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/sbezverk/gobmp/pkg/base"
	"github.com/sbezverk/gobmp/pkg/bgp"
	"github.com/sbezverk/tools"
)

// RouteMonitor defines a structure of BMP Route Monitoring message
type RouteMonitor struct {
	Update *bgp.Update
	L3p    base.L3Pkt
}

// UnmarshalBMPRouteMonitorMessage builds BMP Route Monitor object
// m
func UnmarshalBMPRouteMonitorMessage(l3p base.L3Pkt, b []byte) (*RouteMonitor, error) {
	if glog.V(6) {
		glog.Infof("BMP Route Monitor Message Raw: %s length: %d", tools.MessageHex(b), len(b))
	}
	rm := RouteMonitor{}
	// 16 bytes marker + 2 bytes update length + 1 byte of type
	if len(b) < 19 {
		return nil, fmt.Errorf("malformed route monitor message")
	}
	p := 0
	// Skip 16 bytes of a marker
	p += 16
	// Skip 2 bytes of the update length
	p += 2
	// Getting update type, currently only type 2 is processed
	t := b[p]
	p++
	switch t {
	case 2:
		// Update type
		u, err := bgp.UnmarshalBGPUpdate(b[p:])
		if err != nil {
			return nil, err
		}
		rm.Update = u
		// m
		rm.L3p = l3p
	default:
	}

	return &rm, nil
}
