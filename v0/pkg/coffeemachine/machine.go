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

// Brew brews the given drink
func (m *Machine) Brew(b Beverage) error {
	bev := b.getBeverage()
	if m.status != Ready {
		return errors.New("coffeemachine: the coffee machine is not yet ready to brew")
	}
	if int(m.cleanliness)-int(bev.erosion) < 0 {
		return errors.New("coffeemachine: the coffee machine is to dirty to brew")
	}
	m.changeStatus(Brewing)
	m.cleanliness -= bev.erosion
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
