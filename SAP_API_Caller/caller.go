package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-position-management-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetPositionManagement(code string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(code)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(code string) {
	headerData, err := c.callPositionManagementSrvAPIRequirementHeader("Position", code)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	regularTemporaryNavData, err := c.callRegularTemporaryNav(headerData[0].RegularTemporaryNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(regularTemporaryNavData)
	
	externalNameTranslationTextNavData, err := c.callExternalNameTranslationTextNav(headerData[0].ExternalNameTranslationTextNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(externalNameTranslationTextNavData)
	
	companyNavData, err := c.callCompanyNav(headerData[0].CompanyNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(companyNavData)

	effectiveStatusNavData, err := c.callEffectiveStatusNav(headerData[0].EffectiveStatusNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(effectiveStatusNavData)
	
	departmentNavData, err := c.callDepartmentNav(headerData[0].DepartmentNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(departmentNavData)
	
	businessUnitNavData, err := c.callBusinessUnitNav(headerData[0].BusinessUnitNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(businessUnitNavData)
	
	locationNavData, err := c.callLocationNav(headerData[0].LocationNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(locationNavData)
	
	jobLevelNavData, err := c.callJobLevelNav(headerData[0].JobLevelNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(jobLevelNavData)
	
	costCenterNavData, err := c.callCostCenterNav(headerData[0].CostCenterNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(costCenterNavData)
	
	jobCodeNavData, err := c.callJobCodeNav(headerData[0].JobCodeNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(jobCodeNavData)
	
	divisionNavData, err := c.callDivisionNav(headerData[0].DivisionNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(divisionNavData)
	
	parentPositionData, err := c.callParentPosition(headerData[0].ParentPosition)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(parentPositionData)
	
	payGradeNavData, err := c.callPayGradeNav(headerData[0].PayGradeNav)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(payGradeNavData)
}

func (c *SAPAPICaller) callPositionManagementSrvAPIRequirementHeader(api, code string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "odata/v2", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, code)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callRegularTemporaryNav(url string) (*sap_api_output_formatter.RegularTemporaryNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRegularTemporaryNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callExternalNameTranslationTextNav(url string) ([]sap_api_output_formatter.ExternalNameTranslationTextNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToExternalNameTranslationTextNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCompanyNav(url string) (*sap_api_output_formatter.CompanyNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompanyNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callEffectiveStatusNav(url string) (*sap_api_output_formatter.EffectiveStatusNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToEffectiveStatusNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callDepartmentNav(url string) (*sap_api_output_formatter.DepartmentNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToDepartmentNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callBusinessUnitNav(url string) (*sap_api_output_formatter.BusinessUnitNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBusinessUnitNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callLocationNav(url string) (*sap_api_output_formatter.LocationNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToLocationNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callJobLevelNav(url string) (*sap_api_output_formatter.JobLevelNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToJobLevelNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callCostCenterNav(url string) (*sap_api_output_formatter.CostCenterNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCostCenterNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callJobCodeNav(url string) (*sap_api_output_formatter.JobCodeNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToJobCodeNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callDivisionNav(url string) (*sap_api_output_formatter.DivisionNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToDivisionNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callParentPosition(url string) (*sap_api_output_formatter.ParentPosition, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToParentPosition(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callPayGradeNav(url string) (*sap_api_output_formatter.PayGradeNav, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPayGradeNav(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, code string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("code eq '%s'", code))
	req.URL.RawQuery = params.Encode()
}
