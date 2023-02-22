package person

import (
	"encoding/json"
)

func NewPerson() Person {
	return Person{}
}
func (s *Person) ToJson() (any, error) {
	resultByte, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var result any
	err = json.Unmarshal(resultByte, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
