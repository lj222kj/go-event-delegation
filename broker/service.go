package broker

import (
	"errors"
	"fmt"
	"sync"
)

type broker struct {
	mutex sync.Mutex
	subs  map[string][]chan []byte
}

const playerReadyEventV1 = "PLAYER_READY_EVENT_V1"

func (s *broker) subscribe(channel string, ch chan []byte) <-chan []byte {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.subs[channel] = append(s.subs[channel], ch)
	return ch
}

func (s *broker) publish(channel string, data []byte) error {
	_, ok := s.subs[channel]
	if !ok {
		return errors.New(fmt.Sprintf("%s has no subscribers", channel))
	}

	for _, ch := range s.subs[channel] {
		go func(ch chan []byte) {
			ch <- data
		}(ch)
	}
	return nil
}

func (b *broker) PublishPlayerReadyEventV1(playerId uint8) error {
	return b.publish(playerReadyEventV1, []byte{playerId})
}

func (b *broker) SubscribePlayerReadyEventV1(ch chan []byte) <-chan []byte {
	return b.subscribe(playerReadyEventV1, ch)
}

func New() *broker {
	subs := make(map[string][]chan []byte)
	return &broker{
		mutex: sync.Mutex{},
		subs:  subs,
	}
}
