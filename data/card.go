package data

type Card struct {
	Name  string `json:name`
	Level int64  `json:level`
	Icon  string `json:icon`
	Key   string `json:key`
}
