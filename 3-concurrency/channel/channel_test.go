package channel_test

import (
	"testing"

	"github.com/ramamimu/go-everything/3-concurrency/channel"
)

func TestChannel(t *testing.T) {
	channel.ChannelImpl()
}

func TestIteration(t *testing.T) {
	channel.Iteration()
}

func TestChannelWithTimeOut(t *testing.T) {
	channel.ChannelWithTimeout()
}

func TestChannelGroup(t *testing.T) {
	channel.ChannelGroup()
}
