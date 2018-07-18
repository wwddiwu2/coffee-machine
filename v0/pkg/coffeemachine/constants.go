package coffeemachine

// Drink represents a drink that the machine can pour
type Drink uint

const (
	// Espresso stands for an Espresso
	Espresso Drink = iota
	// Americano stands for an Americano
	Americano
)

// Status represents the current state a coffe machine can be in
type Status uint8

const (
	// Ready represents that a machine is ready to perform some action
	Ready Status = iota
	// Pouring represents that a machine is currently pouring coffee
	Pouring
	// Cleaning represents that a machine is in the cleaning process
	Cleaning
)
