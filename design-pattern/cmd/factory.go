/**
 * Factory Method Pattern
 * Create a gun factory that can produce different types of guns.
 */
package main

import "fmt"

/**
 * Base interface for all gun products. + Base Methods for all gun products.
 */

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
	getSound() string
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

/**
 * Concrete Products
 */

// Ak47 is a concrete product
type Ak47 struct {
	Gun          // embed Gun, Ak47 is a Gun, so it has all the Gun's methods
	sound string // Ak47 has an additional field
}

// rewrite old method for Ak47
func (g *Ak47) getName() string {
	return g.name + "!"
}

// write less method for interface
func (g *Ak47) getSound() string {
	return g.sound
}

// M16 is a concrete product
type M16 struct {
	Gun // embed Gun
}

// write less method for interface
func (g *M16) getSound() string {
	return "Tak Tak Tak..."
}

/**
 * Concrete Creator
 */

// GunFactory is a concrete creator
type GunFactory struct{}

// MakeGun is a factory method
func (g *GunFactory) MakeGun(gunType string) IGun {
	switch gunType {
	case "ak47":
		return &Ak47{Gun{"AK47", 4}, "Bang Bang Bang..."}
	case "m16":
		return &M16{Gun{"M16", 5}}
	default:
		return nil
	}
}

// use interface as a parameter
func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}

func main() {
	gunFactory := &GunFactory{}

	ak47 := gunFactory.MakeGun("ak47")
	m16 := gunFactory.MakeGun("m16")

	println(ak47.getName())
	println(ak47.getPower())

	println(m16.getName())
	println(m16.getPower())

	// set power
	ak47.setPower(10)
	println(ak47.getPower())

	// set name
	ak47.setName("AK-47")
	println(ak47.getName())

	// print details
	printDetails(ak47)
	printDetails(m16)

	// print sound
	println(ak47.getSound())
	println(m16.getSound())
}
