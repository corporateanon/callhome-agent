package consumer

type Options struct {
	brokerHost    string
	messageTopic  string
	onTextMessage func(msg string) error
}

func NewOptions() *Options {
	return &Options{
		onTextMessage: func(msg string) error { return nil },
	}
}

func (o *Options) SetBrokerHost(host string) *Options {
	o.brokerHost = host
	return o
}

func (o *Options) SetMessageTopic(topic string) *Options {
	o.messageTopic = topic
	return o
}

func (o *Options) SetOnTextMessage(callback func(msg string) error) *Options {
	o.onTextMessage = callback
	return o
}
