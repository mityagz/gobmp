package bmppg

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang/glog"
	bmp_message "github.com/sbezverk/gobmp/pkg/message"
)

func (ups *UnicastPrefixS) insertUnicastPrefixV4_HistPG() {
	q0 := `insert into ucast_pfx(action, router_id, router_hash, router_ip, base_attr_hash, 
			    peer_hash, peer_ip, peer_type, peer_asn, prefix, prefix_len,
			    is_ipv4, nexthop, is_nexthop_ipv4, labels, 
			    path_id, origin_as, is_adj_rib_in_post_policy, is_adj_rib_out_post_policy,
			    is_loc_rib_filtered, stimestamp, rtimestamp, ltimestamp) 
			    values(` + `'` + ups.Action + `', '` + ups.RouterID + `', '` + ups.RouterHash + `', 
			           '` + ups.RouterIP + `', '` + ups.BaseAttributes + `', '` + ups.PeerHash + `', 
				   '` + ups.PeerIP + `', ` + strconv.FormatInt(int64(ups.PeerType), 10) + `, 
				   ` + strconv.FormatInt(int64(ups.PeerASN), 10) + `, '` + ups.Prefix + `', 
				   ` + strconv.FormatInt(int64(ups.PrefixLen), 10) + `, 
				   ` + strconv.FormatBool(ups.IsIPv4) + `, '` + ups.Nexthop + `', 
				   ` + strconv.FormatBool(ups.IsNexthopIPv4) + `, '` + ups.Labels + `',  
				   ` + strconv.FormatInt(int64(ups.PathID), 10) + `, 
				   ` + strconv.FormatInt(int64(ups.OriginAS), 10) + `, 
				   ` + strconv.FormatBool(ups.IsAdjRIBInPost) + `, 
				   ` + strconv.FormatBool(ups.IsAdjRIBOutPost) + `, 
				   ` + strconv.FormatBool(ups.IsLocRIBFiltered) + `, '` + ups.Timestamp + `', 
				   '` + str2TS(ups.Timestamp) + `',  now())`
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

/*
func (ba *BaseAttributesS) searchBaseAttrByHash_HistPG(h string) bool {
	var (
		base_attr_hash string
	)
	q0 := `select base_attr_hash
			       from base_attrs where base_attr_hash like '` + h + `'`
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		switch err := rows.Scan(&base_attr_hash); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
			return false
		case nil:
			return true
		default:
		}
	}
	return false
}

func (ba *BaseAttributesS) searchBaseAttrByKey_HistPG(h string) bool {
	return false
}

func (ba *BaseAttributesS) insertBaseAttr_HistPG() {
	if ba.searchBaseAttrByHash_HistPG(ba.BaseAttrHash) {
		return
	}
	q0 := `insert into base_attrs(base_attr_hash, origin, as_path, as_path_count, nexthop,
				      med, local_pref, is_atomic_agg, aggregator, community_list,
				      originator_id, cluster_list, ext_community_list, as4_path,
				      as4_path_count, as4_aggregator, tunnel_encap_attr,
				      large_community_list, stimestamp, rtimestamp, ltimestamp)
				      values(` + `'` + ba.BaseAttrHash + `', '` + ba.Origin + `',
				      	    '` + ba.ASPath + `', ` + strconv.FormatInt(int64(ba.ASPathCount), 10) + `,
					    '` + ba.Nexthop + `', ` + strconv.FormatInt(int64(ba.MED), 10) + `,
					    ` + strconv.FormatInt(int64(ba.LocalPref), 10) + `,
					    ` + strconv.FormatBool(ba.IsAtomicAgg) + `,
					    '` + `', '` + ba.CommunityList + `',
					    '` + ba.OriginatorID + `', '` + ba.ClusterList + `',
					    '` + ba.ExtCommunityList + `', '` + ba.AS4Path + `',
					    ` + strconv.FormatInt(int64(ba.AS4PathCount), 10) + `,
					    '` + ba.AS4Aggregator + `',
					    '` + ba.TunnelEncapAttr + `', '` + ba.LgCommunityList + `',
					    '` + ba.Timestamp + `', '` + str2TS(ba.Timestamp) + `',  now())`
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}
*/

func (ups *UnicastPrefixS) searchPeerByHash_HistPG(h string) bool {
	var (
		peer_hash string
	)
	q0 := `select peer_hash
			       from peers where peer_hash like '` + h + `'`
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		switch err := rows.Scan(&peer_hash); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
			return false
		case nil:
			return true
		default:
		}
	}
	return false
}

func (ups *UnicastPrefixS) searchPeerByKey_HistPG(h string) bool {
	return false
}

func (ups *UnicastPrefixS) insertPeer_HistPG() {
	if ups.searchPeerByHash_HistPG(ups.PeerHash) {
		return
	}
	q0 := `insert into peers (peer_hash, peer_ip, peer_type, peer_asn, stimestamp, rtimestamp, ltimestamp)
				values('` + ups.PeerHash + `', '` + ups.PeerIP + `', 
				       ` + strconv.FormatInt(int64(ups.PeerType), 10) + `, 
				       ` + strconv.FormatInt(int64(ups.PeerASN), 10) + `, 
				       '` + ups.Timestamp + `', '` + str2TS(ups.Timestamp) + `',  now())`
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (bas *BaseAttributesS) baseAttrUcastPfx2S(pf bmp_message.UnicastPrefix) {
	ba := pf.BaseAttributes
	bas.BaseAttrHash = ba.BaseAttrHash                   //     string `json:"base_attr_hash,omitempty"`
	bas.Origin = ba.Origin                               //          string `json:"origin,omitempty"`
	bas.ASPath = arrInt2S(ba.ASPath)                     //           string // []uint32 `json:"as_path,omitempty"`
	bas.ASPathCount = ba.ASPathCount                     //     int32  `json:"as_path_count,omitempty"`
	bas.Nexthop = ba.Nexthop                             //         string `json:"nexthop,omitempty"`
	bas.MED = ba.MED                                     //              uint32 `json:"med,omitempty"`
	bas.LocalPref = ba.LocalPref                         //        uint32 `json:"local_pref,omitempty"`
	bas.IsAtomicAgg = ba.IsAtomicAgg                     //      bool   `json:"is_atomic_agg"`
	bas.Aggregator = arrByte2S(ba.Aggregator)            //      string // []byte   `json:"aggregator,omitempty"`
	bas.CommunityList = arrStr2S(ba.CommunityList)       //   string // []string `json:"community_list,omitempty"`
	bas.OriginatorID = ba.OriginatorID                   //     string `json:"originator_id,omitempty"`
	bas.ClusterList = ba.ClusterList                     //     string `json:"cluster_list,omitempty"`
	bas.ExtCommunityList = arrStr2S(ba.ExtCommunityList) // string // []string `json:"ext_community_list,omitempty"`
	bas.AS4Path = arrInt2S(ba.AS4Path)                   //         string // []uint32 `json:"as4_path,omitempty"`
	bas.AS4PathCount = ba.AS4PathCount                   //    int32  `json:"as4_path_count,omitempty"`
	bas.AS4Aggregator = arrByte2S(ba.AS4Aggregator)      //   string // []byte   `json:"as4_aggregator,omitempty"`
	// PMSITunnel
	bas.TunnelEncapAttr = arrByte2S(ba.TunnelEncapAttr) // string // []byte `json:"-"`
	// TraficEng
	// IPv6SpecExtCommunity
	// AIGP
	// PEDistinguisherLable
	bas.LgCommunityList = arrStr2S(ba.LgCommunityList) // string // []string `json:"large_community_list,omitempty"`
	// SecPath
	// AttrSet
	bas.Timestamp = pf.Timestamp //     string              `json:"timestamp,omitempty"
}

func (ups *UnicastPrefixS) ucastPfx2S(upp bmp_message.UnicastPrefix) {
	ups.Key = upp.Key                                    //           string              `json:"_key,omitempty"`
	ups.ID = upp.ID                                      //            string              `json:"_id,omitempty"`
	ups.Rev = upp.Rev                                    //           string              `json:"_rev,omitempty"`
	ups.Action = upp.Action                              //        string              `json:"action,omitempty"` // Action can be "add" or "del"
	ups.RouterID = upp.RouterID                          //      string              `json:"router_id,omitempty"`
	ups.Sequence = upp.Sequence                          //      int                 `json:"sequence,omitempty"`
	ups.Hash = upp.Hash                                  //          string              `json:"hash,omitempty"`
	ups.RouterHash = upp.RouterHash                      //    string              `json:"router_hash,omitempty"`
	ups.RouterIP = upp.RouterIP                          //      string              `json:"router_ip,omitempty"`
	ups.BaseAttributes = upp.BaseAttributes.BaseAttrHash // *bgp.BaseAttributes `json:"base_attrs,omitempty"`
	ups.PeerHash = upp.PeerHash                          //      string              `json:"peer_hash,omitempty"`
	ups.PeerIP = upp.PeerIP                              //        string              `json:"peer_ip,omitempty"`
	ups.PeerType = upp.PeerType                          //      uint8               `json:"peer_type"`
	ups.PeerASN = upp.PeerASN                            //       uint32              `json:"peer_asn,omitempty"`
	ups.Timestamp = upp.Timestamp                        //     string              `json:"timestamp,omitempty"`
	ups.Prefix = upp.Prefix                              //        string              `json:"prefix,omitempty"`
	ups.PrefixLen = upp.PrefixLen                        //     int32               `json:"prefix_len,omitempty"`
	ups.IsIPv4 = upp.IsIPv4                              //        bool                `json:"is_ipv4"`
	ups.OriginAS = upp.OriginAS                          //      uint32              `json:"origin_as,omitempty"`
	ups.Nexthop = upp.Nexthop                            //       string              `json:"nexthop,omitempty"`
	//ups.ClusterList = upp.ClusterList                    //   string              `json:"cluster_list,omitempty"`
	ups.IsNexthopIPv4 = upp.IsNexthopIPv4 // bool                `json:"is_nexthop_ipv4"`
	ups.PathID = upp.PathID               //        int32               `json:"path_id,omitempty"`
	ups.Labels = arrInt2S(upp.Labels)     //        []uint32            `json:"labels,omitempty"`
	//ups.VPNRD = upp.VPNRD                                //         string              `json:"vpn_rd,omitempty"`
	//ups.VPNRDType = upp.VPNRDType                        //     uint16              `json:"vpn_rd_type"`
	//l3v.PrefixSID = l3vp.PrefixSID                 //     *prefixsid.PSid     `json:"prefix_sid,omitempty"`
	// Values are assigned based on PerPeerHeader flas
	ups.IsAdjRIBInPost = upp.IsAdjRIBInPost     //  bool `json:"is_adj_rib_in_post_policy"`
	ups.IsAdjRIBOutPost = upp.IsAdjRIBOutPost   // bool `json:"is_adj_rib_out_post_policy"`
	ups.IsLocRIBFiltered = upp.IsLocRIBFiltered // bool `json:"is_loc_rib_filtered"`

}

func updUnicastPrefixV4_HistPG(uPfx *bmp_message.UnicastPrefix, ups UnicastPrefixS, bas BaseAttributesS) {
	bas.baseAttrUcastPfx2S(*uPfx)
	ups.ucastPfx2S(*uPfx)

	if glog.V(6) {
		glog.Infof("updUcastPfx SQL -> %s\n", uPfx)
		glog.Infof("updUcastPfx bas SQL -> %s\n", bas)
	}
	bas.insertBaseAttr_HistPG()
	ups.insertPeer_HistPG()
	ups.insertUnicastPrefixV4_HistPG()
}

/*
 CurrPG
*/

func (ups *UnicastPrefixS) searchIdByUcastPfx_CurrPG() (id int) {
	q0 := `select l.id from ucast_pfx_curr l, base_attrs b 
	      where l.base_attr_hash = b.base_attr_hash 
	      and prefix like '` + ups.Prefix + `' 
	      and prefix_len = ` + strconv.FormatInt(int64(ups.PrefixLen), 10) + ` 
	      and router_id like '` + ups.RouterID + `' 
	      and peer_ip like '` + ups.PeerIP + `'`
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows.Next()
	switch err := rows.Scan(&id); err {
	case sql.ErrNoRows:
		return -1
	case nil:
	default:
	}
	if glog.V(6) {
		glog.Infof("SQL ID -> %s\n", id)
	}
	var id1 int
	rows.Next()
	switch err := rows.Scan(&id1); err {
	case sql.ErrNoRows:
		return id
	case nil:
		glog.Infof("Too many rows were returned")
		os.Exit(1)
		return -2
	default:
	}
	if glog.V(6) {
		glog.Infof("SQL ID1 -> %s\n", id1)
	}
	return id
}
func (ups *UnicastPrefixS) insertUnicastPrefixV4_CurrPG() {
	q0 := `insert into ucast_pfx_curr(action, router_id, router_hash, router_ip, base_attr_hash, 
			    peer_hash, peer_ip, peer_type, peer_asn, prefix, prefix_len,
			    is_ipv4, nexthop, is_nexthop_ipv4, labels, 
			    path_id, origin_as, is_adj_rib_in_post_policy, is_adj_rib_out_post_policy,
			    is_loc_rib_filtered, stimestamp, rtimestamp, ltimestamp) 
			    values(` + `'` + ups.Action + `', '` + ups.RouterID + `', '` + ups.RouterHash + `', 
			           '` + ups.RouterIP + `', '` + ups.BaseAttributes + `', '` + ups.PeerHash + `', 
				   '` + ups.PeerIP + `', ` + strconv.FormatInt(int64(ups.PeerType), 10) + `, 
				   ` + strconv.FormatInt(int64(ups.PeerASN), 10) + `, '` + ups.Prefix + `', 
				   ` + strconv.FormatInt(int64(ups.PrefixLen), 10) + `, 
				   ` + strconv.FormatBool(ups.IsIPv4) + `, '` + ups.Nexthop + `', 
				   ` + strconv.FormatBool(ups.IsNexthopIPv4) + `, '` + ups.Labels + `', 
				   ` + strconv.FormatInt(int64(ups.PathID), 10) + `, 
				   ` + strconv.FormatInt(int64(ups.OriginAS), 10) + `, 
				   ` + strconv.FormatBool(ups.IsAdjRIBInPost) + `, 
				   ` + strconv.FormatBool(ups.IsAdjRIBOutPost) + `, 
				   ` + strconv.FormatBool(ups.IsLocRIBFiltered) + `, '` + ups.Timestamp + `', 
				   '` + str2TS(ups.Timestamp) + `',  now())`

	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (ups *UnicastPrefixS) updateUnicastPrefixV4_CurrPG(id int) {
	/*
		q0 := `update l3vpnV4_curr set action, router_id, router_hash, router_ip, base_attr_hash,
				    peer_hash, peer_ip, peer_type, peer_asn, prefix, prefix_len,
				    is_ipv4, nexthop, is_nexthop_ipv4, labels, vpn_rd, vpn_rd_type,
				    path_id, origin_as, is_adj_rib_in_post_policy, is_adj_rib_out_post_policy,
				    is_loc_rib_filtered, stimestamp, rtimestamp, ltimestamp)
				    values(` + `'` + l3v.Action + `', '` + l3v.RouterID + `', '` + l3v.RouterHash + `',
	*/
	q0 := `update ucast_pfx_curr set  base_attr_hash = '` + ups.BaseAttributes + `', 
			    nexthop = '` + ups.Nexthop + `', 
			    labels = '` + ups.Labels + `', 
			    path_id = ` + strconv.FormatInt(int64(ups.PathID), 10) + `, 
			    origin_as = ` + strconv.FormatInt(int64(ups.OriginAS), 10) + `, 
			    stimestamp = '` + ups.Timestamp + `', 
			    rtimestamp = '` + str2TS(ups.Timestamp) + `',  
			    ltimestamp = now() where id = ` + strconv.FormatInt(int64(id), 10)

	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (ups *UnicastPrefixS) deleteUcastPfxById_CurrPG(id int) {
	q0 := `delete from ucast_pfx_curr
	       where id = ` + strconv.FormatInt(int64(id), 10)
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (ups *UnicastPrefixS) modifyUnicastPrefixV4_CurrPG() {
	id := ups.searchIdByUcastPfx_CurrPG()
	// no row was found out
	if id == 0 {
		ups.insertUnicastPrefixV4_CurrPG()
	} else {
		if ups.Action == "add" {
			ups.updateUnicastPrefixV4_CurrPG(id)
		} else if ups.Action == "del" {
			ups.deleteUcastPfxById_CurrPG(id)
		}
	}

}

func updUnicastPrefixV4_CurrPG(uPfx *bmp_message.UnicastPrefix, ups UnicastPrefixS, bas BaseAttributesS) {
	bas.baseAttrUcastPfx2S(*uPfx)
	ups.ucastPfx2S(*uPfx)

	if glog.V(6) {
		glog.Infof("updUcastPfx SQL Curr -> %s\n", uPfx)
		glog.Infof("updUcast bas SQL Curr -> %s\n", bas)
	}
	//bas.insertBaseAttr_HistPG()
	//ups.insertPeer_HistPG()
	ups.modifyUnicastPrefixV4_CurrPG()
}

/*
type L3VPNPrefixS struct {
	Key            string `json:"_key,omitempty"`
	ID             string `json:"_id,omitempty"`
	Rev            string `json:"_rev,omitempty"`
	Action         string `json:"action,omitempty"` // Action can be "add" or "del"
	RouterID       string `json:"router_id,omitempty"`
	Sequence       int    `json:"sequence,omitempty"`
	Hash           string `json:"hash,omitempty"`
	RouterHash     string `json:"router_hash,omitempty"`
	RouterIP       string `json:"router_ip,omitempty"`
	BaseAttributes string //*bgp.BaseAttributes `json:"base_attrs,omitempty"`
	PeerHash       string `json:"peer_hash,omitempty"`
	PeerIP         string `json:"peer_ip,omitempty"`
	PeerType       uint8  `json:"peer_type"`
	PeerASN        uint32 `json:"peer_asn,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
	Prefix         string `json:"prefix,omitempty"`
	PrefixLen      int32  `json:"prefix_len,omitempty"`
	IsIPv4         bool   `json:"is_ipv4"`
	OriginAS       uint32 `json:"origin_as,omitempty"`
	Nexthop        string `json:"nexthop,omitempty"`
	ClusterList    string `json:"cluster_list,omitempty"`
	IsNexthopIPv4  bool   `json:"is_nexthop_ipv4"`
	PathID         int32  `json:"path_id,omitempty"`
	Labels         string //[]uint32            `json:"labels,omitempty"`
	VPNRD          string `json:"vpn_rd,omitempty"`
	VPNRDType      uint16 `json:"vpn_rd_type"`
	PrefixSID      string //*prefixsid.PSid `json:"prefix_sid,omitempty"`
	// Values are assigned based on PerPeerHeader flas
	IsAdjRIBInPost   bool `json:"is_adj_rib_in_post_policy"`
	IsAdjRIBOutPost  bool `json:"is_adj_rib_out_post_policy"`
	IsLocRIBFiltered bool `json:"is_loc_rib_filtered"`
}
*/

type UnicastPrefixS struct {
	Key            string `json:"_key,omitempty"`
	ID             string `json:"_id,omitempty"`
	Rev            string `json:"_rev,omitempty"`
	Action         string `json:"action,omitempty"` // Action can be "add" or "del"
	RouterID       string `json:"router_id,omitempty"`
	Sequence       int    `json:"sequence,omitempty"`
	Hash           string `json:"hash,omitempty"`
	RouterHash     string `json:"router_hash,omitempty"`
	RouterIP       string `json:"router_ip,omitempty"`
	BaseAttributes string //bgp.BaseAttributes `json:"base_attrs,omitempty"`
	PeerHash       string `json:"peer_hash,omitempty"`
	PeerIP         string `json:"peer_ip,omitempty"`
	PeerType       uint8  `json:"peer_type"`
	PeerASN        uint32 `json:"peer_asn,omitempty"`
	Timestamp      string `json:"timestamp,omitempty"`
	Prefix         string `json:"prefix,omitempty"`
	PrefixLen      int32  `json:"prefix_len,omitempty"`
	IsIPv4         bool   `json:"is_ipv4"`
	OriginAS       uint32 `json:"origin_as,omitempty"`
	Nexthop        string `json:"nexthop,omitempty"`
	IsNexthopIPv4  bool   `json:"is_nexthop_ipv4"`
	PathID         int32  `json:"path_id,omitempty"`
	Labels         string //[]uint32            `json:"labels,omitempty"`
	PrefixSID      string //*prefixsid.PSid     `json:"prefix_sid,omitempty"`
	IsEOR          bool   `json:"is_eor,omitempty"`
	// Values are assigned based on PerPeerHeader flags
	IsAdjRIBInPost   bool `json:"is_adj_rib_in_post_policy"`
	IsAdjRIBOutPost  bool `json:"is_adj_rib_out_post_policy"`
	IsLocRIBFiltered bool `json:"is_loc_rib_filtered"`
}

/*
type BaseAttributesS struct {
	BaseAttrHash     string `json:"base_attr_hash,omitempty"`
	Origin           string `json:"origin,omitempty"`
	ASPath           string // []uint32 `json:"as_path,omitempty"`
	ASPathCount      int32  `json:"as_path_count,omitempty"`
	Nexthop          string `json:"nexthop,omitempty"`
	MED              uint32 `json:"med,omitempty"`
	LocalPref        uint32 `json:"local_pref,omitempty"`
	IsAtomicAgg      bool   `json:"is_atomic_agg"`
	Aggregator       string // []byte   `json:"aggregator,omitempty"`
	CommunityList    string // []string `json:"community_list,omitempty"`
	OriginatorID     string `json:"originator_id,omitempty"`
	ClusterList      string `json:"cluster_list,omitempty"`
	ExtCommunityList string // []string `json:"ext_community_list,omitempty"`
	AS4Path          string // []uint32 `json:"as4_path,omitempty"`
	AS4PathCount     int32  `json:"as4_path_count,omitempty"`
	AS4Aggregator    string // []byte   `json:"as4_aggregator,omitempty"`
	// PMSITunnel
	TunnelEncapAttr string // []byte `json:"-"`
	// TraficEng
	// IPv6SpecExtCommunity
	// AIGP
	// PEDistinguisherLable
	LgCommunityList string // []string `json:"large_community_list,omitempty"`
	// SecPath
	// AttrSet
	Timestamp string //            `json:"timestamp,omitempty"
}
*/
