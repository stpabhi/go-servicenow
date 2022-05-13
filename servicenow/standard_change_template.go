package servicenow

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// StandardChangeTemplatesService handles the communication with the StandardChangeTemplate related
// methods of the ServiceNow API
type StandardChangeTemplatesService service

// StandardChangeTemplate represents a Standard Change Template
type StandardChangeTemplate struct {
	Status                         *string `json:"__status,omitempty"`
	Active                         *string `json:"active,omitempty"`
	ActivityDue                    *string `json:"activity_due,omitempty"`
	AdditionalAssigneeList         *string `json:"additional_assignee_list,omitempty"`
	Approval                       *string `json:"approval,omitempty"`
	ApprovalHistory                *string `json:"approval_history,omitempty"`
	ApprovalSet                    *string `json:"approval_set,omitempty"`
	AssignedTo                     *string `json:"assigned_to,omitempty"`
	AssignmentGroup                *string `json:"assignment_group,omitempty"`
	BusinessDuration               *string `json:"business_duration,omitempty"`
	BusinessJustification          *string `json:"business_justification,omitempty"`
	BusinessService                *string `json:"business_service,omitempty"`
	CalendarDuration               *string `json:"calendar_duration,omitempty"`
	Catalog                        *string `json:"catalog,omitempty"`
	Category                       *string `json:"category,omitempty"`
	ChangeRequests                 *string `json:"change_requests,omitempty"`
	ClosedAt                       *string `json:"closed_at,omitempty"`
	ClosedBy                       *string `json:"closed_by,omitempty"`
	CloseNotes                     *string `json:"close_notes,omitempty"`
	CmdbCi                         *string `json:"cmdb_ci,omitempty"`
	Comments                       *string `json:"comments,omitempty"`
	CommentsAndWorkNotes           *string `json:"comments_and_work_notes,omitempty"`
	Company                        *string `json:"company,omitempty"`
	ContactType                    *string `json:"contact_type,omitempty"`
	CorrelationDisplay             *string `json:"correlation_display,omitempty"`
	CorrelationID                  *string `json:"correlation_id,omitempty"`
	CreatedFromChange              *string `json:"created_from_change,omitempty"`
	Description                    *string `json:"description,omitempty"`
	DueDate                        *string `json:"due_date,omitempty"`
	Escalation                     *string `json:"escalation,omitempty"`
	ExpectedStart                  *string `json:"expected_start,omitempty"`
	FollowUp                       *string `json:"follow_up,omitempty"`
	GroupList                      *string `json:"group_list,omitempty"`
	Impact                         *string `json:"impact,omitempty"`
	Knowledge                      *string `json:"knowledge,omitempty"`
	Location                       *string `json:"location,omitempty"`
	MadeSLA                        *string `json:"made_sla,omitempty"`
	Number                         *string `json:"number,omitempty"`
	OpenedAt                       *string `json:"opened_at,omitempty"`
	OpenedBy                       *string `json:"opened_by,omitempty"`
	Order                          *string `json:"order,omitempty"`
	Parent                         *string `json:"parent,omitempty"`
	Priority                       *string `json:"priority,omitempty"`
	ProposalType                   *string `json:"proposal_type,omitempty"`
	ReassignmentCount              *string `json:"reassignment_count,omitempty"`
	RouteReason                    *string `json:"route_reason,omitempty"`
	ServiceOffering                *string `json:"service_offering,omitempty"`
	ShortDescription               *string `json:"short_description,omitempty"`
	Skills                         *string `json:"skills,omitempty"`
	SLADue                         *string `json:"sla_due,omitempty"`
	SnEsignDocument                *string `json:"sn_esign_document,omitempty"`
	SnEsignEsignatureConfiguration *string `json:"sn_esign_esignature_configuration,omitempty"`
	State                          *string `json:"state,omitempty"`
	StdChangeProducer              *string `json:"std_change_producer,omitempty"`
	StdChangeProducerVersion       *string `json:"std_change_producer_version,omitempty"`
	SysClassName                   *string `json:"sys_class_name,omitempty"`
	SysCreatedBy                   *string `json:"sys_created_by,omitempty"`
	SysCreatedOn                   *string `json:"sys_created_on,omitempty"`
	SysDomain                      *string `json:"sys_domain,omitempty"`
	SysDomainPath                  *string `json:"sys_domain_path,omitempty"`
	SysID                          *string `json:"sys_id,omitempty"`
	SysModCount                    *string `json:"sys_mod_count,omitempty"`
	SysTags                        *string `json:"sys_tags,omitempty"`
	SysUpdatedBy                   *string `json:"sys_updated_by,omitempty"`
	SysUpdatedOn                   *string `json:"sys_updated_on,omitempty"`
	TaskEffectiveNumber            *string `json:"task_effective_number,omitempty"`
	TemplateName                   *string `json:"template_name,omitempty"`
	TemplateValue                  *string `json:"template_value,omitempty"`
	TimeWorked                     *string `json:"time_worked,omitempty"`
	UniversalRequest               *string `json:"universal_request,omitempty"`
	UponApproval                   *string `json:"upon_approval,omitempty"`
	UponReject                     *string `json:"upon_reject,omitempty"`
	Urgency                        *string `json:"urgency,omitempty"`
	UserInput                      *string `json:"user_input,omitempty"`
	WatchList                      *string `json:"watch_list,omitempty"`
	WorkEnd                        *string `json:"work_end,omitempty"`
	WorkNotes                      *string `json:"work_notes,omitempty"`
	WorkNotesList                  *string `json:"work_notes_list,omitempty"`
	WorkStart                      *string `json:"work_start,omitempty"`

	Extra map[string]string `json:"-"`
}

func (s StandardChangeTemplate) String() string {
	return Stringify(s)
}

func (s StandardChangeTemplate) MarshalJSON() ([]byte, error) {
	type changeTemplate StandardChangeTemplate
	b, _ := json.Marshal(changeTemplate(s))

	var m map[string]json.RawMessage
	_ = json.Unmarshal(b, &m)

	for k, v := range s.Extra {
		b, _ := json.Marshal(v)
		m[k] = b
	}

	return json.Marshal(m)
}

// List standard change templates
func (s *StandardChangeTemplatesService) List(ctx context.Context, opts ListOptions) ([]*StandardChangeTemplate, *Response, error) {
	u := fmt.Sprint("/std_change_proposal.do")
	var params []string
	for _, v := range opts.QueryOpts {
		params = append(params, fmt.Sprintf("%s%s%s", v.Key, v.Op, v.Val))
	}
	opts.internalFields.SysparmQuery = strings.Join(params[:], "+")
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		StandardChangeTemplates []*StandardChangeTemplate `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res.StandardChangeTemplates, resp, nil
}

// Get a single standard change template
func (s *StandardChangeTemplatesService) Get(ctx context.Context, number string, opts GetOptions) (*StandardChangeTemplate, *Response, error) {
	u := fmt.Sprint("/std_change_proposal.do")
	if number == "" {
		return nil, nil, errors.New("standard change template number cannot be empty")
	}
	opts.internalFields.SysparmQuery = fmt.Sprintf("%s=%s", "number", number)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		StandardChangeTemplates []*StandardChangeTemplate `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	templates := &StandardChangeTemplate{}
	if len(res.StandardChangeTemplates) > 0 {
		templates = res.StandardChangeTemplates[0]
	}

	return templates, resp, nil
}

// Create a standard change template
func (s *StandardChangeTemplatesService) Create(ctx context.Context, template *StandardChangeTemplate, opts CreateOptions) (*StandardChangeTemplate, *Response, error) {
	u := fmt.Sprint("/std_change_proposal.do")
	opts.internalFields.SysparmAction = SysparmActionInsert
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, template)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		StandardChangeTemplates []*StandardChangeTemplate `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	resTemplate := &StandardChangeTemplate{}
	if len(res.StandardChangeTemplates) > 0 {
		resTemplate = res.StandardChangeTemplates[0]
	}

	return resTemplate, resp, nil
}

// Update an existing standard change template on the specified CMDB CI.
func (s *StandardChangeTemplatesService) Update(ctx context.Context, number string, template *StandardChangeTemplate, opts UpdateOptions) (*StandardChangeTemplate, *Response, error) {
	u := fmt.Sprint("/std_change_proposal.do")
	if number == "" {
		return nil, nil, errors.New("standard change template number cannot be empty")
	}
	opts.internalFields.SysparmQuery = fmt.Sprintf("%s=%s", "number", number)
	opts.internalFields.SysparmAction = SysparmActionUpdate
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, template)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		StandardChangeTemplates []*StandardChangeTemplate `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	resTemplate := &StandardChangeTemplate{}
	if len(res.StandardChangeTemplates) > 0 {
		resTemplate = res.StandardChangeTemplates[0]
	}

	return resTemplate, resp, nil
}
