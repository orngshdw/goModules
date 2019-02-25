package utility

// List type allows methods to be called on maps
type List map[string]int

// get method on List which returns a default value if key doesn't exist
// inspired by Python's .get(key, default)
func (mapObj List) get(key string, defaultVal int) int {
	value, err := mapObj[key]
	if err != false {
		return defaultVal
	}
	return value
}
