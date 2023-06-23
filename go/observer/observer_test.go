package observer

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var got = make([][]byte, 0) // For storing received messages to assert later
var wg sync.WaitGroup       // For waiting goroutines to finish in test

// SubscriberImpl implements Subscriber interface
type SubscriberImpl struct {
}

// Update is called when a message is received
func (s SubscriberImpl) Update(m Message) {
	got = append(got, m.Data)
	defer wg.Done()
}

func TestObserver(t *testing.T) {
	sub1 := SubscriberImpl{}           // Create a new subscriber
	sub2 := SubscriberImpl{}           // Create another
	broker := NewPublisher(sub1, sub2) // Create a broker with 2 subscribers

	wg.Add(6) // We will send 3 messages to 2 subscribers so 6
	for i := 0; i < 3; i++ {
		go broker.Pub(Message{Data: []byte("Hello")}) // Publish messages
	}
	wg.Wait() // Wait for goroutines to finish

	assert.Len(t, got, 6)
	for _, v := range got {
		assert.Equal(t, []byte("Hello"), v)
	}
}
