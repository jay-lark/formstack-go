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

func GetForms(fo FormstackOptions) ([]Form, error) {
	path := "/form.json"
	fmt.Println(fo)
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
