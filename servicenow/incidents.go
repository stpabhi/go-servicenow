package servicenow

import (
	"context"
	"errors"
	"fmt"
)

// IncidentsService handles communication with the Incident related
// methods of the ServiceNow API.
type IncidentsService service

// Incident represents a ServiceNow incident.
type Incident struct {
	Status                 *string `json:"__status,omitempty"`
	UponApproval           *string `json:"upon_approval,omitempty"`
	Location               *string `json:"location,omitempty"`
	ExpectedStart          *string `json:"expected_start,omitempty"`
	ReopenCount            *string `json:"reopen_count,omitempty"`
	CloseNotes             *string `json:"close_notes,omitempty"`
	AdditionalAssigneeList *string `json:"additional_assignee_list,omitempty"`
	Impact                 *string `json:"impact,omitempty"`
	Urgency                *string `json:"urgency,omitempty"`
	CorrelationID          *string `json:"correlation_id,omitempty"`
	SysTags                *string `json:"sys_tags,omitempty"`
	SysDomain              *string `json:"sys_domain,omitempty"`
	Description            *string `json:"description,omitempty"`
	GroupList              *string `json:"group_list,omitempty"`
	Priority               *string `json:"priority,omitempty"`
	DeliveryPlan           *string `json:"delivery_plan,omitempty"`
	SysModCount            *string `json:"sys_mod_count,omitempty"`
	WorkNotesList          *string `json:"work_notes_list,omitempty"`
	BusinessService        *string `json:"business_service,omitempty"`
	FollowUp               *string `json:"follow_up,omitempty"`
	ClosedAt               *string `json:"closed_at,omitempty"`
	SLADue                 *string `json:"sla_due,omitempty"`
	DeliveryTask           *string `json:"delivery_task,omitempty"`
	SysUpdatedOn           *string `json:"sys_updated_on,omitempty"`
	Parent                 *string `json:"parent,omitempty"`
	WorkEnd                *string `json:"work_end,omitempty"`
	Number                 *string `json:"number,omitempty"`
	ClosedBy               *string `json:"closed_by,omitempty"`
	WorkStart              *string `json:"work_start,omitempty"`
	CalendarStc            *string `json:"calendar_stc,omitempty"`
	Category               *string `json:"category,omitempty"`
	BusinessDuration       *string `json:"business_duration,omitempty"`
	IncidentState          *string `json:"incident_state,omitempty"`
	ActivityDue            *string `json:"activity_due,omitempty"`
	CorrelationDisplay     *string `json:"correlation_display,omitempty"`
	Company                *string `json:"company,omitempty"`
	Active                 *string `json:"active,omitempty"`
	DueDate                *string `json:"due_date,omitempty"`
	AssignmentGroup        *string `json:"assignment_group,omitempty"`
	CallerID               *string `json:"caller_id,omitempty"`
	Knowledge              *string `json:"knowledge,omitempty"`
	MadeSLA                *string `json:"made_sla,omitempty"`
	CommentsAndWorkNotes   *string `json:"comments_and_work_notes,omitempty"`
	ParentIncident         *string `json:"parent_incident,omitempty"`
	State                  *string `json:"state,omitempty"`
	UserInput              *string `json:"user_input,omitempty"`
	SysCreatedOn           *string `json:"sys_created_on,omitempty"`
	ApprovalSet            *string `json:"approval_set,omitempty"`
	ReassignmentCount      *string `json:"reassignment_count,omitempty"`
	Rfc                    *string `json:"rfc,omitempty"`
	ChildIncidents         *string `json:"child_incidents,omitempty"`
	OpenedAt               *string `json:"opened_at,omitempty"`
	ShortDescription       *string `json:"short_description,omitempty"`
	Order                  *string `json:"order,omitempty"`
	SysUpdatedBy           *string `json:"sys_updated_by,omitempty"`
	ResolvedBy             *string `json:"resolved_by,omitempty"`
	Notify                 *string `json:"notify,omitempty"`
	UponReject             *string `json:"upon_reject,omitempty"`
	ApprovalHistory        *string `json:"approval_history,omitempty"`
	ProblemID              *string `json:"problem_id,omitempty"`
	WorkNotes              *string `json:"work_notes,omitempty"`
	CalendarDuration       *string `json:"calendar_duration,omitempty"`
	CloseCode              *string `json:"close_code,omitempty"`
	SysID                  *string `json:"sys_id,omitempty"`
	Approval               *string `json:"approval,omitempty"`
	CausedBy               *string `json:"caused_by,omitempty"`
	Severity               *string `json:"severity,omitempty"`
	SysCreatedBy           *string `json:"sys_created_by,omitempty"`
	ResolvedAt             *string `json:"resolved_at,omitempty"`
	AssignedTo             *string `json:"assigned_to,omitempty"`
	BusinessStc            *string `json:"business_stc,omitempty"`
	WfActivity             *string `json:"wf_activity,omitempty"`
	SysDomainPath          *string `json:"sys_domain_path,omitempty"`
	CmdbCi                 *string `json:"cmdb_ci,omitempty"`
	OpenedBy               *string `json:"opened_by,omitempty"`
	Subcategory            *string `json:"subcategory,omitempty"`
	RejectionGoto          *string `json:"rejection_goto,omitempty"`
	SysClassName           *string `json:"sys_class_name,omitempty"`
	WatchList              *string `json:"watch_list,omitempty"`
	TimeWorked             *string `json:"time_worked,omitempty"`
	ContactType            *string `json:"contact_type,omitempty"`
	Escalation             *string `json:"escalation,omitempty"`
	Comments               *string `json:"comments,omitempty"`
}

func (i Incident) String() string {
	return Stringify(i)
}

// List incidents.
func (s *IncidentsService) List(ctx context.Context, opts ListOptions) ([]*Incident, *Response, error) {
	u := fmt.Sprint("/incident.do")
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		Incidents []*Incident `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res.Incidents, resp, nil
}

// Get a single incident.
func (s *IncidentsService) Get(ctx context.Context, number string, opts GetOptions) (*Incident, *Response, error) {
	u := fmt.Sprint("/incident.do")
	if number == "" {
		return nil, nil, errors.New("incident number cannot be empty")
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
		Incidents []*Incident `json:"records,omitempty"` // Though servicenow docs say they return a record, we get records (array).
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	inc := &Incident{}
	if len(res.Incidents) > 0 {
		inc = res.Incidents[0]
	}

	return inc, resp, nil
}

// Create a new incident on the specified CMDB CI.
func (s *IncidentsService) Create(ctx context.Context, inc *Incident, opts CreateOptions) (*Incident, *Response, error) {
	u := fmt.Sprint("/incident.do")
	opts.internalFields.SysparmAction = SysparmActionInsert
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, inc)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		Incidents []*Incident `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, nil
	}

	resInc := &Incident{}
	if len(res.Incidents) > 0 {
		resInc = res.Incidents[0]
	}

	return resInc, resp, nil
}

// Update an existing incident on the specified CMDB CI.
func (s *IncidentsService) Update(ctx context.Context, number string, inc *Incident, opts UpdateOptions) (*Incident, *Response, error) {
	u := fmt.Sprint("/incident.do")
	if number == "" {
		return nil, nil, errors.New("incident number cannot be empty")
	}
	opts.internalFields.SysparmQuery = fmt.Sprintf("%s=%s", "number", number)
	opts.internalFields.SysparmAction = SysparmActionUpdate
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, inc)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		Incidents []*Incident `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, nil
	}

	resInc := &Incident{}
	if len(res.Incidents) > 0 {
		resInc = res.Incidents[0]
	}

	return resInc, resp, nil
}
