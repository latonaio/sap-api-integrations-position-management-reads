package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Code                        string      `json:"code"`
			EffectiveStartDate          string      `json:"effectiveStartDate"`
			BusinessUnit                string      `json:"businessUnit"`
			LastModifiedDateTime        string      `json:"lastModifiedDateTime"`
			JobTitle                    string      `json:"jobTitle"`
			Criticality                 string      `json:"criticality"`
			CreatedDateTime             string      `json:"createdDateTime"`
			JobCode                     string      `json:"jobCode"`
			MdfSystemVersionID          string      `json:"mdfSystemVersionId"`
			Type                        string      `json:"type"`
			Incumbent                   string      `json:"incumbent"`
			Division                    string      `json:"division"`
			ExternalNameZhTW            string      `json:"externalName_zh_TW"`
			ExternalNameJaJP            string      `json:"externalName_ja_JP"`
			MdfSystemEntityID           string      `json:"mdfSystemEntityId"`
			PayRange                    string      `json:"payRange"`
			RegularTemporary            string      `json:"regularTemporary"`
			ExternalNamePtPT            string      `json:"externalName_pt_PT"`
			CostCenter                  string      `json:"costCenter"`
			StandardHours               string      `json:"standardHours"`
			LegacyPositionID            string      `json:"legacyPositionId"`
			ExternalNameZhCN            string      `json:"externalName_zh_CN"`
			ExternalNameLocalized       string      `json:"externalName_localized"`
			MdfSystemRecordStatus       string      `json:"mdfSystemRecordStatus"`
			Vacant                      bool        `json:"vacant"`
			EffectiveStatus             string      `json:"effectiveStatus"`
			ExternalNameFrFR            string      `json:"externalName_fr_FR"`
			TechnicalParameters         string      `json:"technicalParameters"`
			ExternalNameEnGB            string      `json:"externalName_en_GB"`
			EffectiveEndDate            string      `json:"effectiveEndDate"`
			PositionCriticality         string      `json:"positionCriticality"`
			ExternalNameNlNL            string      `json:"externalName_nl_NL"`
			PositionTitle               string      `json:"positionTitle"`
			Description                 string      `json:"description"`
			ExternalNameDefaultValue    string      `json:"externalName_defaultValue"`
			ExternalNameEsES            string      `json:"externalName_es_ES"`
			PositionControlled          bool        `json:"positionControlled"`
			ExternalNamePtBR            string      `json:"externalName_pt_BR"`
			PayGrade                    string      `json:"payGrade"`
			Company                     string      `json:"company"`
			Department                  string      `json:"department"`
			EmployeeClass               string      `json:"employeeClass"`
			MdfSystemObjectType         string      `json:"mdfSystemObjectType"`
			CreationSource              string      `json:"creationSource"`
			ChangeReason                string      `json:"changeReason"`
			TargetFTE                   string      `json:"targetFTE"`
			LastModifiedDate            string      `json:"lastModifiedDate"`
			ExternalNameRuRU            string      `json:"externalName_ru_RU"`
			LastModifiedBy              string      `json:"lastModifiedBy"`
			LastModifiedDateWithTZ      string      `json:"lastModifiedDateWithTZ"`
			JobLevel                    string      `json:"jobLevel"`
			TransactionSequence         string      `json:"transactionSequence"`
			ExternalNameDeDE            string      `json:"externalName_de_DE"`
			ExternalNameKoKR            string      `json:"externalName_ko_KR"`
			CreatedDate                 string      `json:"createdDate"`
			CreatedBy                   string      `json:"createdBy"`
			MdfSystemOptimisticLockUUID string      `json:"mdfSystemOptimisticLockUUID"`
			MdfSystemRecordID           string      `json:"mdfSystemRecordId"`
			Comment                     string      `json:"comment"`
			Location                    string      `json:"location"`
			MultipleIncumbentsAllowed   bool        `json:"multipleIncumbentsAllowed"`
			ExternalNameEnUS            string      `json:"externalName_en_US"`
			ExternalNameEnDEBUG         string      `json:"externalName_en_DEBUG"`
			CreationSourceNav           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"creationSourceNav"`
			RegularTemporaryNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"regularTemporaryNav"`
			CreatedByNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"createdByNav"`
			ExternalNameTranslationTextNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"externalNameTranslationTextNav"`
			IncumbentNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"incumbentNav"`
			PositionCriticalityNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"positionCriticalityNav"`
			ChangeReasonNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"changeReasonNav"`
			CompanyNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"companyNav"`
			SuccessorNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"successorNav"`
			EffectiveStatusNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"effectiveStatusNav"`
			DepartmentNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"departmentNav"`
			RightToReturn struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"rightToReturn"`
			BusinessUnitNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"businessUnitNav"`
			PositionMatrixRelationship struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"positionMatrixRelationship"`
			LocationNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"locationNav"`
			PayRangeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"payRangeNav"`
			JobLevelNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobLevelNav"`
			CostCenterNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"costCenterNav"`
			EmployeeClassNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"employeeClassNav"`
			JobCodeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"jobCodeNav"`
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
			DivisionNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"divisionNav"`
			ParentPosition struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"parentPosition"`
			WfRequestNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"wfRequestNav"`
			PayGradeNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"payGradeNav"`
		} `json:"results"`
	} `json:"d"`
}
