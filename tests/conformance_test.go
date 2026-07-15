package events_test

import (
	"testing"

	"github.com/tinywasm/events"
	"github.com/tinywasm/events/conformance"
	"github.com/tinywasm/events/mock"
)

// TestMockConformance holds the reference in-process broker to the same contract
// every events.Broker must meet — the same discipline router/mock proves against
// router/conformance.
func TestMockConformance(t *testing.T) {
	conformance.Run(t, conformance.Factory{
		New: func(t *testing.T) events.Broker { return &mock.Broker{} },
	})
}
