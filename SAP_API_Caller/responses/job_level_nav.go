package responses

type JobLevelNav struct {
	D struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		ExternalCode                 string      `json:"externalCode"`
		PickListV2EffectiveStartDate string      `json:"PickListV2_effectiveStartDate"`
		PickListV2ID                 string      `json:"PickListV2_id"`
		LabelNlNL                    string      `json:"label_nl_NL"`
		LastModifiedDateTime         string      `json:"lastModifiedDateTime"`
		MdfSystemEffectiveEndDate    string      `json:"mdfSystemEffectiveEndDate"`
		CreatedDateTime              string      `json:"createdDateTime"`
		MdfSystemVersionID           string      `json:"mdfSystemVersionId"`
		LabelEsES                    string      `json:"label_es_ES"`
		MaxVal                       string      `json:"maxVal"`
		MinVal                       string      `json:"minVal"`
		LabelPtBR                    string      `json:"label_pt_BR"`
		MdfSystemEntityID            string      `json:"mdfSystemEntityId"`
		LabelEnDEBUG                 string      `json:"label_en_DEBUG"`
		NonUniqueExternalCode        string      `json:"nonUniqueExternalCode"`
		LegacyStatus                 string      `json:"legacyStatus"`
		LabelRuRU                    string      `json:"label_ru_RU"`
		MdfSystemRecordStatus        string      `json:"mdfSystemRecordStatus"`
		LabelDefaultValue            string      `json:"label_defaultValue"`
		LabelEnUS                    string      `json:"label_en_US"`
		LabelDeDE                    string      `json:"label_de_DE"`
		LabelLocalized               string      `json:"label_localized"`
		LabelKoKR                    string      `json:"label_ko_KR"`
		OptionID                     string      `json:"optionId"`
		Status                       string      `json:"status"`
		OptValue                     string      `json:"optValue"`
		LValue                       string      `json:"lValue"`
		MdfSystemEffectiveStartDate  string      `json:"mdfSystemEffectiveStartDate"`
		LabelZhTW                    string      `json:"label_zh_TW"`
		LabelJaJP                    string      `json:"label_ja_JP"`
		ExternalStandardizedCode     string      `json:"externalStandardizedCode"`
		MdfSystemObjectType          string      `json:"mdfSystemObjectType"`
		LabelPtPT                    string      `json:"label_pt_PT"`
		LastModifiedDate             string      `json:"lastModifiedDate"`
		LastModifiedBy               string      `json:"lastModifiedBy"`
		LastModifiedDateWithTZ       string      `json:"lastModifiedDateWithTZ"`
		LabelZhCN                    string      `json:"label_zh_CN"`
		RValue                       string      `json:"rValue"`
		MdfSystemTransactionSequence string      `json:"mdfSystemTransactionSequence"`
		CreatedDate                  string      `json:"createdDate"`
		CreatedBy                    string      `json:"createdBy"`
		ParentPickListValue          string      `json:"parentPickListValue"`
		MdfSystemRecordID            string      `json:"mdfSystemRecordId"`
		LabelFrFR                    string      `json:"label_fr_FR"`
		LabelEnGB                    string      `json:"label_en_GB"`
		CreatedByNav                 struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"createdByNav"`
		StatusNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"statusNav"`
		LabelTranslationTextNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"labelTranslationTextNav"`
		LastModifiedByNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"lastModifiedByNav"`
		MdfSystemRecordStatusNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"mdfSystemRecordStatusNav"`
		ParentPickListValueNav struct {
			Deferred struct {
				URI string `json:"uri"`
			} `json:"__deferred"`
		} `json:"parentPickListValueNav"`
	} `json:"d"`
}
