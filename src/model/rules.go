package model

type Rule struct {
	MinVersion int `json:"min_version"`
	MaxVersion int `json:"max_version"`
	MinUserDID int `json:"min_user_did"`
	MaxUserDID int `json:"max_user_did"`

	GreyLink string `json:"grey_link"`
}

func GetRules() *[]Rule {
	rules := []Rule{}

	rules = append(rules, Rule{
		MinUserDID: 10,
		MaxUserDID: 20,
		MinVersion: 10,
		MaxVersion: 20,
		GreyLink:   "123123",
	})
	return &rules
}
