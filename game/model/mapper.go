package model

var strengthMapper = map[int]int{
	1: 12,
	2: 8,
	3: 10,
	4: 7,
	5: 9,
	6: 11,
}

var agilityMapper = map[int]int{
	1: 12,
	2: 8,
	3: 10,
	4: 7,
	5: 9,
	6: 11,
}

func GetStrength(dice int) int {
	return strengthMapper[dice]
}

func GetAgility(dice int) int {
	return agilityMapper[dice]
}
