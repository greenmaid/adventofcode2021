package day1

func Step1_countIncreaseValues(values []int) int {
	count := 0
	for index := range values {
		if index > 0 {
			if values[index-1] < values[index] {
				count += 1
			}
		}
	}
	return count
}

func Step2_countIncreaseValuesByGroup(values []int) int {

	var groupedValues []int
	for index := range values {
		if index < len(values)-2 {
			groupedValues = append(groupedValues, values[index]+values[index+1]+values[index+2])
		}
	}
	return Step1_countIncreaseValues(groupedValues)
}
