// Package coffeemachine offers the possibility to create a simple mock coffee machine
package coffeemachine

import (
	"errors"
)

// Machine is a struct which represents the coffee machine
type Machine struct {
	cleanliness   uint8
	status        Status
	StatusChannel chan Status
}

// New returns a brand new coffee machine
func New() Machine {
	return Machine{
		cleanliness:   100,
		status:        Ready,
		StatusChannel: make(chan Status, 1),
	}
}

// Pour pours the give drink
func (m *Machine) Pour(d Drink) error {
	if m.status != Ready {
		return errors.New("coffeemachine: the coffee machine is not yet ready to pour")
	}
	m.changeStatus(Pouring)
	if d == Americano {
		m.cleanliness -= 2
	} else {
		m.cleanliness--
	}
	m.changeStatus(Ready)
	return nil
}

// Clean resets the cleanliness value to 100
func (m *Machine) Clean() error {
	if m.status != Ready {
		return errors.New("coffeemachine: the coffee machine is not yet ready to clean")
	}
	m.changeStatus(Cleaning)
	m.cleanliness = 100
	m.changeStatus(Ready)
	return nil
}

// Cleanliness returns the cleanliness of the coffee machine
func (m Machine) Cleanliness() uint8 {
	return m.cleanliness
}

// Status returs the current status of the coffee machine
func (m Machine) Status() Status {
	return m.status
}

func (m Machine) changeStatus(s Status) {
	m.status = s
	select {
	case m.StatusChannel <- s:
	case <-m.StatusChannel:
		m.StatusChannel <- s
	}
}
