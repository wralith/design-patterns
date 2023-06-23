package observer

import "sync"

// Message is the object that will be passed to the subscribers
type Message struct {
	Data []byte
}

// Publisher is the  object that will manage the subscribers and publish the messages
type Publisher struct {
	Subscribers []Subscriber
}

// Subscriber is the interface that all subscribers must implement
type Subscriber interface {
	Update(Message)
}

// NewPublisher creates a new broker with the given subscribers,  if no subscribers are given it will create an empty slice
func NewPublisher(subscribers ...Subscriber) *Publisher {
	if subscribers == nil {
		subscribers = make([]Subscriber, 0)
	}
	return &Publisher{
		Subscribers: subscribers,
	}
}

// Sub adds a subscriber to the broker, if the subscriber is already in the broker won't do anything
func (b *Publisher) Sub(sub Subscriber) {
	for _, alreadySub := range b.Subscribers {
		if sub == alreadySub {
			return
		}
	}
	b.Subscribers = append(b.Subscribers, sub)
}

// Pub executes the Update method of each subscriber
func (b *Publisher) Pub(message Message) {
	// Some concurrency things I don't even know if I test those right...
	wg := sync.WaitGroup{}
	for _, sub := range b.Subscribers {
		wg.Add(1)
		go func(sub Subscriber) {
			defer wg.Done()
			sub.Update(message)
		}(sub)
	}
}
