package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	var output bool
	switch knightIsAwake {
	case true:
		output = false
	case false:
		output = true
	}
	return output
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake

}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var output bool
	if prisonerIsAwake {
		output = true
	}
	if prisonerIsAwake && archerIsAwake {
		output = false
	}

	return output

}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var output bool
	if petDogIsPresent && !archerIsAwake {
		output = true
	}
	if !petDogIsPresent && prisonerIsAwake && (!archerIsAwake && !knightIsAwake) {
		output = true
	}

	return output
}
