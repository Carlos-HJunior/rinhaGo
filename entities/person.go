package entities

type Person struct {
	Nickname string   `json:"nickname"`
	Name     string   `json:"name"`
	Birth    string   `json:"birth"`
	Stack    []string `json:"stack"`
}
