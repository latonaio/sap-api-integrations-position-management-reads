package responses

type EffectiveStatusNav struct {
	D struct {
		Metadata struct {
			URI  string `json:"uri"`
			Type string `json:"type"`
		} `json:"__metadata"`
		Value     string `json:"value"`
		Key       string `json:"key"`
		EnDEBUG   string `json:"en_DEBUG"`
		RuRU      string `json:"ru_RU"`
		Localized string `json:"localized"`
		PtBR      string `json:"pt_BR"`
		FrFR      string `json:"fr_FR"`
		JaJP      string `json:"ja_JP"`
		DeDE      string `json:"de_DE"`
		EnGB      string `json:"en_GB"`
		ZhTW      string `json:"zh_TW"`
		KoKR      string `json:"ko_KR"`
		EnUS      string `json:"en_US"`
		EsES      string `json:"es_ES"`
		ZhCN      string `json:"zh_CN"`
		NlNL      string `json:"nl_NL"`
		PtPT      string `json:"pt_PT"`
	} `json:"d"`
}
