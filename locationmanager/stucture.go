/*
 * Structures for the locationmanager.
 *
 * These structures track properties of objects managed by the locationmanager.
 */

// Package locationmanager provides all abilities to detect other entities in an environment.
package locationmanager

import "github.com/DiscoViking/goBrains/entity"

// Coord structs hold a specific co-ordinate (point) within the environment.
type coord struct {
	locX, locY float64
}

// CircleHitbox represents a circular hitbox (handy, right?)
// They have two active values, representing the centre of the hitbox, and the radius.  It also references the entity it represents.
type circleHitbox struct {

	// Whether the hitbox is currently in use.
	active bool

	// Active values, holding state.
	centre      coord
	orientation float64
	radius      float64

	// External reference, to the entity that the hitbox represents.
	entity entity.Entity
}

// CoordDelta structs represent a position relative to an entity.
type CoordDelta struct {
	distance float64
	rotation float64
}

// LocationManager is an instance of a location manager.
// It holds all the state about entities in the environment.
type LocationManager struct {
	hitboxes []locatable
}