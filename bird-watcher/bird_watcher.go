package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirdCount := 0
	for v := range birdsPerDay {
		totalBirdCount += birdsPerDay[v]
	}
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	totalBirdCount := 0
	if len(birdsPerDay) >= week {

		weekCount := birdsPerDay[(week-1)*7 : (week-1)*7+7]
		for v := range weekCount {
			totalBirdCount += weekCount[v]
		}
		return totalBirdCount
	} else {
		for v := range birdsPerDay {
			totalBirdCount += birdsPerDay[v]
		}
		return totalBirdCount
	}
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for v := range birdsPerDay {
		if v%2 == 0 {
			birdsPerDay[v] = birdsPerDay[v] + 1
		}
	}
	return birdsPerDay
}
