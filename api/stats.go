package api

var growthValues = map[int]map[int]int{
	1: {0: 6, 1: 8, 2: 9, 3: 11, 4: 13, 5: 14, 6: 16, 7: 18, 8: 19, 9: 21, 10: 23, 11: 24, 12: 26, 13: 28},
	2: {0: 7, 1: 8, 2: 10, 3: 12, 4: 14, 5: 15, 6: 17, 7: 19, 8: 21, 9: 23, 10: 25, 11: 26, 12: 28, 13: 30},
	3: {0: 7, 1: 9, 2: 11, 3: 13, 4: 15, 5: 17, 6: 19, 7: 21, 8: 23, 9: 25, 10: 27, 11: 29, 12: 31, 13: 33},
	4: {0: 8, 1: 10, 2: 12, 3: 15, 4: 16, 5: 18, 6: 20, 7: 22, 8: 24, 9: 26, 10: 28, 11: 31, 12: 33, 13: 35},
	5: {0: 9, 1: 10, 2: 13, 3: 15, 4: 17, 5: 19, 6: 22, 7: 24, 8: 26, 9: 28, 10: 30, 11: 33, 12: 35, 13: 37},
}

// Stats act as a single struct for each of the 5 stats.
// This can be use for base stats and growth points.
type Stats struct {
	HP  int `json:"hp"`
	Atk int `json:"atk"`
	Spd int `json:"spd"`
	Def int `json:"def"`
	Res int `json:"res"`
}

// StatRange represents a range of value for each stat.
// This can be used for base and max stat variations for banes and boons.
type StatRange struct {
	HP  []int `json:"hp"`
	Atk []int `json:"atk"`
	Spd []int `json:"spd"`
	Def []int `json:"def"`
	Res []int `json:"res"`
}

func getGrowthValue(rarity, growthPoint int) int {
	return growthValues[rarity][growthPoint]
}
