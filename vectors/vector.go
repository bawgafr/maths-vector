package vector

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

// p5js has the following methods:
// (https://p5js.org/reference/p5/p5.Vector/)
// *set(x,y,z) set() and all in between
// *copy()
// *add(Vector)
// rem -- modulo
// *sub(Vector)
// *mult(float64)
// *dev(float64)
// *mag() -- calculates the magnitude of the vector
// *magSq() -- calculates the square of the magnitude
// *dot(Vector) -- dot product of 2 2d vectors
// *dist(Vector) -- the distance between 2 vectors
// *normalise -- scales the vector so the mag = 1
// *limit(float64) sets a max value for magnitude if the value is > than it
// **angleBetween(Vector) returns the angle between this and the passe vector
// *equals(Vector) -- x==X && y==Y && z==Z
// *setMag(float64) sets the magnitude of the vector

// heading() calcs the angle a 2d vector makes with the positive x axis. Angles increase clockwise
// setHeading() rotates a 2d vector to a specific angle without changing magnitude
// rotate(float64) rotates a vector without changing magnitude
// lerp, slerp
// fromAngle(float64) -- creates a 2d vector from the passed angle
// random2d() -- creates a new 2d unit vector with a random heading
// random3d() -- creates a new 3d unit vector with a random heading

func (v Vector) String() string {
	return fmt.Sprintf("{%2f, %2f, %2f}", v.X, v.Y, v.Z)
}

// check if the components of the two vectors are the same
func Equals(v1, v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

// check if the passed Vector has the same components as this Vector
func (v1 Vector) Equals(v2 Vector) bool {
	return v1.X == v2.X && v1.Y == v2.Y && v1.Z == v2.Z
}

func NewVector(values ...float64) Vector {
	x := 0.0
	y := 0.0
	z := 0.0

	l := len(values)
	if l > 0 {
		x = values[0]
	}
	if l > 1 {
		y = values[1]
	}
	if l > 2 {
		z = values[2]
	}

	return Vector{x, y, z}
}

// sets the values of the components
//
// Set() will set the components to {0,0,0}
// Set(a) will set {a, 0, 0}.
// Set(a,b) will set {a, b, 0}.
// Set(a, b, c) will set {a, b, c}
func (v *Vector) Set(values ...float64) {
	l := len(values)

	if l == 0 {
		v.X = 0.0
		v.Y = 0.0
		v.Z = 0.0
	}

	if l > 0 {
		v.X = values[0]
		v.Y = 0
		v.Z = 0
	}
	if l > 1 {
		v.Y = values[1]
		v.Z = 0
	}
	if l > 2 {
		v.Z = values[2]
	}

}

// returns a new copy of the vector
func (v Vector) Copy() Vector {
	return Vector{v.X, v.Y, v.Z}
}

// adds the two vectors and retuns a new Vector
func Add(v1, v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// adds the vector to this one
func (v *Vector) Add(other Vector) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

// subtract the two vectors and return a new Vector
func Sub(v1, v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// subtract the vector from this one
func (v *Vector) Sub(other Vector) {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
}

// multiply the vector by m and return a new Vector
func Mult(v Vector, m float64) Vector {
	return Vector{v.X * m, v.Y * m, v.Z * m}
}

// multiply this vector by m
func (v *Vector) Mult(m float64) {
	v.X *= m
	v.Y *= m
	v.Z *= m
}

// scalar divide the vector by d
func Div(v Vector, d float64) Vector {
	return Vector{v.X / d, v.Y / d, v.Z / d}
}

// scalar divide this by amount d
func (v *Vector) Div(d float64) {
	v.X /= d
	v.Y /= d
	v.Z /= d
}

// returns the magnitude of the passed in Vector
func Mag(v Vector) float64 {
	return math.Sqrt(MagSq(v))
}

// returns the magnitude squared of the passed Vector
func MagSq(v Vector) float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// return the magnitude squared of this vector
func (v Vector) MagSq() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

// return the magnitude of this vector
func (v Vector) Mag() float64 {
	return math.Sqrt(v.MagSq())
}

// angle between 2 vectors
func AngleBetween(v1, v2 Vector) float64 {
	// acos( (v1.v2)/(|v1| |v2|)
	v1m := v1.Mag()
	v2m := v2.Mag()

	dp := v1.DotProduct(v2)

	return math.Acos(dp / (v1m * v2m))
}

// angle between passed in vector and this vector
func (v Vector) AngleBetween(other Vector) float64 {
	v1m := v.Mag()
	v2m := other.Mag()

	dp := v.DotProduct(other)

	return math.Acos(dp / (v1m * v2m))
}

// returns the dot product of the Vectors
func DotProduct(v1, v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// returns the dot product of this vector with the passed in one
func (v Vector) DotProduct(other Vector) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

// Distance between the two vectors
func Dist(v1, v2 Vector) float64 {
	dx := v1.X - v2.X
	dy := v1.Y - v2.Y
	dz := v1.Z - v2.Z

	sq := (dx * dx) + (dy * dy) + (dz * dz)

	return math.Sqrt(sq)
}

// distance between this vector and one passed in
func (v Vector) Dist(other Vector) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z

	sq := (dx * dx) + (dy * dy) + (dz * dz)

	return math.Sqrt(sq)
}

// normalise the vector
func Normalise(v Vector) Vector {
	m := v.Mag()

	return Div(v, m)
}

// normalise this vector
func (v *Vector) Normalise() {
	m := v.Mag()

	v.Div(m)
}

// limits the magnitude of the vector to passed in float64
// todo: work out if inlining the normalise if faster
func Limit(v Vector, l float64) Vector {
	m := v.Mag()

	if m <= l {
		return v
	}

	n := Normalise(v)
	return Mult(n, l)
}

// limit the magnitude of this vector
func (v *Vector) Limit(l float64) {
	m := v.Mag()
	if m <= l {
		return
	}

	v.Normalise()
	v.Mult(l)
}

// set magnitude of the vector
func SetMag(v Vector, m float64) Vector {
	n := Normalise(v)
	return Mult(n, m)
}

// set the magnitude of this vector
func (v *Vector) SetMag(m float64) {
	v.Normalise()
	v.Mult(m)
}
