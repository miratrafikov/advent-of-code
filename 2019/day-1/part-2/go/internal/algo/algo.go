package algo

func CalculateTotalFuelNeeded(moduleMasses []int) int {
	totalFuelMass := 0
	for _, moduleMass := range moduleMasses {
		totalFuelMass += calculateFuelNeededForModule(moduleMass)
	}
	return totalFuelMass
}

func calculateFuelNeededForModule(moduleMass int) int {
	totalFuelMass := 0
	fuelApproximation := moduleMass/3 - 2
	for fuelApproximation > 0 {
		totalFuelMass += fuelApproximation
		fuelApproximation = fuelApproximation/3 - 2
	}
	return totalFuelMass
}
