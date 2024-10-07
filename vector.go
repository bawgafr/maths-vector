package vector

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type Vector struct {
	X, Y, Z float64
}

//test
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
// *heading() calcs the angle a 2d vector makes with the positive x axis. Angles increase clockwise
// *rotate(float64) rotates a vector without changing magnitude
// *random2d() -- creates a new 2d unit vector with a random heading
// *random3d() -- creates a new 3d unit vector with a random heading
// *fromAngle(float64) -- creates a 2d vector from the passed angle

// setHeading() rotates a 2d vector to a specific angle without changing magnitude

// lerp, slerp

func (v Vector) String() string {
	return fmt.Sprintf("{%2f, %2f, %2f}", v.X, v.Y, v.Z)
}

// check if the components of the two vectors are the same
func Equals(v1, v2 Vector) bool {

	x := math.Abs(v1.X - v2.X)
	y := math.Abs(v1.Y - v2.Y)
	z := math.Abs(v1.Z - v2.Z)

	return x < 1e-9 && y < 1e-9 && z < 1e-9
}

// check if the passed Vector has the same components as this Vector
func (v1 Vector) Equals(v2 Vector) bool {
	x := math.Abs(v1.X - v2.X)
	y := math.Abs(v1.Y - v2.Y)
	z := math.Abs(v1.Z - v2.Z)

	return x < 1e-9 && y < 1e-9 && z < 1e-9
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

// create a unit vector in a random direction
func Random2d() Vector {
	v := NewVector(rand.Float64(), rand.Float64())
	v.Normalise()
	return v
}

func Random3d() Vector {
	v := NewVector(rand.Float64(), rand.Float64(), rand.Float64())
	v.Normalise()
	return v
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
	}
	if l > 1 {
		v.Y = values[1]
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

// angle 2d vector makes with with positive x axis. Angle increases clockwise
func Heading(v Vector) float64 {
	base := NewVector(10, 0)
	angle := AngleBetween(v, base)

	return angle
}

// angle this 2d vector makes with the positive x axis
func (v Vector) Heading() float64 {
	base := NewVector(10, 0)
	return base.AngleBetween(v)
}

// sets the angle of the vector without changing its magnitude
func Rotate(v Vector, angle float64) Vector {
	// x2 = cos()x1 - sin()y1
	// y2 = sin()x1 + cos()y1

	c := math.Cos(-angle)
	s := math.Sin(-angle)

	return NewVector(c*v.X-s*v.Y, s*v.X+c*v.Y)
}

// rotates the vector by angle
func (v *Vector) Rotate(angle float64) {
	c := math.Cos(-angle)
	s := math.Sin(-angle)

	v.X = c*v.X - s*v.Y
	v.Y = s*v.X + c*v.Y
}

// creates a vector of length l in the direction angle
//
// FromAngle(Angle float64, length float64). If length omitted then unit vector created
func FromAngle(values ...float64) Vector {
	angle := -1.0 * values[0]
	length := 1.0
	if len(values) == 2 {
		length = values[1]
	}

	x := math.Cos(angle)
	y := math.Sin(angle)

	v := NewVector(x, y)

	v.Normalise()
	if length != 1 {
		v.Mult(length)
	}

	switch quadrant(angle) {
	case "ne":
		v.X = math.Abs(v.X)
		v.Y = math.Abs(v.Y)
	case "se":
		v.X = math.Abs(v.X)
		v.Y = -1.0 * math.Abs(v.Y)
	case "sw":
		v.X = -1.0 * math.Abs(v.X)
		v.Y = -1.0 * math.Abs(v.Y)
	case "nw":
		v.X = -1.0 * math.Abs(v.X)
		v.Y = math.Abs(v.Y)
	}

	return v
}

func quadrant(angle float64) string {
	angle = math.Mod(angle, 2*math.Pi)

	if angle >= 0 && angle <= math.Pi/2 {
		log.Printf("%.2fπ  %s", angle, "se")
		return "se"
	}

	if angle >= math.Pi/2 && angle <= math.Pi {
		log.Printf("%.2fπ  %s", angle, "sw")
		return "sw"
	}

	if angle >= math.Pi && angle < 3*math.Pi/2 {
		log.Printf("%.2fπ  %s", angle, "nw")
		return "nw"
	}
	log.Printf("%.2fπ  %s", angle, "ne")

	return "ne"

}
