package pokeapi

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `josn:"weight"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	BaseExperience int `json:"base_experience"`
}
