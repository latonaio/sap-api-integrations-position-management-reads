package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-position-management-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	Code:                        data.Code,
	EffectiveStartDate:          data.EffectiveStartDate,
	BusinessUnit:                data.BusinessUnit,
	LastModifiedDateTime:        data.LastModifiedDateTime,
	JobTitle:                    data.JobTitle,
	Criticality:                 data.Criticality,
	CreatedDateTime:             data.CreatedDateTime,
	JobCode:                     data.JobCode,
	MdfSystemVersionID:          data.MdfSystemVersionID,
	Type:                        data.Type,
	Incumbent:                   data.Incumbent,
	Division:                    data.Division,
	ExternalNameZhTW:            data.ExternalNameZhTW,
	ExternalNameJaJP:            data.ExternalNameJaJP,
	MdfSystemEntityID:           data.MdfSystemEntityID,
	PayRange:                    data.PayRange,
	RegularTemporary:            data.RegularTemporary,
	ExternalNamePtPT:            data.ExternalNamePtPT,
	CostCenter:                  data.CostCenter,
	StandardHours:               data.StandardHours,
	LegacyPositionID:            data.LegacyPositionID,
	ExternalNameZhCN:            data.ExternalNameZhCN,
	ExternalNameLocalized:       data.ExternalNameLocalized,
	MdfSystemRecordStatus:       data.MdfSystemRecordStatus,
	Vacant:                      data.Vacant,
	EffectiveStatus:             data.EffectiveStatus,
	ExternalNameFrFR:            data.ExternalNameFrFR,
	TechnicalParameters:         data.TechnicalParameters,
	ExternalNameEnGB:            data.ExternalNameEnGB,
	EffectiveEndDate:            data.EffectiveEndDate,
	PositionCriticality:         data.PositionCriticality,
	ExternalNameNlNL:            data.ExternalNameNlNL,
	PositionTitle:               data.PositionTitle,
	Description:                 data.Description,
	ExternalNameDefaultValue:    data.ExternalNameDefaultValue,
	ExternalNameEsES:            data.ExternalNameEsES,
	PositionControlled:          data.PositionControlled,
	ExternalNamePtBR:            data.ExternalNamePtBR,
	PayGrade:                    data.PayGrade,
	Company:                     data.Company,
	Department:                  data.Department,
	EmployeeClass:               data.EmployeeClass,
	MdfSystemObjectType:         data.MdfSystemObjectType,
	CreationSource:              data.CreationSource,
	ChangeReason:                data.ChangeReason,
	TargetFTE:                   data.TargetFTE,
	LastModifiedDate:            data.LastModifiedDate,
	ExternalNameRuRU:            data.ExternalNameRuRU,
	LastModifiedBy:              data.LastModifiedBy,
	LastModifiedDateWithTZ:      data.LastModifiedDateWithTZ,
	JobLevel:                    data.JobLevel,
	TransactionSequence:         data.TransactionSequence,
	ExternalNameDeDE:            data.ExternalNameDeDE,
	ExternalNameKoKR:            data.ExternalNameKoKR,
	CreatedDate:                 data.CreatedDate,
	CreatedBy:                   data.CreatedBy,
	MdfSystemOptimisticLockUUID: data.MdfSystemOptimisticLockUUID,
	MdfSystemRecordID:           data.MdfSystemRecordID,
	Comment:                     data.Comment,
	Location:                    data.Location,
	MultipleIncumbentsAllowed:   data.MultipleIncumbentsAllowed,
	ExternalNameEnUS:            data.ExternalNameEnUS,
	ExternalNameEnDEBUG:         data.ExternalNameEnDEBUG,
	RegularTemporaryNav:         data.RegularTemporaryNav.Deferred.URI,
	ExternalNameTranslationTextNav: data.ExternalNameTranslationTextNav.Deferred.URI,
	CompanyNav:                  data.CompanyNav.Deferred.URI,
	EffectiveStatusNav:          data.EffectiveStatusNav.Deferred.URI,
	DepartmentNav:               data.DepartmentNav.Deferred.URI,
	BusinessUnitNav:             data.BusinessUnitNav.Deferred.URI,
	LocationNav:                 data.LocationNav.Deferred.URI,
	CostCenterNav:               data.CostCenterNav.Deferred.URI,
	JobCodeNav:                  data.JobCodeNav.Deferred.URI,
	DivisionNav:                 data.DivisionNav.Deferred.URI,
	JobLevelNav:                 data.JobLevelNav.Deferred.URI,
	ParentPosition:              data.ParentPosition.Deferred.URI,
	PayGradeNav:                 data.PayGradeNav.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToRegularTemporaryNav(raw []byte, l *logger.Logger) (*RegularTemporaryNav, error) {
	pm := &responses.RegularTemporaryNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to RegularTemporaryNav. unmarshal error: %w", err)
	}
	
	return &RegularTemporaryNav{
	ExternalCode:                 pm.D.ExternalCode,
	PickListV2EffectiveStartDate: pm.D.PickListV2EffectiveStartDate,
	PickListV2ID:                 pm.D.PickListV2ID,
	LabelNlNL:                    pm.D.LabelNlNL,
	LastModifiedDateTime:         pm.D.LastModifiedDateTime,
	MdfSystemEffectiveEndDate:    pm.D.MdfSystemEffectiveEndDate,
	CreatedDateTime:              pm.D.CreatedDateTime,
	MdfSystemVersionID:           pm.D.MdfSystemVersionID,
	LabelEsES:                    pm.D.LabelEsES,
	MaxVal:                       pm.D.MaxVal,
	MinVal:                       pm.D.MinVal,
	LabelPtBR:                    pm.D.LabelPtBR,
	MdfSystemEntityID:            pm.D.MdfSystemEntityID,
	LabelEnDEBUG:                 pm.D.LabelEnDEBUG,
	NonUniqueExternalCode:        pm.D.NonUniqueExternalCode,
	LegacyStatus:                 pm.D.LegacyStatus,
	LabelRuRU:                    pm.D.LabelRuRU,
	MdfSystemRecordStatus:        pm.D.MdfSystemRecordStatus,
	LabelDefaultValue:            pm.D.LabelDefaultValue,
	LabelEnUS:                    pm.D.LabelEnUS,
	LabelDeDE:                    pm.D.LabelDeDE,
	LabelLocalized:               pm.D.LabelLocalized,
	LabelKoKR:                    pm.D.LabelKoKR,
	OptionID:                     pm.D.OptionID,
	Status:                       pm.D.Status,
	OptValue:                     pm.D.OptValue,
	LValue:                       pm.D.LValue,
	MdfSystemEffectiveStartDate:  pm.D.MdfSystemEffectiveStartDate,
	LabelZhTW:                    pm.D.LabelZhTW,
	LabelJaJP:                    pm.D.LabelJaJP,
	ExternalStandardizedCode:     pm.D.ExternalStandardizedCode,
	MdfSystemObjectType:          pm.D.MdfSystemObjectType,
	LabelPtPT:                    pm.D.LabelPtPT,
	LastModifiedDate:             pm.D.LastModifiedDate,
	LastModifiedBy:               pm.D.LastModifiedBy,
	LastModifiedDateWithTZ:       pm.D.LastModifiedDateWithTZ,
	LabelZhCN:                    pm.D.LabelZhCN,
	RValue:                       pm.D.RValue,
	MdfSystemTransactionSequence: pm.D.MdfSystemTransactionSequence,
	CreatedDate:                  pm.D.CreatedDate,
	CreatedBy:                    pm.D.CreatedBy,
	ParentPickListValue:          pm.D.ParentPickListValue,
	MdfSystemRecordID:            pm.D.MdfSystemRecordID,
	LabelFrFR:                    pm.D.LabelFrFR,
	LabelEnGB:                    pm.D.LabelEnGB,
	}, nil
}

func ConvertToExternalNameTranslationTextNav(raw []byte, l *logger.Logger) ([]ExternalNameTranslationTextNav, error) {
	pm := &responses.ExternalNameTranslationTextNav{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ExternalNameTranslationTextNav. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	externalNameTranslationTextNav := make([]ExternalNameTranslationTextNav, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		externalNameTranslationTextNav = append(externalNameTranslationTextNav, ExternalNameTranslationTextNav{
	Locale:                      data.Locale,
	Value:                       data.Value,
		})
	}

	return externalNameTranslationTextNav, nil
}

func ConvertToCompanyNav(raw []byte, l *logger.Logger) (*CompanyNav, error) {
	pm := &responses.CompanyNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CompanyNav. unmarshal error: %w", err)
	}
	
	return &CompanyNav{
	ExternalCode:            pm.D.ExternalCode,
	StartDate:               pm.D.StartDate,
	Country:                 pm.D.Country,
	DescriptionKoKR:         pm.D.DescriptionKoKR,
	LastModifiedDateTime:    pm.D.LastModifiedDateTime,
	NameLocalized:           pm.D.NameLocalized,
	EndDate:                 pm.D.EndDate,
	NameKoKR:                pm.D.NameKoKR,
	CreatedDateTime:         pm.D.CreatedDateTime,
	DescriptionPtBR:         pm.D.DescriptionPtBR,
	DescriptionEsES:         pm.D.DescriptionEsES,
	NamePtBR:                pm.D.NamePtBR,
	DescriptionNlNL:         pm.D.DescriptionNlNL,
	NameNlNL:                pm.D.NameNlNL,
	StandardHours:           pm.D.StandardHours,
	DescriptionDefaultValue: pm.D.DescriptionDefaultValue,
	NameDeDE:                pm.D.NameDeDE,
	NameZhTW:                pm.D.NameZhTW,
	Name:                    pm.D.Name,
	NameEsES:                pm.D.NameEsES,
	DescriptionEnUS:         pm.D.DescriptionEnUS,
	DescriptionEnDEBUG:      pm.D.DescriptionEnDEBUG,
	DescriptionRuRU:         pm.D.DescriptionRuRU,
	Status:                  pm.D.Status,
	NameRuRU:                pm.D.NameRuRU,
	DescriptionJaJP:         pm.D.DescriptionJaJP,
	DescriptionFrFR:         pm.D.DescriptionFrFR,
	NamePtPT:                pm.D.NamePtPT,
	Description:             pm.D.Description,
	DescriptionDeDE:         pm.D.DescriptionDeDE,
	NameFrFR:                pm.D.NameFrFR,
	NameEnDEBUG:             pm.D.NameEnDEBUG,
	NameJaJP:                pm.D.NameJaJP,
	CreatedOn:               pm.D.CreatedOn,
	NameEnUS:                pm.D.NameEnUS,
	Currency:                pm.D.Currency,
	DescriptionZhTW:         pm.D.DescriptionZhTW,
	NameZhCN:                pm.D.NameZhCN,
	NameDefaultValue:        pm.D.NameDefaultValue,
	DescriptionEnGB:         pm.D.DescriptionEnGB,
	LastModifiedBy:          pm.D.LastModifiedBy,
	DefaultPayGroup:         pm.D.DefaultPayGroup,
	NameEnGB:                pm.D.NameEnGB,
	LastModifiedOn:          pm.D.LastModifiedOn,
	DescriptionZhCN:         pm.D.DescriptionZhCN,
	CreatedBy:               pm.D.CreatedBy,
	DescriptionLocalized:    pm.D.DescriptionLocalized,
	DescriptionPtPT:         pm.D.DescriptionPtPT,
	DefaultLocation:         pm.D.DefaultLocation,
	}, nil
}

func ConvertToEffectiveStatusNav(raw []byte, l *logger.Logger) (*EffectiveStatusNav, error) {
	pm := &responses.EffectiveStatusNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to EffectiveStatusNav. unmarshal error: %w", err)
	}
	
	return &EffectiveStatusNav{
	Value:     pm.D.Value,
	Key:       pm.D.Key,
	EnDEBUG:   pm.D.EnDEBUG,
	RuRU:      pm.D.RuRU,
	Localized: pm.D.Localized,
	PtBR:      pm.D.PtBR,
	FrFR:      pm.D.FrFR,
	JaJP:      pm.D.JaJP,
	DeDE:      pm.D.DeDE,
	EnGB:      pm.D.EnGB,
	ZhTW:      pm.D.ZhTW,
	KoKR:      pm.D.KoKR,
	EnUS:      pm.D.EnUS,
	EsES:      pm.D.EsES,
	ZhCN:      pm.D.ZhCN,
	NlNL:      pm.D.NlNL,
	PtPT:      pm.D.PtPT,
	}, nil
}

func ConvertToDepartmentNav(raw []byte, l *logger.Logger) (*DepartmentNav, error) {
	pm := &responses.DepartmentNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to DepartmentNav. unmarshal error: %w", err)
	}
	
	return &DepartmentNav{
	ExternalCode:            pm.D.ExternalCode,
	StartDate:               pm.D.StartDate,
	Parent:                  pm.D.Parent,
	DescriptionKoKR:         pm.D.DescriptionKoKR,
	LastModifiedDateTime:    pm.D.LastModifiedDateTime,
	NameLocalized:           pm.D.NameLocalized,
	EndDate:                 pm.D.EndDate,
	EntityUUID:              pm.D.EntityUUID,
	NameKoKR:                pm.D.NameKoKR,
	CreatedDateTime:         pm.D.CreatedDateTime,
	DescriptionPtBR:         pm.D.DescriptionPtBR,
	DescriptionEsES:         pm.D.DescriptionEsES,
	NamePtBR:                pm.D.NamePtBR,
	DescriptionNlNL:         pm.D.DescriptionNlNL,
	NameNlNL:                pm.D.NameNlNL,
	CostCenter:              pm.D.CostCenter,
	DescriptionDefaultValue: pm.D.DescriptionDefaultValue,
	NameDeDE:                pm.D.NameDeDE,
	NameZhTW:                pm.D.NameZhTW,
	Name:                    pm.D.Name,
	NameEsES:                pm.D.NameEsES,
	DescriptionEnUS:         pm.D.DescriptionEnUS,
	DescriptionEnDEBUG:      pm.D.DescriptionEnDEBUG,
	DescriptionRuRU:         pm.D.DescriptionRuRU,
	Status:                  pm.D.Status,
	NameRuRU:                pm.D.NameRuRU,
	DescriptionJaJP:         pm.D.DescriptionJaJP,
	DescriptionFrFR:         pm.D.DescriptionFrFR,
	NamePtPT:                pm.D.NamePtPT,
	Description:             pm.D.Description,
	DescriptionDeDE:         pm.D.DescriptionDeDE,
	NameFrFR:                pm.D.NameFrFR,
	NameEnDEBUG:             pm.D.NameEnDEBUG,
	NameJaJP:                pm.D.NameJaJP,
	CreatedOn:               pm.D.CreatedOn,
	HeadOfUnit:              pm.D.HeadOfUnit,
	NameEnUS:                pm.D.NameEnUS,
	DescriptionZhTW:         pm.D.DescriptionZhTW,
	NameZhCN:                pm.D.NameZhCN,
	NameDefaultValue:        pm.D.NameDefaultValue,
	DescriptionEnGB:         pm.D.DescriptionEnGB,
	LastModifiedBy:          pm.D.LastModifiedBy,
	NameEnGB:                pm.D.NameEnGB,
	LastModifiedOn:          pm.D.LastModifiedOn,
	DescriptionZhCN:         pm.D.DescriptionZhCN,
	CreatedBy:               pm.D.CreatedBy,
	DescriptionLocalized:    pm.D.DescriptionLocalized,
	DescriptionPtPT:         pm.D.DescriptionPtPT,
	}, nil
}

func ConvertToBusinessUnitNav(raw []byte, l *logger.Logger) (*BusinessUnitNav, error) {
	pm := &responses.BusinessUnitNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to BusinessUnitNav. unmarshal error: %w", err)
	}
	
	return &BusinessUnitNav{
	ExternalCode:            pm.D.ExternalCode,
	StartDate:               pm.D.StartDate,
	DescriptionKoKR:         pm.D.DescriptionKoKR,
	LastModifiedDateTime:    pm.D.LastModifiedDateTime,
	NameLocalized:           pm.D.NameLocalized,
	EndDate:                 pm.D.EndDate,
	EntityUUID:              pm.D.EntityUUID,
	NameKoKR:                pm.D.NameKoKR,
	CreatedDateTime:         pm.D.CreatedDateTime,
	DescriptionPtBR:         pm.D.DescriptionPtBR,
	DescriptionEsES:         pm.D.DescriptionEsES,
	NamePtBR:                pm.D.NamePtBR,
	DescriptionNlNL:         pm.D.DescriptionNlNL,
	NameNlNL:                pm.D.NameNlNL,
	DescriptionDefaultValue: pm.D.DescriptionDefaultValue,
	NameDeDE:                pm.D.NameDeDE,
	NameZhTW:                pm.D.NameZhTW,
	Name:                    pm.D.Name,
	NameEsES:                pm.D.NameEsES,
	DescriptionEnUS:         pm.D.DescriptionEnUS,
	DescriptionEnDEBUG:      pm.D.DescriptionEnDEBUG,
	DescriptionRuRU:         pm.D.DescriptionRuRU,
	Status:                  pm.D.Status,
	NameRuRU:                pm.D.NameRuRU,
	DescriptionJaJP:         pm.D.DescriptionJaJP,
	DescriptionFrFR:         pm.D.DescriptionFrFR,
	NamePtPT:                pm.D.NamePtPT,
	Description:             pm.D.Description,
	DescriptionDeDE:         pm.D.DescriptionDeDE,
	NameFrFR:                pm.D.NameFrFR,
	NameEnDEBUG:             pm.D.NameEnDEBUG,
	NameJaJP:                pm.D.NameJaJP,
	CreatedOn:               pm.D.CreatedOn,
	HeadOfUnit:              pm.D.HeadOfUnit,
	NameEnUS:                pm.D.NameEnUS,
	DescriptionZhTW:         pm.D.DescriptionZhTW,
	NameZhCN:                pm.D.NameZhCN,
	NameDefaultValue:        pm.D.NameDefaultValue,
	DescriptionEnGB:         pm.D.DescriptionEnGB,
	LastModifiedBy:          pm.D.LastModifiedBy,
	NameEnGB:                pm.D.NameEnGB,
	LastModifiedOn:          pm.D.LastModifiedOn,
	DescriptionZhCN:         pm.D.DescriptionZhCN,
	CreatedBy:               pm.D.CreatedBy,
	DescriptionLocalized:    pm.D.DescriptionLocalized,
	DescriptionPtPT:         pm.D.DescriptionPtPT,
	}, nil
}

func ConvertToLocationNav(raw []byte, l *logger.Logger) (*LocationNav, error) {
	pm := &responses.LocationNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to LocationNav. unmarshal error: %w", err)
	}
	
	return &LocationNav{
	ExternalCode:         pm.D.ExternalCode,
	StartDate:            pm.D.StartDate,
	LastModifiedDateTime: pm.D.LastModifiedDateTime,
	EndDate:              pm.D.EndDate,
	StandardHours:        pm.D.StandardHours,
	Timezone:             pm.D.Timezone,
	LastModifiedBy:       pm.D.LastModifiedBy,
	CreatedDateTime:      pm.D.CreatedDateTime,
	Description:          pm.D.Description,
	CreatedOn:            pm.D.CreatedOn,
	LastModifiedOn:       pm.D.LastModifiedOn,
	CreatedBy:            pm.D.CreatedBy,
	Name:                 pm.D.Name,
	GeozoneFlx:           pm.D.GeozoneFlx,
	LocationGroup:        pm.D.LocationGroup,
	Status:               pm.D.Status,
	}, nil
}

func ConvertToJobLevelNav(raw []byte, l *logger.Logger) (*JobLevelNav, error) {
	pm := &responses.JobLevelNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to JobLevelNav. unmarshal error: %w", err)
	}
	
	return &JobLevelNav{
	ExternalCode:                 pm.D.ExternalCode,
	PickListV2EffectiveStartDate: pm.D.PickListV2EffectiveStartDate,
	PickListV2ID:                 pm.D.PickListV2ID,
	LabelNlNL:                    pm.D.LabelNlNL,
	LastModifiedDateTime:         pm.D.LastModifiedDateTime,
	MdfSystemEffectiveEndDate:    pm.D.MdfSystemEffectiveEndDate,
	CreatedDateTime:              pm.D.CreatedDateTime,
	MdfSystemVersionID:           pm.D.MdfSystemVersionID,
	LabelEsES:                    pm.D.LabelEsES,
	MaxVal:                       pm.D.MaxVal,
	MinVal:                       pm.D.MinVal,
	LabelPtBR:                    pm.D.LabelPtBR,
	MdfSystemEntityID:            pm.D.MdfSystemEntityID,
	LabelEnDEBUG:                 pm.D.LabelEnDEBUG,
	NonUniqueExternalCode:        pm.D.NonUniqueExternalCode,
	LegacyStatus:                 pm.D.LegacyStatus,
	LabelRuRU:                    pm.D.LabelRuRU,
	MdfSystemRecordStatus:        pm.D.MdfSystemRecordStatus,
	LabelDefaultValue:            pm.D.LabelDefaultValue,
	LabelEnUS:                    pm.D.LabelEnUS,
	LabelDeDE:                    pm.D.LabelDeDE,
	LabelLocalized:               pm.D.LabelLocalized,
	LabelKoKR:                    pm.D.LabelKoKR,
	OptionID:                     pm.D.OptionID,
	Status:                       pm.D.Status,
	OptValue:                     pm.D.OptValue,
	LValue:                       pm.D.LValue,
	MdfSystemEffectiveStartDate:  pm.D.MdfSystemEffectiveStartDate,
	LabelZhTW:                    pm.D.LabelZhTW,
	LabelJaJP:                    pm.D.LabelJaJP,
	ExternalStandardizedCode:     pm.D.ExternalStandardizedCode,
	MdfSystemObjectType:          pm.D.MdfSystemObjectType,
	LabelPtPT:                    pm.D.LabelPtPT,
	LastModifiedDate:             pm.D.LastModifiedDate,
	LastModifiedBy:               pm.D.LastModifiedBy,
	LastModifiedDateWithTZ:       pm.D.LastModifiedDateWithTZ,
	LabelZhCN:                    pm.D.LabelZhCN,
	RValue:                       pm.D.RValue,
	MdfSystemTransactionSequence: pm.D.MdfSystemTransactionSequence,
	CreatedDate:                  pm.D.CreatedDate,
	CreatedBy:                    pm.D.CreatedBy,
	ParentPickListValue:          pm.D.ParentPickListValue,
	MdfSystemRecordID:            pm.D.MdfSystemRecordID,
	LabelFrFR:                    pm.D.LabelFrFR,
	LabelEnGB:                    pm.D.LabelEnGB,
	}, nil
}

func ConvertToCostCenterNav(raw []byte, l *logger.Logger) (*CostCenterNav, error) {
	pm := &responses.CostCenterNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CostCenterNav. unmarshal error: %w", err)
	}
	
	return &CostCenterNav{
	ExternalCode:               pm.D.ExternalCode,
	StartDate:                  pm.D.StartDate,
	Parent:                     pm.D.Parent,
	DescriptionKoKR:            pm.D.DescriptionKoKR,
	LastModifiedDateTime:       pm.D.LastModifiedDateTime,
	NameLocalized:              pm.D.NameLocalized,
	EndDate:                    pm.D.EndDate,
	EntityUUID:                 pm.D.EntityUUID,
	NameKoKR:                   pm.D.NameKoKR,
	CreatedDateTime:            pm.D.CreatedDateTime,
	DescriptionPtBR:            pm.D.DescriptionPtBR,
	DescriptionEsES:            pm.D.DescriptionEsES,
	NamePtBR:                   pm.D.NamePtBR,
	DescriptionNlNL:            pm.D.DescriptionNlNL,
	NameNlNL:                   pm.D.NameNlNL,
	DescriptionDefaultValue:    pm.D.DescriptionDefaultValue,
	NameDeDE:                   pm.D.NameDeDE,
	NameZhTW:                   pm.D.NameZhTW,
	Name:                       pm.D.Name,
	NameEsES:                   pm.D.NameEsES,
	DescriptionEnUS:            pm.D.DescriptionEnUS,
	GlStatementCode:            pm.D.GlStatementCode,
	DescriptionEnDEBUG:         pm.D.DescriptionEnDEBUG,
	DescriptionRuRU:            pm.D.DescriptionRuRU,
	Status:                     pm.D.Status,
	NameRuRU:                   pm.D.NameRuRU,
	DescriptionJaJP:            pm.D.DescriptionJaJP,
	DescriptionFrFR:            pm.D.DescriptionFrFR,
	NamePtPT:                   pm.D.NamePtPT,
	Description:                pm.D.Description,
	DescriptionDeDE:            pm.D.DescriptionDeDE,
	NameFrFR:                   pm.D.NameFrFR,
	NameEnDEBUG:                pm.D.NameEnDEBUG,
	NameJaJP:                   pm.D.NameJaJP,
	CreatedOn:                  pm.D.CreatedOn,
	NameEnUS:                   pm.D.NameEnUS,
	DescriptionZhTW:            pm.D.DescriptionZhTW,
	NameZhCN:                   pm.D.NameZhCN,
	CostcenterExternalObjectID: pm.D.CostcenterExternalObjectID,
	NameDefaultValue:           pm.D.NameDefaultValue,
	DescriptionEnGB:            pm.D.DescriptionEnGB,
	LastModifiedBy:             pm.D.LastModifiedBy,
	NameEnGB:                   pm.D.NameEnGB,
	LastModifiedOn:             pm.D.LastModifiedOn,
	DescriptionZhCN:            pm.D.DescriptionZhCN,
	CreatedBy:                  pm.D.CreatedBy,
	DescriptionLocalized:       pm.D.DescriptionLocalized,
	DescriptionPtPT:            pm.D.DescriptionPtPT,
	CostcenterManager:          pm.D.CostcenterManager,
	}, nil
}

func ConvertToJobCodeNav(raw []byte, l *logger.Logger) (*JobCodeNav, error) {
	pm := &responses.JobCodeNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to JobCodeNav. unmarshal error: %w", err)
	}
	
	return &JobCodeNav{
	ExternalCode:            pm.D.ExternalCode,
	StartDate:               pm.D.StartDate,
	SupervisorLevel:         pm.D.SupervisorLevel,
	NameKoKR:                pm.D.NameKoKR,
	CreatedDateTime:         pm.D.CreatedDateTime,
	CustString6PtPT:         pm.D.CustString6PtPT,
	NamePtBR:                pm.D.NamePtBR,
	CustString7JaJP:         pm.D.CustString7JaJP,
	ParentJobCode:           pm.D.ParentJobCode,
	WorkerCompCode:          pm.D.WorkerCompCode,
	CustString6Localized:    pm.D.CustString6Localized,
	CustString1DeDE:         pm.D.CustString1DeDE,
	CustString7RuRU:         pm.D.CustString7RuRU,
	CustString2FrFR:         pm.D.CustString2FrFR,
	Grade:                   pm.D.Grade,
	NameEsES:                pm.D.NameEsES,
	CustString7DeDE:         pm.D.CustString7DeDE,
	CustString1NlNL:         pm.D.CustString1NlNL,
	IsFulltimeEmployee:      pm.D.IsFulltimeEmployee,
	CustString9EnGB:         pm.D.CustString9EnGB,
	CustString8RuRU:         pm.D.CustString8RuRU,
	NameEnUS:                pm.D.NameEnUS,
	CustString2JaJP:         pm.D.CustString2JaJP,
	JobFunction:             pm.D.JobFunction,
	CustString1PtPT:         pm.D.CustString1PtPT,
	DescriptionZhTW:         pm.D.DescriptionZhTW,
	CustString1RuRU:         pm.D.CustString1RuRU,
	NameZhCN:                pm.D.NameZhCN,
	CustString2RuRU:         pm.D.CustString2RuRU,
	CustString6DeDE:         pm.D.CustString6DeDE,
	DescriptionEnGB:         pm.D.DescriptionEnGB,
	CustString7PtPT:         pm.D.CustString7PtPT,
	CustString6EsES:         pm.D.CustString6EsES,
	NameEnGB:                pm.D.NameEnGB,
	LastModifiedOn:          pm.D.LastModifiedOn,
	CustString1JaJP:         pm.D.CustString1JaJP,
	CustString2DefaultValue: pm.D.CustString2DefaultValue,
	DescriptionZhCN:         pm.D.DescriptionZhCN,
	CustString8JaJP:         pm.D.CustString8JaJP,
	CustString9ZhTW:         pm.D.CustString9ZhTW,
	CustString6EnUS:         pm.D.CustString6EnUS,
	CustString6ZhCN:         pm.D.CustString6ZhCN,
	DescriptionKoKR:         pm.D.DescriptionKoKR,
	LastModifiedDateTime:    pm.D.LastModifiedDateTime,
	CustString9DefaultValue: pm.D.CustString9DefaultValue,
	DescriptionPtBR:         pm.D.DescriptionPtBR,
	CustString8EnGB:         pm.D.CustString8EnGB,
	CustString6RuRU:         pm.D.CustString6RuRU,
	CustString9EnUS:         pm.D.CustString9EnUS,
	CustString9ZhCN:         pm.D.CustString9ZhCN,
	CustString2PtPT:         pm.D.CustString2PtPT,
	CustString8NlNL:         pm.D.CustString8NlNL,
	RegularTemporary:        pm.D.RegularTemporary,
	NameNlNL:                pm.D.NameNlNL,
	CustString9JaJP:         pm.D.CustString9JaJP,
	NameDeDE:                pm.D.NameDeDE,
	NameZhTW:                pm.D.NameZhTW,
	CustString7KoKR:         pm.D.CustString7KoKR,
	Name:                    pm.D.Name,
	CustString9Localized:    pm.D.CustString9Localized,
	DescriptionRuRU:         pm.D.DescriptionRuRU,
	NameRuRU:                pm.D.NameRuRU,
	CustString6EnGB:         pm.D.CustString6EnGB,
	NamePtPT:                pm.D.NamePtPT,
	CustString5:             pm.D.CustString5,
	CustString4:             pm.D.CustString4,
	CustString3:             pm.D.CustString3,
	CustString7NlNL:         pm.D.CustString7NlNL,
	Description:             pm.D.Description,
	DescriptionDeDE:         pm.D.DescriptionDeDE,
	CustString8EnUS:         pm.D.CustString8EnUS,
	CustString1FrFR:         pm.D.CustString1FrFR,
	CustString6KoKR:         pm.D.CustString6KoKR,
	EmployeeClass:           pm.D.EmployeeClass,
	CustString7EnDEBUG:      pm.D.CustString7EnDEBUG,
	DefaultSupervisorLevel:  pm.D.DefaultSupervisorLevel,
	DefaultJobLevel:         pm.D.DefaultJobLevel,
	CustString1DefaultValue: pm.D.CustString1DefaultValue,
	CustString7ZhCN:         pm.D.CustString7ZhCN,
	CustString1PtBR:         pm.D.CustString1PtBR,
	CustString2DeDE:         pm.D.CustString2DeDE,
	CustString6NlNL:         pm.D.CustString6NlNL,
	CreatedBy:               pm.D.CreatedBy,
	CustString9EsES:         pm.D.CustString9EsES,
	CustString7EnGB:         pm.D.CustString7EnGB,
	CustString6ZhTW:         pm.D.CustString6ZhTW,
	DescriptionPtPT:         pm.D.DescriptionPtPT,
	EndDate:                 pm.D.EndDate,
	CustString2EsES:         pm.D.CustString2EsES,
	CustString8KoKR:         pm.D.CustString8KoKR,
	CustString7ZhTW:         pm.D.CustString7ZhTW,
	CustString2PtBR:         pm.D.CustString2PtBR,
	CustString1EnGB:         pm.D.CustString1EnGB,
	DescriptionNlNL:         pm.D.DescriptionNlNL,
	CustString9PtPT:         pm.D.CustString9PtPT,
	CustString10:            pm.D.CustString10,
	IsRegular:               pm.D.IsRegular,
	CustString8EsES:         pm.D.CustString8EsES,
	CustString2ZhCN:         pm.D.CustString2ZhCN,
	CustString8ZhCN:         pm.D.CustString8ZhCN,
	CustString7EnUS:         pm.D.CustString7EnUS,
	CustString1ZhTW:         pm.D.CustString1ZhTW,
	Status:                  pm.D.Status,
	CustString7EsES:         pm.D.CustString7EsES,
	CustString8PtPT:         pm.D.CustString8PtPT,
	DescriptionFrFR:         pm.D.DescriptionFrFR,
	NameFrFR:                pm.D.NameFrFR,
	NameEnDEBUG:             pm.D.NameEnDEBUG,
	CreatedOn:               pm.D.CreatedOn,
	CustString8EnDEBUG:      pm.D.CustString8EnDEBUG,
	CustString8ZhTW:         pm.D.CustString8ZhTW,
	CustString2ZhTW:         pm.D.CustString2ZhTW,
	CustString9RuRU:         pm.D.CustString9RuRU,
	NameDefaultValue:        pm.D.NameDefaultValue,
	CustString1EnDEBUG:      pm.D.CustString1EnDEBUG,
	CustString9KoKR:         pm.D.CustString9KoKR,
	JobLevel:                pm.D.JobLevel,
	CustString2KoKR:         pm.D.CustString2KoKR,
	CustString6DefaultValue: pm.D.CustString6DefaultValue,
	CustString8DefaultValue: pm.D.CustString8DefaultValue,
	CustString6EnDEBUG:      pm.D.CustString6EnDEBUG,
	DefaultEmployeeClass:    pm.D.DefaultEmployeeClass,
	CustString9NlNL:         pm.D.CustString9NlNL,
	CustString2NlNL:         pm.D.CustString2NlNL,
	NameLocalized:           pm.D.NameLocalized,
	EntityUUID:              pm.D.EntityUUID,
	CustString1KoKR:         pm.D.CustString1KoKR,
	DescriptionEsES:         pm.D.DescriptionEsES,
	CustString8Localized:    pm.D.CustString8Localized,
	CustString1EsES:         pm.D.CustString1EsES,
	CustString6FrFR:         pm.D.CustString6FrFR,
	StandardHours:           pm.D.StandardHours,
	CustString6PtBR:         pm.D.CustString6PtBR,
	CustString9PtBR:         pm.D.CustString9PtBR,
	DescriptionDefaultValue: pm.D.DescriptionDefaultValue,
	CustString9FrFR:         pm.D.CustString9FrFR,
	CustString2EnGB:         pm.D.CustString2EnGB,
	CustString2Localized:    pm.D.CustString2Localized,
	CustString9EnDEBUG:      pm.D.CustString9EnDEBUG,
	CustString6JaJP:         pm.D.CustString6JaJP,
	DescriptionEnUS:         pm.D.DescriptionEnUS,
	DescriptionEnDEBUG:      pm.D.DescriptionEnDEBUG,
	DescriptionJaJP:         pm.D.DescriptionJaJP,
	CustString1ZhCN:         pm.D.CustString1ZhCN,
	NameJaJP:                pm.D.NameJaJP,
	CustString7Localized:    pm.D.CustString7Localized,
	CustString8DeDE:         pm.D.CustString8DeDE,
	CustString2EnDEBUG:      pm.D.CustString2EnDEBUG,
	CustString7PtBR:         pm.D.CustString7PtBR,
	CustString8FrFR:         pm.D.CustString8FrFR,
	CustString1EnUS:         pm.D.CustString1EnUS,
	LastModifiedBy:          pm.D.LastModifiedBy,
	CustomString5:           pm.D.CustomString5,
	CustomString4:           pm.D.CustomString4,
	CustomString3:           pm.D.CustomString3,
	CustomString2:           pm.D.CustomString2,
	CustString9DeDE:         pm.D.CustString9DeDE,
	CustomString9:           pm.D.CustomString9,
	CustomString8:           pm.D.CustomString8,
	CustomString7:           pm.D.CustomString7,
	CustString1Localized:    pm.D.CustString1Localized,
	CustomString6:           pm.D.CustomString6,
	CustomString1:           pm.D.CustomString1,
	CustString7DefaultValue: pm.D.CustString7DefaultValue,
	DescriptionLocalized:    pm.D.DescriptionLocalized,
	CustString2EnUS:         pm.D.CustString2EnUS,
	CustString8PtBR:         pm.D.CustString8PtBR,
	CustString7FrFR:         pm.D.CustString7FrFR,
	}, nil
}

func ConvertToDivisionNav(raw []byte, l *logger.Logger) (*DivisionNav, error) {
	pm := &responses.DivisionNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to DivisionNav. unmarshal error: %w", err)
	}
	
	return &DivisionNav{
	ExternalCode:            pm.D.ExternalCode,
	StartDate:               pm.D.StartDate,
	Parent:                  pm.D.Parent,
	DescriptionKoKR:         pm.D.DescriptionKoKR,
	LastModifiedDateTime:    pm.D.LastModifiedDateTime,
	NameLocalized:           pm.D.NameLocalized,
	EndDate:                 pm.D.EndDate,
	EntityUUID:              pm.D.EntityUUID,
	NameKoKR:                pm.D.NameKoKR,
	CreatedDateTime:         pm.D.CreatedDateTime,
	DescriptionPtBR:         pm.D.DescriptionPtBR,
	DescriptionEsES:         pm.D.DescriptionEsES,
	NamePtBR:                pm.D.NamePtBR,
	DescriptionNlNL:         pm.D.DescriptionNlNL,
	NameNlNL:                pm.D.NameNlNL,
	DescriptionDefaultValue: pm.D.DescriptionDefaultValue,
	NameDeDE:                pm.D.NameDeDE,
	NameZhTW:                pm.D.NameZhTW,
	Name:                    pm.D.Name,
	NameEsES:                pm.D.NameEsES,
	DescriptionEnUS:         pm.D.DescriptionEnUS,
	DescriptionEnDEBUG:      pm.D.DescriptionEnDEBUG,
	DescriptionRuRU:         pm.D.DescriptionRuRU,
	Status:                  pm.D.Status,
	NameRuRU:                pm.D.NameRuRU,
	DescriptionJaJP:         pm.D.DescriptionJaJP,
	DescriptionFrFR:         pm.D.DescriptionFrFR,
	NamePtPT:                pm.D.NamePtPT,
	Description:             pm.D.Description,
	DescriptionDeDE:         pm.D.DescriptionDeDE,
	NameFrFR:                pm.D.NameFrFR,
	NameEnDEBUG:             pm.D.NameEnDEBUG,
	NameJaJP:                pm.D.NameJaJP,
	CreatedOn:               pm.D.CreatedOn,
	HeadOfUnit:              pm.D.HeadOfUnit,
	NameEnUS:                pm.D.NameEnUS,
	DescriptionZhTW:         pm.D.DescriptionZhTW,
	NameZhCN:                pm.D.NameZhCN,
	NameDefaultValue:        pm.D.NameDefaultValue,
	DescriptionEnGB:         pm.D.DescriptionEnGB,
	LastModifiedBy:          pm.D.LastModifiedBy,
	NameEnGB:                pm.D.NameEnGB,
	LastModifiedOn:          pm.D.LastModifiedOn,
	DescriptionZhCN:         pm.D.DescriptionZhCN,
	CreatedBy:               pm.D.CreatedBy,
	DescriptionLocalized:    pm.D.DescriptionLocalized,
	DescriptionPtPT:         pm.D.DescriptionPtPT,
	}, nil
}

func ConvertToParentPosition(raw []byte, l *logger.Logger) (*ParentPosition, error) {
	pm := &responses.ParentPosition{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ParentPosition. unmarshal error: %w", err)
	}
	
	return &ParentPosition{
	Code:                        pm.D.Code,
	EffectiveStartDate:          pm.D.EffectiveStartDate,
	BusinessUnit:                pm.D.BusinessUnit,
	LastModifiedDateTime:        pm.D.LastModifiedDateTime,
	JobTitle:                    pm.D.JobTitle,
	Criticality:                 pm.D.Criticality,
	CreatedDateTime:             pm.D.CreatedDateTime,
	JobCode:                     pm.D.JobCode,
	MdfSystemVersionID:          pm.D.MdfSystemVersionID,
	Type:                        pm.D.Type,
	Incumbent:                   pm.D.Incumbent,
	Division:                    pm.D.Division,
	ExternalNameZhTW:            pm.D.ExternalNameZhTW,
	ExternalNameJaJP:            pm.D.ExternalNameJaJP,
	MdfSystemEntityID:           pm.D.MdfSystemEntityID,
	PayRange:                    pm.D.PayRange,
	RegularTemporary:            pm.D.RegularTemporary,
	ExternalNamePtPT:            pm.D.ExternalNamePtPT,
	CostCenter:                  pm.D.CostCenter,
	StandardHours:               pm.D.StandardHours,
	LegacyPositionID:            pm.D.LegacyPositionID,
	ExternalNameZhCN:            pm.D.ExternalNameZhCN,
	ExternalNameLocalized:       pm.D.ExternalNameLocalized,
	MdfSystemRecordStatus:       pm.D.MdfSystemRecordStatus,
	Vacant:                      pm.D.Vacant,
	EffectiveStatus:             pm.D.EffectiveStatus,
	ExternalNameFrFR:            pm.D.ExternalNameFrFR,
	TechnicalParameters:         pm.D.TechnicalParameters,
	ExternalNameEnGB:            pm.D.ExternalNameEnGB,
	EffectiveEndDate:            pm.D.EffectiveEndDate,
	PositionCriticality:         pm.D.PositionCriticality,
	ExternalNameNlNL:            pm.D.ExternalNameNlNL,
	PositionTitle:               pm.D.PositionTitle,
	Description:                 pm.D.Description,
	ExternalNameDefaultValue:    pm.D.ExternalNameDefaultValue,
	ExternalNameEsES:            pm.D.ExternalNameEsES,
	PositionControlled:          pm.D.PositionControlled,
	ExternalNamePtBR:            pm.D.ExternalNamePtBR,
	PayGrade:                    pm.D.PayGrade,
	Company:                     pm.D.Company,
	Department:                  pm.D.Department,
	EmployeeClass:               pm.D.EmployeeClass,
	MdfSystemObjectType:         pm.D.MdfSystemObjectType,
	CreationSource:              pm.D.CreationSource,
	ChangeReason:                pm.D.ChangeReason,
	TargetFTE:                   pm.D.TargetFTE,
	LastModifiedDate:            pm.D.LastModifiedDate,
	ExternalNameRuRU:            pm.D.ExternalNameRuRU,
	LastModifiedBy:              pm.D.LastModifiedBy,
	LastModifiedDateWithTZ:      pm.D.LastModifiedDateWithTZ,
	JobLevel:                    pm.D.JobLevel,
	TransactionSequence:         pm.D.TransactionSequence,
	ExternalNameDeDE:            pm.D.ExternalNameDeDE,
	ExternalNameKoKR:            pm.D.ExternalNameKoKR,
	CreatedDate:                 pm.D.CreatedDate,
	CreatedBy:                   pm.D.CreatedBy,
	MdfSystemOptimisticLockUUID: pm.D.MdfSystemOptimisticLockUUID,
	MdfSystemRecordID:           pm.D.MdfSystemRecordID,
	Comment:                     pm.D.Comment,
	Location:                    pm.D.Location,
	MultipleIncumbentsAllowed:   pm.D.MultipleIncumbentsAllowed,
	ExternalNameEnUS:            pm.D.ExternalNameEnUS,
	ExternalNameEnDEBUG:         pm.D.ExternalNameEnDEBUG,
	}, nil
}

func ConvertToPayGradeNav(raw []byte, l *logger.Logger) (*PayGradeNav, error) {
	pm := &responses.PayGradeNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PayGradeNav. unmarshal error: %w", err)
	}
	
	return &PayGradeNav{
	ExternalCode:              pm.D.ExternalCode,
	StartDate:                 pm.D.StartDate,
	LastModifiedDateTime:      pm.D.LastModifiedDateTime,
	EndDate:                   pm.D.EndDate,
	LastModifiedBy:            pm.D.LastModifiedBy,
	CreatedDateTime:           pm.D.CreatedDateTime,
	Description:               pm.D.Description,
	CreatedOn:                 pm.D.CreatedOn,
	LastModifiedOn:            pm.D.LastModifiedOn,
	PaygradeLevel:             pm.D.PaygradeLevel,
	CustomString1:             pm.D.CustomString1,
	CreatedBy:                 pm.D.CreatedBy,
	Name:                      pm.D.Name,
	Status:                    pm.D.Status,
	}, nil
}
