package intercom

import (
	"encoding/json"
	"fmt"
	"github.com/openline-ai/openline-customer-os/packages/runner/sync-customer-os-data/entity"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/utils"
	"strings"
	"time"
)

func MapUser(inputJson string) (string, error) {
	var input struct {
		ID    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	}

	if err := json.Unmarshal([]byte(inputJson), &input); err != nil {
		return "", err
	}

	output := entity.UserData{
		BaseData: entity.BaseData{
			ExternalId: input.ID,
		},
		Name:  input.Name,
		Email: input.Email,
	}
	if input.ID == "" {
		output.Skip = true
		output.SkipReason = "Missing external id"
	}

	return utils.ToJson(output)
}

func MapOrganization(inputJSON string) (string, error) {
	var input struct {
		Email     string `json:"email,omitempty"`
		ID        string `json:"id,omitempty"`
		CreatedAt int64  `json:"created_at,omitempty"`
	}

	err := json.Unmarshal([]byte(inputJSON), &input)
	if err != nil {
		return "", fmt.Errorf("failed to parse input JSON: %v", err)
	}
	if input.Email == "" {
		output := entity.OrganizationData{
			BaseData: entity.BaseData{
				Skip:       true,
				SkipReason: "Missing email",
			},
		}
		return utils.ToJson(output)
	}
	if input.ID == "" {
		output := entity.OrganizationData{
			BaseData: entity.BaseData{
				Skip:       true,
				SkipReason: "Missing id",
			},
		}
		return utils.ToJson(output)
	}

	output := entity.OrganizationData{
		BaseData: entity.BaseData{
			ExternalId: input.ID,
		},
		CreateByDomain:      true,
		ExternalSourceTable: utils.StringPtr("contacts"),
	}
	output.Domains = []string{extractDomain(input.Email)}
	if input.CreatedAt != 0 {
		output.CreatedAtStr = tsStrToRFC3339(input.CreatedAt)
	}

	return utils.ToJson(output)
}

func extractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func tsStrToRFC3339(timestamp int64) string {
	t := time.Unix(timestamp, 0).UTC()
	layout := "2006-01-02T15:04:05Z"
	return t.Format(layout)
}

//func MapContact(inputJSON string) (string, error) {
//	var input struct {
//		ID         int64  `json:"id,omitempty"`
//		Name       string `json:"name,omitempty"`
//		FirstName  string `json:"first_name,omitempty"`
//		LastName   string `json:"last_name,omitempty"`
//		Active     bool   `json:"active_flag,omitempty"`
//		AddTime    string `json:"add_time,omitempty"`
//		UpdateTime string `json:"update_time,omitempty"`
//		OrgId      int64  `json:"org_id,omitempty"`
//		OwnerId    int64  `json:"owner_id,omitempty"`
//		Emails     []struct {
//			Value   string `json:"value,omitempty"`
//			Primary bool   `json:"primary,omitempty"`
//		} `json:"email,omitempty"`
//		Phones []struct {
//			Value   string `json:"value,omitempty"`
//			Primary bool   `json:"primary,omitempty"`
//		} `json:"phone,omitempty"`
//	}
//
//	err := json.Unmarshal([]byte(inputJSON), &input)
//	if err != nil {
//		return "", fmt.Errorf("failed to parse input JSON: %v", err)
//	}
//
//	output := model.Output{
//		ExternalId: fmt.Sprintf("%d", input.ID),
//		Name:       input.Name,
//		FirstName:  input.FirstName,
//		LastName:   input.LastName,
//		CreatedAt:  input.AddTime,
//		UpdatedAt:  input.UpdateTime,
//	}
//	if input.ID == 0 {
//		output.Skip = true
//		output.SkipReason = "Missing external id"
//	}
//	if input.OrgId != 0 {
//		output.ExternalOrganizationId = fmt.Sprintf("%d", input.OrgId)
//	}
//	if input.OwnerId != 0 {
//		output.ExternalUserId = fmt.Sprintf("%d", input.OwnerId)
//	}
//
//	var primaryEmailFound = false
//	for _, email := range input.Emails {
//		if email.Value != "" {
//			if email.Primary && !primaryEmailFound {
//				output.Email = email.Value
//				primaryEmailFound = true
//			} else {
//				output.AdditionalEmails = append(output.AdditionalEmails, email.Value)
//			}
//		}
//	}
//	var primaryPhoneNumberFound = false
//	for _, phone := range input.Phones {
//		if phone.Value != "" {
//			if phone.Primary && !primaryPhoneNumberFound {
//				output.PhoneNumber = phone.Value
//				primaryPhoneNumberFound = true
//			} else {
//				output.AdditionalPhoneNumbers = append(output.AdditionalPhoneNumbers, phone.Value)
//			}
//		}
//	}
//
//	outputJSON, err := json.Marshal(output)
//	if err != nil {
//		return "", fmt.Errorf("failed to marshal output JSON: %v", err)
//	}
//
//	return string(outputJSON), nil
//}
//
//func MapNote(inputJSON string) (string, error) {
//	var input struct {
//		ID         int64  `json:"id,omitempty"`
//		Content    string `json:"content,omitempty"`
//		UserId     int64  `json:"user_id,omitempty"`
//		AddTime    string `json:"add_time,omitempty"`
//		UpdateTime string `json:"update_time,omitempty"`
//		PersonId   int64  `json:"person_id,omitempty"`
//		OrgId      int64  `json:"org_id,omitempty"`
//	}
//
//	err := json.Unmarshal([]byte(inputJSON), &input)
//	if err != nil {
//		return "", fmt.Errorf("failed to parse input JSON: %v", err)
//	}
//
//	output := model.Output{
//		ExternalId: fmt.Sprintf("%d", input.ID),
//		CreatedAt:  input.AddTime,
//		UpdatedAt:  input.UpdateTime,
//	}
//	if input.ID == 0 {
//		output.Skip = true
//		output.SkipReason = "Missing external id"
//	}
//	if input.UserId != 0 {
//		output.ExternalUserId = fmt.Sprintf("%d", input.UserId)
//	}
//	if input.PersonId != 0 {
//		output.ExternalContactsIds = append(output.ExternalContactsIds, fmt.Sprintf("%d", input.PersonId))
//	}
//	if input.OrgId != 0 {
//		output.ExternalOrganizationsIds = append(output.ExternalOrganizationsIds, fmt.Sprintf("%d", input.OrgId))
//	}
//	if strings.Contains(input.Content, "<") {
//		output.Content = input.Content
//		output.ContentType = "text/html"
//	} else {
//		output.Content = input.Content
//		output.ContentType = "text/plain"
//	}
//
//	outputJSON, err := json.Marshal(output)
//	if err != nil {
//		return "", fmt.Errorf("failed to marshal output JSON: %v", err)
//	}
//
//	return string(outputJSON), nil
//}
