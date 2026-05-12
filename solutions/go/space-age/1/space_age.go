package space

type Planet string

func Age(seconds float64, planet Planet) float64 {

	YinS := float64(31556952)

	p := map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
		"Earth":   1.0,
	}

	_, ok := p[string(planet)]

	if ok {
		return (seconds / YinS) / p[string(planet)]
	}

	return -1

}
