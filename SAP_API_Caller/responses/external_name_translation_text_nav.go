package responses

type ExternalNameTranslationTextNav struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Locale string `json:"locale"`
			Value  string `json:"value"`
		} `json:"results"`
	} `json:"d"`
}
