package consumer

// Code generated by github.com/launchdarkly/go-options.  DO NOT EDIT.

import "fmt"

import "github.com/google/go-cmp/cmp"

type ApplyOptionFunc func(c *Consumer) error

func (f ApplyOptionFunc) apply(c *Consumer) error {
	return f(c)
}

func newConsumer(options ...Option) (Consumer, error) {
	var c Consumer
	err := applyConsumerOptions(&c, options...)
	return c, err
}

func applyConsumerOptions(c *Consumer, options ...Option) error {
	for _, o := range options {
		if err := o.apply(c); err != nil {
			return err
		}
	}
	return nil
}

type Option interface {
	apply(*Consumer) error
}

type withBrokerHostImpl struct {
	o string
}

func (o withBrokerHostImpl) apply(c *Consumer) error {
	c.brokerHost = o.o
	return nil
}

func (o withBrokerHostImpl) Equal(v withBrokerHostImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o withBrokerHostImpl) String() string {
	name := "WithBrokerHost"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

func WithBrokerHost(o string) Option {
	return withBrokerHostImpl{
		o: o,
	}
}

type withMessageTopicImpl struct {
	o string
}

func (o withMessageTopicImpl) apply(c *Consumer) error {
	c.messageTopic = o.o
	return nil
}

func (o withMessageTopicImpl) Equal(v withMessageTopicImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o withMessageTopicImpl) String() string {
	name := "WithMessageTopic"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

func WithMessageTopic(o string) Option {
	return withMessageTopicImpl{
		o: o,
	}
}

type withOnTextMessageImpl struct {
	o func(msg string) error
}

func (o withOnTextMessageImpl) apply(c *Consumer) error {
	c.onTextMessage = o.o
	return nil
}

func (o withOnTextMessageImpl) Equal(v withOnTextMessageImpl) bool {
	switch {
	case !cmp.Equal(o.o, v.o):
		return false
	}
	return true
}

func (o withOnTextMessageImpl) String() string {
	name := "WithOnTextMessage"

	// hack to avoid go vet error about passing a function to Sprintf
	var value interface{} = o.o
	return fmt.Sprintf("%s: %+v", name, value)
}

func WithOnTextMessage(o func(msg string) error) Option {
	return withOnTextMessageImpl{
		o: o,
	}
}
