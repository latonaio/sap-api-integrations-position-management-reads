# sap-api-integrations-position-management-reads
sap-api-integrations-position-management-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で ポジション管理 データを取得するマイクロサービスです。    
sap-api-integrations-position-management-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-position-management-reads は、オンプレミス版である（＝クラウド版ではない）SAP SuccessFactors API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/ECEmploymentInformation/overview   

## 動作環境  
sap-api-integrations-position-management-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-position-management-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-position-management-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/ECPositionManagement/overview    
* APIサービス名(=baseURL): https://sandbox.api.sap.com/successfactors/

## 本レポジトリ に 含まれる API名
sap-api-integrations-position-management-reads には、次の API をコールするためのリソースが含まれています。  

* Position（ポジション管理 - ヘッダ）※ポジション管理の詳細データを取得するために、RegularTemporaryNav、ExternalNameTranslationTextNav、CompanyNav、EffectiveStatusNav、DepartmentNav、BusinessUnitNav、LocationNav、JobLevelNav、CostCenterNav、JobCodeNav。DivisionNav、ParentPosition、PayGradeNav、と合わせて取得されます。
* RegularTemporaryNav（ポジション管理 - 正規非正規情報 ※To）
* ExternalNameTranslationTextNav（ポジション管理 - 外部名称翻訳テキスト情報 ※To）
* CompanyNav（ポジション管理 - 会社情報 ※To）
* EffectiveStatusNav（ポジション管理 - 有効ステータス情報 ※To）
* DepartmentNav（ポジション管理 - 部門_Department 情報 ※To）
* BusinessUnitNav（ポジション管理 - 事業ユニット情報 ※To）
* LocationNav（ポジション管理 - ロケーション情報 ※To）
* JobLevelNav（ポジション管理 - ジョブレベル情報 ※To）
* CostCenterNav（ポジション管理 - 原価センタ情報 ※To）
* JobCodeNav（ポジション管理 - ジョブコード情報 ※To）
* DivisionNav（ポジション管理 - 部門_Division 情報 ※To）
* ParentPosition（ポジション管理 - 親ポジション情報 ※To）
* PayGradeNav（ポジション管理 - 支払等級情報 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-position-management-readsにおいて、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.PositionManagement.Code（コード）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "Position",
	"accepter": ["Header"],
	"code": "3000320",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "Position",
	"accepter": ["All"],
	"code": "3000320",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
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
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、ポジション管理 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"Code" ～ "PayGradeNav" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-position-management-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-position-management-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"Code": "3000320",
			"EffectiveStartDate": "/Date(631152000000)/",
			"BusinessUnit": "PRODS",
			"LastModifiedDateTime": "/Date(1515681212000+0000)/",
			"JobTitle": "SVP Operations & Maintenance",
			"Criticality": "",
			"CreatedDateTime": "/Date(1442544406000+0000)/",
			"JobCode": "50071000",
			"MdfSystemVersionID": "",
			"Type": "RP",
			"Incumbent": "",
			"Division": "MANU",
			"ExternalNameZhTW": "運營和維護高級副總裁",
			"ExternalNameJaJP": "SVP 作業/メンテナンス",
			"MdfSystemEntityID": "1C910D43480E49D7A6D13133EDB1C701",
			"PayRange": "",
			"RegularTemporary": "R",
			"ExternalNamePtPT": "",
			"CostCenter": "8500-4200",
			"StandardHours": "",
			"LegacyPositionID": "",
			"ExternalNameZhCN": "运营和维护高级副总裁",
			"ExternalNameLocalized": "SVP Operations & Maintenance",
			"MdfSystemRecordStatus": "N",
			"Vacant": true,
			"EffectiveStatus": "A",
			"ExternalNameFrFR": "SVP opérations et maintenance",
			"TechnicalParameters": "",
			"ExternalNameEnGB": "SVP Operations & Maintenance",
			"EffectiveEndDate": "/Date(253402214400000)/",
			"PositionCriticality": "0",
			"ExternalNameNlNL": "",
			"PositionTitle": "SVP Operations & Maintenance",
			"Description": "",
			"ExternalNameDefaultValue": "SVP Operations & Maintenance",
			"ExternalNameEsES": "Vicepresidente sénior (operaciones y mantenimiento)",
			"PositionControlled": "",
			"ExternalNamePtBR": "SVP Operações e manutenção",
			"PayGrade": "GR-14",
			"Company": "4000",
			"Department": "5000017",
			"EmployeeClass": "1",
			"MdfSystemObjectType": "Position",
			"CreationSource": "",
			"ChangeReason": "",
			"TargetFTE": "1",
			"LastModifiedDate": "/Date(1515663212000)/",
			"ExternalNameRuRU": "Старший вице-президент по эксплуатации и обслуживанию",
			"LastModifiedBy": "sfadmin",
			"LastModifiedDateWithTZ": "/Date(1515681212000+0000)/",
			"JobLevel": "SR PRES",
			"TransactionSequence": "1",
			"ExternalNameDeDE": "Unternehmensbereichsleiter Betrieb und Wartung",
			"ExternalNameKoKR": "SVP 운영 및 유지보수",
			"CreatedDate": "/Date(1442530006000)/",
			"CreatedBy": "sfadmin",
			"MdfSystemOptimisticLockUUID": "E3BAD896990C479E9C02A1B8D5E1EB35",
			"MdfSystemRecordID": "EDF56EAF7B844BB78DFDCCC8AA682854",
			"Comment": "",
			"Location": "6200-0001",
			"MultipleIncumbentsAllowed": false,
			"ExternalNameEnUS": "SVP Operations & Maintenance",
			"ExternalNameEnDEBUG": "",
			"RegularTemporaryNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/regularTemporaryNav",
			"ExternalNameTranslationTextNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/externalNameTranslationTextNav",
			"CompanyNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/companyNav",
			"EffectiveStatusNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/effectiveStatusNav",
			"DepartmentNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/departmentNav",
			"BusinessUnitNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/businessUnitNav",
			"LocationNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/locationNav",
			"CostCenterNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/costCenterNav",
			"JobCodeNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/jobCodeNav",
			"DivisionNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/divisionNav",
			"JobLevelNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/jobLevelNav",
			"ParentPosition": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/parentPosition",
			"PayGradeNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/Position(code='3000320',effectiveStartDate=datetime'1990-01-01T00:00:00')/payGradeNav"
		}
	],
	"time": "2022-01-14T22:53:36.54887+09:00"
}
```
