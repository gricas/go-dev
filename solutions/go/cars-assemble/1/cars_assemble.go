package cars


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var x float64 = successRate / 100
    y := float64(productionRate)
    var cars_per_hour float64 = y * x
    return cars_per_hour
    
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var x float64 = successRate / 100
    y := float64(productionRate)
    var cars_per_hour float64 = y * x
    z := int(cars_per_hour)
    var cars_per_minute int = z / 60
    result := int(cars_per_minute)
    return result
    
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groups := carsCount / 10
    remaining := carsCount % 10

    cost := groups*95000 + remaining*10000
    return uint(cost)
    
    
	panic("CalculateCost not implemented")
}
