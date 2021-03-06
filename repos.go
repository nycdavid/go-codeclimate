package codeclimate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Repo struct {
	Body struct {
		Attributes Attributes `json:"attributes"`
	} `json:"data"`
}

type Attributes struct {
	Score           float64 `json:"score"`
	AnalysisVersion int     `json:"analysis_version"`
	Branch          string  `json:"branch"`
}

func (c client) GetRepo() (Repo, error) {
	var repo Repo
	u, e := url.Parse(c.BaseUrl)
	if e != nil {
		return Repo{}, e
	}
	u.Path = fmt.Sprintf("/repos/%s", c.AppId)
	req, e := http.NewRequest(http.MethodGet, u.String(), nil)
	if e != nil {
		return Repo{}, e
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.ApiKey))
	httpres, e := c.httpClient.Do(req)
	if e != nil {
		return Repo{}, e
	}
	var buf bytes.Buffer
	buf.ReadFrom(httpres.Body)
	dec := json.NewDecoder(&buf)
	e = dec.Decode(&repo)
	if e != nil {
		fmt.Println(e.Error())
	}
	return repo, nil
}
