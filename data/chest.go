package data

type Chest struct {
	Upcoming     []string `json:"upcoming"`
	SuperMagical int64    `json:superMagical`
	Magical      int64    `json:magical`
	Legendary    int64    `json:legendary`
	Epic         int64    `json:epic`
	Giant        int64    `json:giant`
}

func (chest Chest) GetChestByKey(key string) int64 {
	switch key {
	case "giant":
		return chest.Giant
	case "magical":
		return chest.Magical
	case "epic":
		return chest.Epic
	case "legendary":
		return chest.Legendary
	case "supermagical":
		return chest.SuperMagical
	}
	return 0
}
