package speed

// TODO: define the 'Car' type struct

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
	track        Track
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	c := Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
		distance:     0,
	}
	return c
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	t := Track{
		distance: distance,
	}
	return t
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	var c Car
	if car.batteryDrain > car.battery {
		c = Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery,
			distance:     car.distance,
			track:        NewTrack(0),
		}

	} else {
		c = Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery - car.batteryDrain,
			distance:     car.speed + car.distance,
			track:        NewTrack(car.track.distance),
		}
	}
	return c
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	var f bool
	d := car.speed * (car.battery / car.batteryDrain)

	if car.battery == track.distance || d > track.distance {
		f = true
	} else {
		f = false
	}
	if d == track.distance {
		f = true
	}

	return f
}
