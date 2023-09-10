package entity

type ReferencedIssue struct {
	ExternalId string `json:"externalId,omitempty"`
}

func (r *ReferencedIssue) Available() bool {
	return r.ReferencedByExternalId()
}

func (r *ReferencedIssue) GetReferencedEntityType() ReferencedEntityType {
	return ReferencedEntityTypeIssue
}

func (r *ReferencedIssue) ReferencedByExternalId() bool {
	return r.ExternalId != ""
}
