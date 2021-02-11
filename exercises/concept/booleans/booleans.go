package booleans

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if prisoner is awake and the other 3 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	panic("Please implement the CanFreePrisoner() function")
}
