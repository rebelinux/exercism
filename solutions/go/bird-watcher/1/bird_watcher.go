package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for _, v := range birdsPerDay {
		sum += v
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	index := 7 * week

	var c int

	if week == 1 {

		c = TotalBirdCount(birdsPerDay[:index])

	} else {

		c = TotalBirdCount(birdsPerDay[:len(birdsPerDay)-index-1])
	}

	return c

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i, v := range birdsPerDay {
		if i == 0 {
			birdsPerDay[i] = v + 1
		}
		if i != 0 && i%2 == 0 {

			birdsPerDay[i] = v + 1

		}
		if i != 0 && i%2 == 1 {

			birdsPerDay[i] = v

		}

	}

	return birdsPerDay
}
