package bmppg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"github.com/sbezverk/gobmp/pkg/kafka"
	bmp_message "github.com/sbezverk/gobmp/pkg/message"
)

var db *sql.DB
var Pg_host string
var Pg_db string
var kafka_topics []kafka_topic

type kafka_topic struct {
	Id               int
	TopicName        string
	TopicType        string
	TopicTypeId      int
	TopicDescription string
}

func Open() {
	var err error
	pgbmp_user, u := os.LookupEnv("pgbmp_user")
	pgbmp_pass, p := os.LookupEnv("pgbmp_pass")
	if !u || !p {
		glog.Infof("Environment variables pgbmp_user and/or pgbmp_pass weren't defined")
		os.Exit(1)
	}

	connStr := "host=" + Pg_host + " user=" + pgbmp_user + " password=" + pgbmp_pass + " dbname=" + Pg_db + " sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(2)
	//db.SetConnMaxIdleTime(1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//defer db.Close()
}

func Close() {
	db.Close()
}

func GetKafkaTopics() []kafka_topic {
	kafka_topics = make([]kafka_topic, 0)
	var (
		id                int
		topic_name        string
		topic_type        string
		topic_type_id     int
		topic_description string
	)
	rows, err := db.Query("select kt.id, kt.topic_name, kt.topic_type, kt.topic_type_id, kt.topic_description from kafka_topics kt")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		switch err := rows.Scan(&id, &topic_name, &topic_type, &topic_type_id, &topic_description); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			kafka_topics = append(kafka_topics, kafka_topic{
				Id:               id,
				TopicName:        topic_name,
				TopicType:        topic_type,
				TopicTypeId:      topic_type_id,
				TopicDescription: topic_description,
			})
		default:
		}
	}

	nkt := len(kafka_topics)
	db.SetMaxOpenConns(nkt)
	db.SetMaxIdleConns(nkt)

	return kafka_topics
}

func (s *store) PeerStateChangeMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) UnicastPrefixMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
	for {
		select {
		case <-s.stopCh:
			return
		case msg := <-topic.TopicChan:
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d", topic.TopicType)
			}
			u := &bmp_message.UnicastPrefix{}
			if err := json.Unmarshal(msg, u); err != nil {
				workersErrChan <- err
				return
			}
			if u.IsEOR {
				continue
			}
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d -> %s\n", topic.TopicType, u)
			}
		}
	}

}
func (s *store) UnicastPrefixV4MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
	upS := UnicastPrefixS{}
	baS := BaseAttributesS{}
	for {
		select {
		case <-s.stopCh:
			return
		case msg := <-topic.TopicChan:
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d", topic.TopicType)
			}
			u := &bmp_message.UnicastPrefix{}
			if err := json.Unmarshal(msg, u); err != nil {
				workersErrChan <- err
				return
			}
			if u.IsEOR {
				continue
			}
			upS = UnicastPrefixS{}
			baS = BaseAttributesS{}
			updUnicastPrefixV4_HistPG(u, upS, baS)
			updUnicastPrefixV4_CurrPG(u, upS, baS)
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d -> %s\n", topic.TopicType, u)
			}
		}
	}
}
func (s *store) UnicastPrefixV6MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
	for {
		select {
		case <-s.stopCh:
			return
		case msg := <-topic.TopicChan:
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d", topic.TopicType)
			}
			u := &bmp_message.UnicastPrefix{}
			if err := json.Unmarshal(msg, u); err != nil {
				workersErrChan <- err
				return
			}
			if u.IsEOR {
				continue
			}
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d -> %s\n", topic.TopicType, u)
			}
		}
	}
}
func (s *store) LSNodeMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) LSLinkMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) L3VPNMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
	for {
		select {
		case <-s.stopCh:
			return
		case msg := <-topic.TopicChan:
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d", topic.TopicType)
			}
			u := &bmp_message.L3VPNPrefix{}
			if err := json.Unmarshal(msg, u); err != nil {
				workersErrChan <- err
				return
			}
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d -> %s\n", topic.TopicType, u)
			}
		}
	}
}
func (s *store) L3VPNV4MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
	l3vpnS := L3VPNPrefixS{}
	baS := BaseAttributesS{}
	for {
		select {
		case <-s.stopCh:
			return
		case msg := <-topic.TopicChan:
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d", topic.TopicType)
			}
			u := &bmp_message.L3VPNPrefix{}
			if err := json.Unmarshal(msg, u); err != nil {
				workersErrChan <- err
				return
			}
			l3vpnS = L3VPNPrefixS{}
			baS = BaseAttributesS{}
			updL3VPNV4_HistPG(u, l3vpnS, baS)
			updL3VPNV4_CurrPG(u, l3vpnS, baS)
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d -> %s\n", topic.TopicType, u)
			}
		}
	}
}
func (s *store) L3VPNV6MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
	for {
		select {
		case <-s.stopCh:
			return
		case msg := <-topic.TopicChan:
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d", topic.TopicType)
			}
			u := &bmp_message.L3VPNPrefix{}
			if err := json.Unmarshal(msg, u); err != nil {
				workersErrChan <- err
				return
			}
			if glog.V(6) {
				glog.Infof("Store received message from topic type: %d -> %s\n", topic.TopicType, u)
			}
		}
	}
}
func (s *store) LSPrefixMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) LSSRv6SIDMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) EVPNMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) SRPolicyMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) SRPolicyV4MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) SRPolicyV6MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) FlowspecMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) FlowspecV4MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) FlowspecV6MsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
func (s *store) StatsReportMsgPG(topic *kafka.TopicDescriptor, done chan struct{}, workersErrChan chan error) {
}
