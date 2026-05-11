package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(slice []string, time int) int {

	if time == 0 {
		return 2 * len(slice)
	} else {
		return time * len(slice)
	}

}

// TODO: define the 'Quantities()' function

func Quantities(slice []string) (int, float64) {

	var n int
	var s float64

	for _, v := range slice {
		if v == "noodles" {
			n += 1
		}
		if v == "sauce" {
			s += 1
		}
	}

	return n * 50, s * 0.2

}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(secret, ingredient []string) {

	for i, v := range ingredient {
		if v == "?" {
			ingredient[i] = secret[len(secret)-1]
		}
	}

}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(q []float64, s int) []float64 {

	var ns []float64

	for _, v := range q {

		if s < 2 && len(q) != 0 {

			p := v / 2

			ns = append(ns, p)

		} else {

			if s%2 == 0 {
				n := s / 2
				ns = append(ns, v*float64(n))

			}
			if s%2 == 1 {
				e := v / 2
				ns = append(ns, float64(e)+v)

			}
		}

	}

	return ns

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
