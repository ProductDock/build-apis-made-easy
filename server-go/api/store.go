package api

import (
	"errors"
	"strings"
)

type EventStorer interface {
	Add(e Event) (Event, error)
	GetAll() []Event
	FindById(id string) (*Event, error)
	DeleteById(id string) bool
	FindByNameFuzzy(name string) []Event
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		events: make([]Event, 0),
	}
}

type InMemoryEventStore struct {
	events []Event
}

func (i *InMemoryEventStore) Add(e Event) (Event, error) {
	i.events = append(i.events, e)
	return e, nil
}

func (i *InMemoryEventStore) GetAll() []Event {
	return i.events
}

func (i *InMemoryEventStore) DeleteById(id string) bool {
	for index, evt := range i.events {
		if evt.Id == id {
			i.events = append(i.events[:index], i.events[index+1:]...)
			return true
		}
	}
	return false
}

func (i *InMemoryEventStore) FindById(id string) (*Event, error) {
	for _, event := range i.events {
		if event.Id == id {
			return &event, nil
		}
	}
	return nil, errors.New("event not found")
}

func (i *InMemoryEventStore) FindByNameFuzzy(name string) []Event {
	filteredEvents := make([]Event, 0)
	for _, e := range i.events {
		if strings.ContainsAny(e.Name, name) {
			filteredEvents = append(filteredEvents, e)
		}
	}
	return filteredEvents
}
