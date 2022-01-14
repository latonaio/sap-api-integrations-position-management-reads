package responses

type PayGradeNav struct {
	D struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ExternalCode              string      `json:"externalCode"`
		StartDate                 string      `json:"startDate"`
		LastModifiedDateTime      string      `json:"lastModifiedDateTime"`
		EndDate                   string      `json:"endDate"`
		LastModifiedBy            string      `json:"lastModifiedBy"`
		CreatedDateTime           string      `json:"createdDateTime"`
		Description               string      `json:"description"`
		CreatedOn                 string      `json:"createdOn"`
		LastModifiedOn            string      `json:"lastModifiedOn"`
		PaygradeLevel             string      `json:"paygradeLevel"`
		CustomString1             string      `json:"customString1"`
		CreatedBy                 string      `json:"createdBy"`
		Name                      string      `json:"name"`
		Status                    string      `json:"status"`
		DescriptionTranslationNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"descriptionTranslationNav"`
		NameTranslationNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"nameTranslationNav"`
	} `json:"d"`
}
