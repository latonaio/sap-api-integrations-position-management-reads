package responses

type DepartmentNav struct {
	D struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ExternalCode            string      `json:"externalCode"`
		StartDate               string      `json:"startDate"`
		Parent                  string      `json:"parent"`
		DescriptionKoKR         string      `json:"description_ko_KR"`
		LastModifiedDateTime    string      `json:"lastModifiedDateTime"`
		NameLocalized           string      `json:"name_localized"`
		EndDate                 string      `json:"endDate"`
		EntityUUID              string      `json:"entityUUID"`
		NameKoKR                string      `json:"name_ko_KR"`
		CreatedDateTime         string      `json:"createdDateTime"`
		DescriptionPtBR         string      `json:"description_pt_BR"`
		DescriptionEsES         string      `json:"description_es_ES"`
		NamePtBR                string      `json:"name_pt_BR"`
		DescriptionNlNL         string      `json:"description_nl_NL"`
		NameNlNL                string      `json:"name_nl_NL"`
		CostCenter              string      `json:"costCenter"`
		DescriptionDefaultValue string      `json:"description_defaultValue"`
		NameDeDE                string      `json:"name_de_DE"`
		NameZhTW                string      `json:"name_zh_TW"`
		Name                    string      `json:"name"`
		NameEsES                string      `json:"name_es_ES"`
		DescriptionEnUS         string      `json:"description_en_US"`
		DescriptionEnDEBUG      string      `json:"description_en_DEBUG"`
		DescriptionRuRU         string      `json:"description_ru_RU"`
		Status                  string      `json:"status"`
		NameRuRU                string      `json:"name_ru_RU"`
		DescriptionJaJP         string      `json:"description_ja_JP"`
		DescriptionFrFR         string      `json:"description_fr_FR"`
		NamePtPT                string      `json:"name_pt_PT"`
		Description             string      `json:"description"`
		DescriptionDeDE         string      `json:"description_de_DE"`
		NameFrFR                string      `json:"name_fr_FR"`
		NameEnDEBUG             string      `json:"name_en_DEBUG"`
		NameJaJP                string      `json:"name_ja_JP"`
		CreatedOn               string      `json:"createdOn"`
		HeadOfUnit              string      `json:"headOfUnit"`
		NameEnUS                string      `json:"name_en_US"`
		DescriptionZhTW         string      `json:"description_zh_TW"`
		NameZhCN                string      `json:"name_zh_CN"`
		NameDefaultValue        string      `json:"name_defaultValue"`
		DescriptionEnGB         string      `json:"description_en_GB"`
		LastModifiedBy          string      `json:"lastModifiedBy"`
		NameEnGB                string      `json:"name_en_GB"`
		LastModifiedOn          string      `json:"lastModifiedOn"`
		DescriptionZhCN         string      `json:"description_zh_CN"`
		CreatedBy               string      `json:"createdBy"`
		DescriptionLocalized    string      `json:"description_localized"`
		DescriptionPtPT         string      `json:"description_pt_PT"`
		CustToLegalEntity       struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"cust_toLegalEntity"`
		NameTranslationTextNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"nameTranslationTextNav"`
		CreatedByNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"createdByNav"`
		HeadOfUnitNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"headOfUnitNav"`
		DivisionFlxNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"divisionFlxNav"`
		DescriptionTranslationTextNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"descriptionTranslationTextNav"`
		StatusNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"statusNav"`
		DescriptionTranslationNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"descriptionTranslationNav"`
		CustToDivision struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"cust_toDivision"`
		CostCenterNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"costCenterNav"`
		LastModifiedByNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"lastModifiedByNav"`
		NameTranslationNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"nameTranslationNav"`
		ToDepartmentApprenticeDetail struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"toDepartmentApprenticeDetail"`
		ParentNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"parentNav"`
	} `json:"d"`
}
