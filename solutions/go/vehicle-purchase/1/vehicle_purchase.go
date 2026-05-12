package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var isLicence bool
	switch kind {
	case "car":
		isLicence = true
	case "truck":
		isLicence = true
	default:
		isLicence = false
	}
	return isLicence
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	x := strings.Compare(option1, option2)
	var f string
	if x == 1 {
		f = option2
	} else {
		f = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", f)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var sp float64
	if age < 3 {
		sp = originalPrice * .8
	}
	if age > 3 {
		sp = originalPrice * .7
	}
	if age >= 10 {
		sp = originalPrice * .5
	}
	return sp
}
