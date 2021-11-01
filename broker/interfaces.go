package broker

type Broker interface {
	SubscribePlayerReadyEventV1()
	PublishPlayerReadyEventV1(playerId uint8)
}
