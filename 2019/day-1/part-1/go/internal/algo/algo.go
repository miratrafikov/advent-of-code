package algo

func CalculateTotalFuelNeeded(moduleMasses []int) int {
	totalFuelMass := 0
	for _, moduleMass := range moduleMasses {
		totalFuelMass += calculateFuelNeededForModule(moduleMass)
	}
	return totalFuelMass
}

func calculateFuelNeededForModule(moduleMass int) int {
	return moduleMass/3 - 2
}
