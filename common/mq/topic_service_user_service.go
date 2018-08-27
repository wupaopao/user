package mq

import (
	"github.com/mz-eco/mz/kafka"
	"github.com/mz-eco/mz/log"
)

const (
	TOPIC_SERVICE_USER_SERVICE = "service-user-service"
)

const (
	IDENT_SERVICE_USER_SERVICE_MODIFY_USER_INFO = "modify_user_info"
)

var (
	topicServiceUserService *TopicServiceUserService = nil
)

func GetTopicServiceUserService() (topic *TopicServiceUserService, err error) {
	if topicServiceUserService != nil {
		topic = topicServiceUserService
		return
	}

	producer, err := kafka.NewAsyncProducer()
	if err != nil {
		log.Warnf("new async producer failed. %s", err)
		return
	}

	topicServiceUserService = &TopicServiceUserService{
		Producer: producer,
	}

	topic = topicServiceUserService

	return
}

type TopicServiceUserService struct {
	Producer *kafka.AsyncProducer
}

func (m *TopicServiceUserService) send(ident string, msg interface{}) (err error) {
	err = m.Producer.SendMessage(TOPIC_SERVICE_USER_SERVICE, ident, msg)
	if err != nil {
		log.Warnf("send topic message failed. %s", err)
		return
	}
	return
}

// 用户信息修改
type ModifyUserInfoMessage struct {
	UserId string
	Values map[string]interface{}
}

func (m *TopicServiceUserService) ModifyUserInfo(msg *ModifyUserInfoMessage) (err error) {
	return m.send(IDENT_SERVICE_USER_SERVICE_MODIFY_USER_INFO, msg)
}
