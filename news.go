package nsddata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//easyjson:json
type NewsResponse []News

//easyjson:json
type News struct {
	TitleRu       string `json:"title_ru"`
	BodyRu        string `json:"body_ru"`
	NewsThemeID   int    `json:"news_theme_id"`
	GroupID       int    `json:"group_id"`
	PubDate       string `json:"pub_date"`
	NewsDate      string `json:"news_date"`
	ContentIDOut  string `json:"content_id_out"`
	ActionID      int    `json:"action_id"`
	AnnounceRu    string `json:"announce_ru"`
	ForQuInvestor string `json:"for_qu_investor"`
	CaType        string `json:"ca_type"`
	Data          struct {
		ID             int `json:"id"`
		CorpActionType struct {
			ID   int    `json:"id"`
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"corp_action_type"`
		State struct {
			ID   int    `json:"id"`
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"state"`
		ActionDatePlan string `json:"action_date_plan"`
		ActionDateCalc string `json:"action_date_calc"`
		RecordDateCalc string `json:"record_date_calc"`
		RecordDatePlan string `json:"record_date_plan"`
		RecordTime     string `json:"record_time"`
		PriorityIssue  struct {
			DecisionDate        string `json:"decision_date"`
			LawSection          string `json:"law_section"`
			DateFrom            string `json:"date_from"`
			DateTo              string `json:"date_to"`
			InitiatorExpireDate string `json:"initiator_expire_date"`
			NsdExpireDate       string `json:"nsd_expire_date"`
			ExpireDate          string `json:"expire_date"`
			NsdActualDate       string `json:"nsd_actual_date"`
			Items               []struct {
				Security struct {
					ID        int    `json:"id"`
					Isin      string `json:"isin"`
					CodeNsd   string `json:"code_nsd"`
					NameFull  string `json:"name_full"`
					InstrType struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"instr_type"`
					StateRegNumber string `json:"state_reg_number"`
				} `json:"security"`
				NewSecurity struct {
					ID        int    `json:"id"`
					Isin      string `json:"isin"`
					CodeNsd   string `json:"code_nsd"`
					NameFull  string `json:"name_full"`
					InstrType struct {
						ID   int    `json:"id"`
						Name string `json:"name"`
					} `json:"instr_type"`
					StateRegNumber string `json:"state_reg_number"`
				} `json:"new_security"`
				Price    int `json:"price"`
				Currency struct {
					ID        int    `json:"id"`
					Code      string `json:"code"`
					NameShort string `json:"name_short"`
					NameFull  string `json:"name_full"`
				} `json:"currency"`
			} `json:"items"`
		} `json:"priority_issue"`
		Securities []struct {
			ID        int    `json:"id"`
			Isin      string `json:"isin"`
			CodeNsd   string `json:"code_nsd"`
			NameFull  string `json:"name_full"`
			InstrType struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"instr_type"`
			StateRegNumber string `json:"state_reg_number"`
			StateRegDate   string `json:"state_reg_date"`
			Share          struct {
				Category struct {
					ID        int    `json:"id"`
					NameFull  string `json:"name_full"`
					NameShort string `json:"name_short"`
				} `json:"category"`
				FaceValue float64 `json:"face_value"`
				Currency  struct {
					ID        int    `json:"id"`
					Code      string `json:"code"`
					NameShort string `json:"name_short"`
					NameFull  string `json:"name_full"`
				} `json:"currency"`
			} `json:"share"`
			Issuer struct {
				ID          int    `json:"id"`
				CodeNsd     string `json:"code_nsd"`
				NameFull    string `json:"name_full"`
				NameShort   string `json:"name_short"`
				NameFullTr  string `json:"name_full_tr"`
				NameShortTr string `json:"name_short_tr"`
				Inn         string `json:"inn"`
				Ogrn        string `json:"ogrn"`
				LeiCode     string `json:"lei_code"`
				TaxNumber   string `json:"tax_number"`
			} `json:"issuer"`
		} `json:"securities"`
	} `json:"data"`
	Category string `json:"category"`
}

const (
	NewsMethod = "get/news"

	FilterEQ     = "$eq"
	FilterNOT    = "$ne"
	FilterIN     = "$in"
	FilterNOTIN  = "$nin"
	FilterMORE   = "$gt"
	FilterMOREEQ = "$gte"
	FilterLESS   = "$lt"
	FilterLESSEQ = "$lte"
)

func (n *NSDDataAPIClient) GetNews(limit, skip int, filter map[string]map[string]interface{}) NewsResponse {
	var resp NewsResponse

	n.debug("Generating base url for get news method")

	url := n.O.BaseURL + NewsMethod

	n.debug("Base url is " + url)

	params := "?apikey=" + n.O.APIKey

	if limit > 0 {
		params += "&"

		n.debug(fmt.Sprintf("Set limit %d", limit))

		params += fmt.Sprintf("limit=%d", limit)
	}

	if skip > 0 {
		params += "&"

		n.debug(fmt.Sprintf("Set skip %d", skip))

		params += fmt.Sprintf("skip=%d", skip)
	}

	if filter != nil && len(filter) > 0 {
		params += "&"

		bFilter, err := json.Marshal(filter)
		if err != nil {

		}

		sFilter := string(bFilter)
		n.debug(fmt.Sprintf("Set filter %s", sFilter))

		params += "filter=" + sFilter
	}
	url += params

	fmt.Println(url)
	r, err := n.O.Client.Get(url)

	if err != nil {
		panic(err)
	}

	bResp, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	err = resp.UnmarshalJSON(bResp)

	if err != nil {
		panic(err)
	}

	return resp
}
