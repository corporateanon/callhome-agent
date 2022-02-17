package consumer

type Options struct {
	BrokerHost    string
	MessageTopic  string
	OnTextMessage func(msg string) error
}

func NewOptions() *Options {
	return &Options{
		OnTextMessage: func(msg string) error { return nil },
	}
}

func (o *Options) SetBrokerHost(host string) *Options {
	o.BrokerHost = host
	return o
}

func (o *Options) SetMessageTopic(topic string) *Options {
	o.MessageTopic = topic
	return o
}

func (o *Options) SetOnTextMessage(callback func(msg string) error) *Options {
	o.OnTextMessage = callback
	return o
}
