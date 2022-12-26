package es

import "fmt"

const (
	aggregateStartVersion                = -1 // used for EventStoreDB
	aggregateAppliedEventsInitialCap     = 10
	aggregateUncommittedEventsInitialCap = 10
)

type Aggregate interface {
	When
	AggregateRoot
}

// AggregateRoot contains all methods of AggregateBase
type AggregateRoot interface {
	GetUncommittedEvents() []Event
	GetID() string
	SetID(id string) *AggregateBase
	GetVersion() int64
	ClearUncommittedEvents()
	ToSnapshot()
	SetType(aggregateType AggregateType)
	GetType() AggregateType
	SetAppliedEvents(events []Event)
	GetAppliedEvents() []Event
	RaiseEvent(event Event) error
	String() string
	Load
	Apply
}

// Apply process Aggregate Event
type Apply interface {
	Apply(event Event) error
}

// Load create Aggregate state from Event's.
type Load interface {
	Load(events []Event) error
}

type When interface {
	When(event Event) error
}

type when func(event Event) error

// AggregateType type of the Aggregate
type AggregateType string

type AggregateBase struct {
	ID                string
	Version           int32
	AppliedEvents     []Event
	UncommittedEvents []Event
	Type              AggregateType
	withAppliedEvents bool
	when              when
}

func NewAggregateBase(when when) *AggregateBase {
	if when == nil {
		return nil
	}

	return &AggregateBase{
		Version:           aggregateStartVersion,
		AppliedEvents:     make([]Event, 0, aggregateAppliedEventsInitialCap),
		UncommittedEvents: make([]Event, aggregateUncommittedEventsInitialCap),
		when:              when,
		withAppliedEvents: false,
	}
}

// SetID set AggregateBase ID
func (a *AggregateBase) SetID(id string) *AggregateBase {
	a.ID = fmt.Sprintf("%s-%s", a.GetType(), id)
	return a
}

// GetID get AggregateBase ID
func (a *AggregateBase) GetID() string {
	return a.ID
}

// SetType set AggregateBase AggregateType
func (a *AggregateBase) SetType(aggregateType AggregateType) {
	a.Type = aggregateType
}

// GetType get AggregateBase AggregateType
func (a *AggregateBase) GetType() AggregateType {
	return a.Type
}

// GetVersion get AggregateBase version
func (a *AggregateBase) GetVersion() int32 {
	return a.Version
}

// ClearUncommittedEvents clear AggregateBase uncommitted Event's
func (a *AggregateBase) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]Event, 0, aggregateUncommittedEventsInitialCap)
}

// GetAppliedEvents get AggregateBase applied Event's
func (a *AggregateBase) GetAppliedEvents() []Event {
	return a.AppliedEvents
}

// SetAppliedEvents set AggregateBase applied Event's
func (a *AggregateBase) SetAppliedEvents(events []Event) {
	a.AppliedEvents = events
}

// GetUncommittedEvents get AggregateBase uncommitted Event's
func (a *AggregateBase) GetUncommittedEvents() []Event {
	return a.UncommittedEvents
}
