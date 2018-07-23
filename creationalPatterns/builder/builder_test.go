package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	t.Run("CarBuilder", func(t *testing.T) {
		carBuilder := &CarBuilder{}
		manufacturingComplex.SetBuilder(carBuilder)
		manufacturingComplex.Contstruct()

		car := carBuilder.Build()

		if car.Wheels != 4 {
			t.Errorf("Wheels on a car must be 4 and they ware %d\n", car.Wheels)
		}

		if car.Structure != "Car" {
			t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
		}

		if car.Seats != 5 {
			t.Errorf("Seats on a car must be 5 and they ware %d\n", car.Seats)
		}
	})

	t.Run("BikeBuilder", func(t *testing.T) {
		bikeBuilder := &BikeBuilder{}

		manufacturingComplex.SetBuilder(bikeBuilder)
		manufacturingComplex.Contstruct()

		motorbike := bikeBuilder.Build()

		if motorbike.Wheels != 2 {
			t.Errorf("Wheels on a motorbike must be 2 and they ware %d\n", motorbike.Wheels)
		}

		if motorbike.Structure != "Motorbike" {
			t.Errorf("Structure on a motorbike must be 'Motorbike' and was %s\n", motorbike.Structure)
		}

		if motorbike.Seats != 1 {
			t.Errorf("Seats on a motorbike must be 1 and they ware %d\n", motorbike.Seats)
		}

	})
}
