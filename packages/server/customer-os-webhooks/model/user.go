package model

import "strings"

type UserData struct {
	BaseData
	Name      string `json:"name,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	// TODO handle phone numbers
	//PhoneNumbers    []PhoneNumber `json:"phoneNumbers,omitempty"`
	ExternalOwnerId string `json:"externalOwnerId,omitempty"`
	ProfilePhotoUrl string `json:"profilePhotoUrl,omitempty"`
	Timezone        string `json:"timezone,omitempty"`
}

func (u *UserData) HasPhoneNumbers() bool {
	//return len(u.PhoneNumbers) > 0
	return false
}

func (u *UserData) HasEmail() bool {
	return u.Email != ""
}

func (u *UserData) Normalize() {
	u.SetTimes()
	u.ExternalSystem = strings.ToLower(strings.TrimSpace(u.ExternalSystem))

	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	//u.PhoneNumbers = GetNonEmptyPhoneNumbers(u.PhoneNumbers)
	//u.PhoneNumbers = RemoveDuplicatedPhoneNumbers(u.PhoneNumbers)
}
