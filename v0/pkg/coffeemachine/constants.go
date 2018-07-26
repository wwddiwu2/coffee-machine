package coffeemachine

type beverage struct {
	id      uint
	erosion uint8
	name    string
}

var beverages = [...]beverage{
	beverage{
		id:      0,
		erosion: 1,
		name:    "Espresso",
	},
	beverage{
		id:      1,
		erosion: 2,
		name:    "Americano",
	},
}

// Beverage represents a beverage that the machine can brew
type Beverage uint

// ID gets the ID of the beverage
func (b Beverage) ID() uint {
	return b.getBeverage().id
}

// Name gets the name of the beverage
func (b Beverage) Name() string {
	return b.getBeverage().name
}

func (b Beverage) getBeverage() beverage {
	return beverages[b]
}

const (
	// Espresso stands for an Espresso
	Espresso Beverage = iota
	// Americano stands for an Americano
	Americano
)

// Status represents the current state a coffe machine can be in
type Status uint8

const (
	// Ready represents that a machine is ready to perform some action
	Ready Status = iota
	// Brewing represents that a machine is currently brewing coffee
	Brewing
	// Cleaning represents that a machine is in the cleaning process
	Cleaning
)
