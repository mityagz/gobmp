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

/*

func arrInt2S(ai []uint32) string {
	var r string
	for _, i := range ai {
		r = r + " " + strconv.FormatInt(int64(i), 10)
	}
	return r
}

func arrByte2S(ab []byte) string {
	r := string(ab[:])
	return r
}

func arrStr2S(as []string) string {
	var r string
	for _, s := range as {
		r = r + " " + s
	}
	return r
}

func str2TS(t string) string {
	var r string
	rr := strings.Split(t, "T")
	r1 := strings.TrimRight(rr[1], "Z")
	r = rr[0] + " " + r1
	return r
}

*/

func (l3v *L3VPNPrefixS) insertL3VPNV4_HistPG() {
	q0 := `insert into l3vpnV4(action, router_id, router_hash, router_ip, base_attr_hash, 
			    peer_hash, peer_ip, peer_type, peer_asn, prefix, prefix_len,
			    is_ipv4, nexthop, is_nexthop_ipv4, labels, vpn_rd, vpn_rd_type,
			    path_id, origin_as, is_adj_rib_in_post_policy, is_adj_rib_out_post_policy,
			    is_loc_rib_filtered, stimestamp, rtimestamp, ltimestamp) 
			    values(` + `'` + l3v.Action + `', '` + l3v.RouterID + `', '` + l3v.RouterHash + `', 
			           '` + l3v.RouterIP + `', '` + l3v.BaseAttributes + `', '` + l3v.PeerHash + `', 
				   '` + l3v.PeerIP + `', ` + strconv.FormatInt(int64(l3v.PeerType), 10) + `, 
				   ` + strconv.FormatInt(int64(l3v.PeerASN), 10) + `, '` + l3v.Prefix + `', 
				   ` + strconv.FormatInt(int64(l3v.PrefixLen), 10) + `, 
				   ` + strconv.FormatBool(l3v.IsIPv4) + `, '` + l3v.Nexthop + `', 
				   ` + strconv.FormatBool(l3v.IsNexthopIPv4) + `, '` + l3v.Labels + `', 
				   '` + l3v.VPNRD + `', ` + strconv.FormatInt(int64(l3v.VPNRDType), 10) + `, 
				   ` + strconv.FormatInt(int64(l3v.PathID), 10) + `, 
				   ` + strconv.FormatInt(int64(l3v.OriginAS), 10) + `, 
				   ` + strconv.FormatBool(l3v.IsAdjRIBInPost) + `, 
				   ` + strconv.FormatBool(l3v.IsAdjRIBOutPost) + `, 
				   ` + strconv.FormatBool(l3v.IsLocRIBFiltered) + `, '` + l3v.Timestamp + `', 
				   '` + str2TS(l3v.Timestamp) + `',  now())`

	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

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

func (l3v *L3VPNPrefixS) searchPeerByHash_HistPG(h string) bool {
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

func (l3v *L3VPNPrefixS) searchPeerByKey_HistPG(h string) bool {
	return false
}

func (l3v *L3VPNPrefixS) insertPeer_HistPG() {
	if l3v.searchPeerByHash_HistPG(l3v.PeerHash) {
		return
	}
	q0 := `insert into peers (peer_hash, peer_ip, peer_type, peer_asn, stimestamp, rtimestamp, ltimestamp)
				values('` + l3v.PeerHash + `', '` + l3v.PeerIP + `', 
				       ` + strconv.FormatInt(int64(l3v.PeerType), 10) + `, 
				       ` + strconv.FormatInt(int64(l3v.PeerASN), 10) + `, 
				       '` + l3v.Timestamp + `', '` + str2TS(l3v.Timestamp) + `',  now())`
	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (bas *BaseAttributesS) baseAttr2S(pf bmp_message.L3VPNPrefix) {
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

func (l3v *L3VPNPrefixS) l3vpnPfx2S(l3vp bmp_message.L3VPNPrefix) {
	l3v.Key = l3vp.Key                                    //           string              `json:"_key,omitempty"`
	l3v.ID = l3vp.ID                                      //            string              `json:"_id,omitempty"`
	l3v.Rev = l3vp.Rev                                    //           string              `json:"_rev,omitempty"`
	l3v.Action = l3vp.Action                              //        string              `json:"action,omitempty"` // Action can be "add" or "del"
	l3v.RouterID = l3vp.RouterID                          //      string              `json:"router_id,omitempty"`
	l3v.Sequence = l3vp.Sequence                          //      int                 `json:"sequence,omitempty"`
	l3v.Hash = l3vp.Hash                                  //          string              `json:"hash,omitempty"`
	l3v.RouterHash = l3vp.RouterHash                      //    string              `json:"router_hash,omitempty"`
	l3v.RouterIP = l3vp.RouterIP                          //      string              `json:"router_ip,omitempty"`
	l3v.BaseAttributes = l3vp.BaseAttributes.BaseAttrHash // *bgp.BaseAttributes `json:"base_attrs,omitempty"`
	l3v.PeerHash = l3vp.PeerHash                          //      string              `json:"peer_hash,omitempty"`
	l3v.PeerIP = l3vp.PeerIP                              //        string              `json:"peer_ip,omitempty"`
	l3v.PeerType = l3vp.PeerType                          //      uint8               `json:"peer_type"`
	l3v.PeerASN = l3vp.PeerASN                            //       uint32              `json:"peer_asn,omitempty"`
	l3v.Timestamp = l3vp.Timestamp                        //     string              `json:"timestamp,omitempty"`
	l3v.Prefix = l3vp.Prefix                              //        string              `json:"prefix,omitempty"`
	l3v.PrefixLen = l3vp.PrefixLen                        //     int32               `json:"prefix_len,omitempty"`
	l3v.IsIPv4 = l3vp.IsIPv4                              //        bool                `json:"is_ipv4"`
	l3v.OriginAS = l3vp.OriginAS                          //      uint32              `json:"origin_as,omitempty"`
	l3v.Nexthop = l3vp.Nexthop                            //       string              `json:"nexthop,omitempty"`
	l3v.ClusterList = l3vp.ClusterList                    //   string              `json:"cluster_list,omitempty"`
	l3v.IsNexthopIPv4 = l3vp.IsNexthopIPv4                // bool                `json:"is_nexthop_ipv4"`
	l3v.PathID = l3vp.PathID                              //        int32               `json:"path_id,omitempty"`
	l3v.Labels = arrInt2S(l3vp.Labels)                    //        []uint32            `json:"labels,omitempty"`
	l3v.VPNRD = l3vp.VPNRD                                //         string              `json:"vpn_rd,omitempty"`
	l3v.VPNRDType = l3vp.VPNRDType                        //     uint16              `json:"vpn_rd_type"`
	//l3v.PrefixSID = l3vp.PrefixSID                 //     *prefixsid.PSid     `json:"prefix_sid,omitempty"`
	// Values are assigned based on PerPeerHeader flas
	l3v.IsAdjRIBInPost = l3vp.IsAdjRIBInPost     //  bool `json:"is_adj_rib_in_post_policy"`
	l3v.IsAdjRIBOutPost = l3vp.IsAdjRIBOutPost   // bool `json:"is_adj_rib_out_post_policy"`
	l3v.IsLocRIBFiltered = l3vp.IsLocRIBFiltered // bool `json:"is_loc_rib_filtered"`

}

func updL3VPNV4_HistPG(l3vpnPfx *bmp_message.L3VPNPrefix, l3v L3VPNPrefixS, bas BaseAttributesS) {
	bas.baseAttr2S(*l3vpnPfx)
	l3v.l3vpnPfx2S(*l3vpnPfx)

	if glog.V(6) {
		glog.Infof("updL3VPNV4 SQL -> %s\n", l3vpnPfx)
		glog.Infof("updL3VPNV4 bas SQL -> %s\n", bas)
	}
	bas.insertBaseAttr_HistPG()
	l3v.insertPeer_HistPG()
	l3v.insertL3VPNV4_HistPG()
}

/*
 CurrPG
*/

func (l3v *L3VPNPrefixS) searchIdByPfx_CurrPG() (id int) {
	q0 := `select l.id from l3vpnv4_curr l, base_attrs b 
	      where l.base_attr_hash = b.base_attr_hash 
	      and prefix like '` + l3v.Prefix + `' 
	      and prefix_len = ` + strconv.FormatInt(int64(l3v.PrefixLen), 10) + ` 
	      and vpn_rd like '` + l3v.VPNRD + `' 
	      and router_id like '` + l3v.RouterID + `' 
	      and peer_ip like '` + l3v.PeerIP + `'`
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
func (l3v *L3VPNPrefixS) insertL3VPNV4_CurrPG() {
	q0 := `insert into l3vpnV4_curr(action, router_id, router_hash, router_ip, base_attr_hash, 
			    peer_hash, peer_ip, peer_type, peer_asn, prefix, prefix_len,
			    is_ipv4, nexthop, is_nexthop_ipv4, labels, vpn_rd, vpn_rd_type,
			    path_id, origin_as, is_adj_rib_in_post_policy, is_adj_rib_out_post_policy,
			    is_loc_rib_filtered, stimestamp, rtimestamp, ltimestamp) 
			    values(` + `'` + l3v.Action + `', '` + l3v.RouterID + `', '` + l3v.RouterHash + `', 
			           '` + l3v.RouterIP + `', '` + l3v.BaseAttributes + `', '` + l3v.PeerHash + `', 
				   '` + l3v.PeerIP + `', ` + strconv.FormatInt(int64(l3v.PeerType), 10) + `, 
				   ` + strconv.FormatInt(int64(l3v.PeerASN), 10) + `, '` + l3v.Prefix + `', 
				   ` + strconv.FormatInt(int64(l3v.PrefixLen), 10) + `, 
				   ` + strconv.FormatBool(l3v.IsIPv4) + `, '` + l3v.Nexthop + `', 
				   ` + strconv.FormatBool(l3v.IsNexthopIPv4) + `, '` + l3v.Labels + `', 
				   '` + l3v.VPNRD + `', ` + strconv.FormatInt(int64(l3v.VPNRDType), 10) + `, 
				   ` + strconv.FormatInt(int64(l3v.PathID), 10) + `, 
				   ` + strconv.FormatInt(int64(l3v.OriginAS), 10) + `, 
				   ` + strconv.FormatBool(l3v.IsAdjRIBInPost) + `, 
				   ` + strconv.FormatBool(l3v.IsAdjRIBOutPost) + `, 
				   ` + strconv.FormatBool(l3v.IsLocRIBFiltered) + `, '` + l3v.Timestamp + `', 
				   '` + str2TS(l3v.Timestamp) + `',  now())`

	if glog.V(6) {
		glog.Infof("SQL -> %s\n", q0)
	}
	rows, err := db.Query(q0)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (l3v *L3VPNPrefixS) updateL3VPNV4_CurrPG(id int) {
	/*
		q0 := `update l3vpnV4_curr set action, router_id, router_hash, router_ip, base_attr_hash,
				    peer_hash, peer_ip, peer_type, peer_asn, prefix, prefix_len,
				    is_ipv4, nexthop, is_nexthop_ipv4, labels, vpn_rd, vpn_rd_type,
				    path_id, origin_as, is_adj_rib_in_post_policy, is_adj_rib_out_post_policy,
				    is_loc_rib_filtered, stimestamp, rtimestamp, ltimestamp)
				    values(` + `'` + l3v.Action + `', '` + l3v.RouterID + `', '` + l3v.RouterHash + `',
	*/
	q0 := `update l3vpnV4_curr set  base_attr_hash = '` + l3v.BaseAttributes + `', 
			    nexthop = '` + l3v.Nexthop + `', 
			    labels = '` + l3v.Labels + `', 
			    path_id = ` + strconv.FormatInt(int64(l3v.PathID), 10) + `, 
			    origin_as = ` + strconv.FormatInt(int64(l3v.OriginAS), 10) + `, 
			    stimestamp = '` + l3v.Timestamp + `', 
			    rtimestamp = '` + str2TS(l3v.Timestamp) + `',  
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

func (l3v *L3VPNPrefixS) deletePfxById_CurrPG(id int) {
	q0 := `delete from l3vpnv4_curr
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

func (l3v *L3VPNPrefixS) modifyL3VPNV4_CurrPG() {
	id := l3v.searchIdByPfx_CurrPG()
	// no row was found out
	if id == 0 {
		l3v.insertL3VPNV4_CurrPG()
	} else {
		if l3v.Action == "add" {
			l3v.updateL3VPNV4_CurrPG(id)
		} else if l3v.Action == "del" {
			l3v.deletePfxById_CurrPG(id)
		}
	}

}

func updL3VPNV4_CurrPG(l3vpnPfx *bmp_message.L3VPNPrefix, l3v L3VPNPrefixS, bas BaseAttributesS) {
	bas.baseAttr2S(*l3vpnPfx)
	l3v.l3vpnPfx2S(*l3vpnPfx)

	if glog.V(6) {
		glog.Infof("updL3VPNV4 SQL Curr -> %s\n", l3vpnPfx)
		glog.Infof("updL3VPNV4 bas SQL Curr -> %s\n", bas)
	}
	//bas.insertBaseAttr_HistPG()
	//l3v.insertPeer_HistPG()
	l3v.modifyL3VPNV4_CurrPG()
}

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
