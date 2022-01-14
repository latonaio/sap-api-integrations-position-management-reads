package responses

type LocationNav struct {
	D struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ExternalCode         string      `json:"externalCode"`
		StartDate            string      `json:"startDate"`
		LastModifiedDateTime string      `json:"lastModifiedDateTime"`
		EndDate              string      `json:"endDate"`
		StandardHours        string      `json:"standardHours"`
		Timezone             string      `json:"timezone"`
		LastModifiedBy       string      `json:"lastModifiedBy"`
		CreatedDateTime      string      `json:"createdDateTime"`
		Description          string      `json:"description"`
		CreatedOn            string      `json:"createdOn"`
		LastModifiedOn       string      `json:"lastModifiedOn"`
		CreatedBy            string      `json:"createdBy"`
		Name                 string      `json:"name"`
		GeozoneFlx           string      `json:"geozoneFlx"`
		LocationGroup        string      `json:"locationGroup"`
		Status               string      `json:"status"`
		AddressNavDEFLT      struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"addressNavDEFLT"`
		DescriptionTranslationNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"descriptionTranslationNav"`
		LocationGroupNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"locationGroupNav"`
		NameTranslationNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"nameTranslationNav"`
		GeozoneFlxNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"geozoneFlxNav"`
		CompanyFlxNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"companyFlxNav"`
	} `json:"d"`
}
