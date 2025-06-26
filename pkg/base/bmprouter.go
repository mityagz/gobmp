package base

var BmpRtrM map[string]BmpRouter

type L3Pkt struct {
	SrcIpPort string // from l3 packet
	SrcIp     string // from l3 packet
	SrcPort   string // from l3 packet
}

type BmpRouter struct {
	SrcIpPort string // from l3 packet
	SrcIp     string // from l3 packet
	SrcPort   string // from l3 packet
	SysFree   string // from initiation bmp message type 0
	SysDescr  string // from initiation bmp message type 1
	SysName   string // from initiation bmp message type 2
	RouterID  string // from pg db get_rid_by_hostname
	NodeID    int    // from pg db get_nodeid_by_hostname
}

func (bmpr *BmpRouter) Set(br L3Pkt) {
	bmpr.SrcIp = br.SrcIp
	bmpr.SrcPort = br.SrcPort
}

//var BmpRtrM map[id string]BmpRouter
// BmpRtrM = make(map[id]BmpRouter)
// we have to give out id to func(id string, ...)
