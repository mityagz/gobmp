CREATE TABLE kafka_topics (
  	id SERIAL NOT NULL,
	topic_name VARCHAR(255) UNIQUE,				--":"b17f50eb425c66f7adc2b98afffb888b"
	topic_type VARCHAR(255),					--":"10.229.170.0",
	topic_type_id INTEGER,					--":0,
	topic_description VARCHAR(255),					--":"10.229.170.0",
  	PRIMARY KEY(topic_name)
);

CREATE TABLE peers (
  	id SERIAL NOT NULL,
	peer_hash VARCHAR(255) UNIQUE,				--":"b17f50eb425c66f7adc2b98afffb888b"
	peer_ip  VARCHAR(255),					--":"10.229.170.0",
	peer_type INTEGER,					--":0,
	peer_asn INTEGER,					--":3333,
	stimestamp  VARCHAR(255),				--":"2025-06-18T17:28:12Z"
	rtimestamp timestamp with time zone,
	ltimestamp timestamp with time zone,
  	PRIMARY KEY(id, peer_hash)
);

CREATE TABLE base_attrs (
  	id SERIAL NOT NULL,
	base_attr_hash VARCHAR(255) UNIQUE,			--":"7a179b6d0631fcd21c5e533419d2b9ab"
	origin VARCHAR(255),					--:"igp"
	as_path	VARCHAR(255),
        as_path_count INTEGER,					--      int32    `json:"as_path_count,omitempty"`
        nexthop VARCHAR(255),					--      string   `json:"nexthop,omitempty"`
        med INTEGER,						--             uint32   `json:"med,omitempty"`
        local_pref INTEGER,					--        uint32   `json:"local_pref,omitempty"`
	is_atomic_agg  BOOLEAN,					--":false,"
        aggregator VARCHAR(255),				--      []byte   `json:"aggregator,omitempty"`
	community_list VARCHAR(255),				--":["3333:10300","3333:31990","3333:32990","3333:38080"],"
	originator_id VARCHAR(255),				--":"192.168.33.131","
	cluster_list VARCHAR(255),				--":"10.229.170.0, 10.233.241.128","
	ext_community_list VARCHAR(255),			--":["rt=3333:20545"]}
        as4_path VARCHAR(255),					--          []uint32 `json:"as4_path,omitempty"`
        as4_path_count INTEGER,					--    int32    `json:"as4_path_count,omitempty"`
        as4_aggregator VARCHAR(255),				--    []byte   `json:"as4_aggregator,omitempty"`
        -- // PMSITunnel
        tunnel_encap_attr VARCHAR(255),				--[]byte `json:"-"`
        -- // TraficEng
        -- // IPv6SpecExtCommunity
        -- // AIGP
        -- // PEDistinguisherLable
        large_community_list VARCHAR(255),
        -- // SecPath
        --// AttrSet
	stimestamp  VARCHAR(255),				--":"2025-06-18T17:28:12Z"
	rtimestamp timestamp with time zone,
	ltimestamp timestamp with time zone,
  	PRIMARY KEY(id, base_attr_hash)
);

CREATE TABLE l3vpnV4 (
  	id SERIAL NOT NULL,
        -- Sequence       int                 `json:"sequence,omitempty"`
        -- Hash           string              `json:"hash,omitempty"`
	action VARCHAR(255),					--":"add",
	router_id VARCHAR(255), 				--":"10.229.132.0:52724:PE1:209"
	router_hash VARCHAR(255),				--":"7c74731780fbc7fcc903929fe2e64ef8"
	router_ip VARCHAR(255),					--":"10.49.4.1"
	base_attr_hash VARCHAR(255),
	peer_hash VARCHAR(255),					--":"b17f50eb425c66f7adc2b98afffb888b"
	peer_ip VARCHAR(255),					--":"10.229.170.0"
	peer_type INTEGER NOT NULL,				--":0
	peer_asn INTEGER NOT NULL,				--":3333
	prefix VARCHAR(255),					--":"62.192.1.30"
	prefix_len INTEGER NOT NULL,				--":32
	is_ipv4 BOOLEAN,					--":true
	nexthop VARCHAR(255),					--":"10.229.4.0"
	is_nexthop_ipv4 BOOLEAN,				--":true
	labels VARCHAR(255),					--":[33]
	vpn_rd VARCHAR(255),					--":"10.229.4.0:2054"
	vpn_rd_type INTEGER NOT NULL,				--":1
        -- PrefixSID      *prefixsid.PSid     `json:"prefix_sid,omitempty"`
        path_id INTEGER,					--`json:"path_id,omitempty"`
        origin_as INTEGER,					--`json:"origin_as,omitempty"`
	is_adj_rib_in_post_policy BOOLEAN,			--":true
	is_adj_rib_out_post_policy BOOLEAN,			--":false
	is_loc_rib_filtered BOOLEAN,				--":false}
	stimestamp  VARCHAR(255),				--":"2025-06-18T17:28:12Z"
	rtimestamp timestamp with time zone,
	ltimestamp timestamp with time zone,
  	PRIMARY KEY(id, router_id, vpn_rd, prefix, prefix_len),
  	FOREIGN KEY (base_attr_hash) REFERENCES base_attrs (base_attr_hash),
  	FOREIGN KEY (peer_hash) REFERENCES peers (peer_hash)
);

-- select action, router_id, vpn_rd, prefix, peer_ip, cluster_list, ext_community_list, l.rtimestamp from l3vpnv4 l, base_attrs b where prefix like '62.192.1.225' and l.base_attr_hash = b.base_attr_hash and vpn_rd like '%2054' and router_id like '10.229.134.0%' group by action, router_id, vpn_rd, prefix, peer_ip, cluster_list, ext_community_list, l.rtimestamp order by l.rtimestamp;

-- select action, router_id, vpn_rd, prefix, peer_ip, cluster_list, ext_community_list, l.rtimestamp, l.ltimestamp from l3vpnv4 l, base_attrs b where prefix like '62.192.1.225' and l.base_attr_hash = b.base_attr_hash and vpn_rd like '%2054' and router_id like '10.229.134.0%' group by action, router_id, vpn_rd, prefix, peer_ip, cluster_list, ext_community_list, l.rtimestamp, l.ltimestamp order by l.rtimestamp, l.ltimestamp;
