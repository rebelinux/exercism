// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	r := Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
	return &r
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {

	value, exists := r.Address["street"]

	if r.Name != "" && exists && r.Age == 0 {
		return true
	}
	if r.Name != "" && value != "" && r.Age != 0 {
		return true
	}

	return false

}

// Delete deletes a resident's information.
func (r *Resident) Delete() {

	*r = Resident{
		Name:    "",
		Age:     0,
		Address: nil,
	}

}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {

	var c int

	for _, v := range residents {
		if v.HasRequiredInfo() {
			c += 1
		}
	}

	return c

}
