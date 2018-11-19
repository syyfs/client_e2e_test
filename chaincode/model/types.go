package model

type SimpleTime string

func (s SimpleTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s + "\""), nil
}
