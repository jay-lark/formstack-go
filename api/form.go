package formstack

import (
	"encoding/json"
	"fmt"
)

type Form struct {
	ID                              string      `json:"id"`
	Created                         string      `json:"created"`
	Db                              string      `json:"db"`
	Deleted                         string      `json:"deleted"`
	Folder                          string      `json:"folder"`
	Language                        string      `json:"language"`
	Name                            string      `json:"name"`
	NumColumns                      string      `json:"num_columns"`
	ProgressMeter                   string      `json:"progress_meter"`
	Submissions                     string      `json:"submissions"`
	SubmissionsUnread               string      `json:"submissions_unread"`
	Updated                         string      `json:"updated"`
	Viewkey                         string      `json:"viewkey"`
	Views                           string      `json:"views"`
	SubmissionsToday                int         `json:"submissions_today"`
	LastSubmissionID                string      `json:"last_submission_id"`
	LastSubmissionTime              string      `json:"last_submission_time"`
	URL                             string      `json:"url"`
	Encrypted                       bool        `json:"encrypted"`
	ThumbnailURL                    interface{} `json:"thumbnail_url"`
	SubmitButtonTitle               string      `json:"submit_button_title"`
	Inactive                        bool        `json:"inactive"`
	Timezone                        string      `json:"timezone"`
	ShouldDisplayOneQuestionAtATime bool        `json:"should_display_one_question_at_a_time"`
	CanAccess1QFeature              bool        `json:"can_access_1q_feature"`
	IsWorkflowForm                  bool        `json:"is_workflow_form"`
	IsWorkflowPublished             bool        `json:"is_workflow_published"`
	HasApprovers                    bool        `json:"has_approvers"`
	EditURL                         string      `json:"edit_url"`
	DataURL                         string      `json:"data_url"`
	SummaryURL                      string      `json:"summary_url"`
	RssURL                          string      `json:"rss_url"`
	Permissions                     int         `json:"permissions"`
	CanEdit                         bool        `json:"can_edit"`
}

type FormId struct {
	Created                         string      `json:"created"`
	Db                              string      `json:"db"`
	Deleted                         string      `json:"deleted"`
	Folder                          string      `json:"folder"`
	ID                              string      `json:"id"`
	Language                        string      `json:"language"`
	Name                            string      `json:"name"`
	NumColumns                      string      `json:"num_columns"`
	ProgressMeter                   string      `json:"progress_meter"`
	Submissions                     string      `json:"submissions"`
	SubmissionsUnread               string      `json:"submissions_unread"`
	Updated                         string      `json:"updated"`
	Viewkey                         string      `json:"viewkey"`
	Views                           string      `json:"views"`
	SubmissionsToday                int         `json:"submissions_today"`
	URL                             string      `json:"url"`
	Encrypted                       bool        `json:"encrypted"`
	ThumbnailURL                    interface{} `json:"thumbnail_url"`
	SubmitButtonTitle               string      `json:"submit_button_title"`
	Inactive                        bool        `json:"inactive"`
	Timezone                        string      `json:"timezone"`
	ShouldDisplayOneQuestionAtATime bool        `json:"should_display_one_question_at_a_time"`
	CanAccess1QFeature              bool        `json:"can_access_1q_feature"`
	IsWorkflowForm                  bool        `json:"is_workflow_form"`
	IsWorkflowPublished             bool        `json:"is_workflow_published"`
	HasApprovers                    bool        `json:"has_approvers"`
	EditURL                         string      `json:"edit_url"`
	DataURL                         string      `json:"data_url"`
	SummaryURL                      string      `json:"summary_url"`
	RssURL                          string      `json:"rss_url"`
	Permissions                     int         `json:"permissions"`
	CanEdit                         bool        `json:"can_edit"`
	Javascript                      string      `json:"javascript"`
	HTML                            string      `json:"html"`
	Fields                          FormFields
}

type FormFields struct {
	ID                    string      `json:"id"`
	Label                 string      `json:"label"`
	HideLabel             string      `json:"hide_label"`
	Description           string      `json:"description"`
	Name                  string      `json:"name"`
	Type                  string      `json:"type"`
	Options               string      `json:"options"`
	Required              string      `json:"required"`
	Uniq                  string      `json:"uniq"`
	Hidden                string      `json:"hidden"`
	Readonly              string      `json:"readonly"`
	Colspan               string      `json:"colspan"`
	Sort                  string      `json:"sort"`
	Logic                 interface{} `json:"logic"`
	Calculation           string      `json:"calculation"`
	WorkflowAccess        string      `json:"workflow_access"`
	Default               string      `json:"default"`
	ShowPrefix            int         `json:"show_prefix"`
	ShowMiddle            int         `json:"show_middle"`
	ShowInitial           int         `json:"show_initial"`
	ShowSuffix            int         `json:"show_suffix"`
	TextSize              int         `json:"text_size"`
	MiddleInitialOptional int         `json:"middle_initial_optional"`
	MiddleNameOptional    int         `json:"middle_name_optional"`
	PrefixOptional        int         `json:"prefix_optional"`
	SuffixOptional        int         `json:"suffix_optional"`
	VisibleSubfields      []string    `json:"visible_subfields"`
}

func GetForms(fo FormstackOptions) ([]Form, error) {
	path := "/form.json"
	res, err := clientDo(fo, "GET", path, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	forms := map[string]json.RawMessage{}
	err = json.NewDecoder(res.Body).Decode(&forms)
	if err != nil {
		return nil, err
	}

	var frm []Form
	json.Unmarshal(forms["forms"], &frm)

	return frm, nil

}

func GetFormById(fo FormstackOptions, id string) (*FormId, error) {
	path := fmt.Sprintf("/form/%s.json", id)
	res, err := clientDo(fo, "GET", path, nil)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var form FormId
	json.NewDecoder(res.Body).Decode(&form)
	frm := &FormId{
		ID:   form.ID,
		Name: form.Name,
	}

	return frm, nil

}

func NewForm(fo FormstackOptions, name string) {
	url := "https://www.formstack.com/api/v2/form.json"

	payload := strings.NewReader("{\"db\":false,\"label_position\":\"top\",\"submit_button_title\":\"Submit Form\",\"use_ssl\":false,\"timezone\":\"US/Eastern\",\"language\":\"en\",\"active\":false,\"name\":name}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer e37f25b2600a521ed38a7d3f40b2559d")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}