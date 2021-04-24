package mango

import "encoding/json"

type Mandate struct {
	ID            string `json:"Id"`
	BankAccountID string `json:"BankAccountId"`
	UserID        string `json:"UserId"`
	ReturnURL     string `json:"ReturnURL"`
	RedirectURL   string `json:"RedirectURL"`
	DocumentURL   string `json:"DocumentURL"`
	Culture       string `json:"Culture"`
	Scheme        string `json:"Scheme"`
	Status        string `json:"Status"`
	CreationDate  int64  `json:"CreationDate"`
	MandateType   string `json:"MandateType"`
	ExecutionType string `json:"ExecutionType"`
	Tag           string `json:"Tag"`
	ResultCode    string `json:"ResultCode"`
	ResultMessage string `json:"ResultMessage"`
}

type CreateMandate struct {
	BankAccountID string `json:"BankAccountId"`
	Culture       string `json:"Culture"`
	ReturnURL     string `json:"ReturnURL"`
}

func (m *MangoPay) CreateNewMandate(data CreateMandate) (Mandate, error) {
	var action = actionCreateMandate

	var jsonData JsonObject

	j, err := json.Marshal(data)
	if err != nil {
		return Mandate{}, err
	}

	err = json.Unmarshal(j, &jsonData)
	if err != nil {
		return Mandate{}, err
	}

	createdMandate, err := m.anyRequest(Mandate{}, action, jsonData)
	if err != nil {
		return Mandate{}, err
	}

	return createdMandate.(Mandate), nil
}

// ViewMandate view a created Mandate by ID
func (m *MangoPay) ViewMandate(mandateID string) (Mandate, error) {
	var action = actionViewMandate

	fetchedMandate, err := m.anyRequest(Mandate{}, action, JsonObject{"MandateID": mandateID})
	if err != nil {
		return Mandate{}, err
	}

	return fetchedMandate.(Mandate), nil
}
