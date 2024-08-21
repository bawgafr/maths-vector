package vector

import (
	"math"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("create an empty vector", func(t *testing.T) {
		v := NewVector()

		if v.X != 0 || v.Y != 0 || v.Z != 0 {
			t.Error("components not equal to zero", v)
		}

		zeroVect := Vector{0.0, 0.0, 0.0}

		if !Equals(v, zeroVect) {
			t.Error("did not create a zero vector")
		}
	})
	t.Run("Test 1d Vector", func(t *testing.T) {
		v := NewVector(1.2)
		if v.X != 1.2 || v.Y != 0 || v.Z != 0 {
			t.Error("1d vector not created", v)
		}

	})
	t.Run("Test 2d Vector", func(t *testing.T) {
		v := NewVector(1.2, 7.9)
		if v.X != 1.2 || v.Y != 7.9 || v.Z != 0 {
			t.Error("1d vector not created", v)
		}

	})
	t.Run("Test 3d Vector", func(t *testing.T) {
		v := NewVector(1.2, 7.1, 6.4)
		if v.X != 1.2 || v.Y != 7.1 || v.Z != 6.4 {
			t.Error("1d vector not created", v)
		}

	})

}

func TestDotProduct(t *testing.T) {
	v1 := NewVector(3, 4)
	v2 := NewVector(3, 0)

	dp := DotProduct(v1, v2)

	if dp != 9.0 {
		t.Errorf("dp not working yet... %f", dp)
	}

}

func TestAngleBetween(t *testing.T) {

	t.Run("check perpindicular lines", func(t *testing.T) {
		v1 := NewVector(0, 10)
		v2 := NewVector(10, 0)

		theta := AngleBetween(v1, v2)

		if theta != math.Pi/2.0 {
			t.Errorf("should have been pi/2 %f", theta)
		}

	})

	t.Run("check diverging lines", func(t *testing.T) {
		v1 := NewVector(0, 10)
		v2 := NewVector(0, -10)

		theta := AngleBetween(v1, v2)

		if theta != math.Pi {
			t.Errorf("should have been pi/2 %f", theta)
		}

	})

	t.Run("check v1.AB(v2) == -v2.AB(v1)", func(t *testing.T) {
		v1 := NewVector(0, 10)
		v2 := NewVector(0, -10)

		theta1 := v1.AngleBetween(v2)
		theta2 := v2.AngleBetween(v1)

		if theta1 != -theta2 {
			t.Errorf("should have been opposite %f and %f", theta1, theta2)
		}

	})

}

func TestDistance(t *testing.T) {
	t.Run("check distance of 2 vectors", func(t *testing.T) {
		v1 := NewVector()
		v2 := NewVector(1, 0)

		d1 := v1.Dist(v2)
		d2 := v2.Dist(v1)

		if d1 != d2 {
			t.Error("Something has gone very wrong, v1.Dist(v2) != v2.Dist(v1)")
		}

		if d1 != 1.0 {
			t.Errorf("Should have equaled 1 -- %f", d1)
		}
	})
}

func TestNormalise(t *testing.T) {
	v := NewVector(3, 0, 0)

	v.Normalise()

	if v.X != 1.0 {
		t.Errorf("v not normalised %v", v)
	}

}
