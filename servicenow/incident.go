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
	Active                 *string `json:"active,omitempty"`
	ActivityDue            *string `json:"activity_due,omitempty"`
	AdditionalAssigneeList *string `json:"additional_assignee_list,omitempty"`
	Approval               *string `json:"approval,omitempty"`
	ApprovalHistory        *string `json:"approval_history,omitempty"`
	ApprovalSet            *string `json:"approval_set,omitempty"`
	AssignedTo             *string `json:"assigned_to,omitempty"`
	AssignmentGroup        *string `json:"assignment_group,omitempty"`
	BusinessDuration       *string `json:"business_duration,omitempty"`
	BusinessService        *string `json:"business_service,omitempty"`
	BusinessStc            *string `json:"business_stc,omitempty"`
	CalendarDuration       *string `json:"calendar_duration,omitempty"`
	CalendarStc            *string `json:"calendar_stc,omitempty"`
	CallerID               *string `json:"caller_id,omitempty"`
	Category               *string `json:"category,omitempty"`
	CausedBy               *string `json:"caused_by,omitempty"`
	ChildIncidents         *string `json:"child_incidents,omitempty"`
	CloseCode              *string `json:"close_code,omitempty"`
	ClosedAt               *string `json:"closed_at,omitempty"`
	ClosedBy               *string `json:"closed_by,omitempty"`
	CloseNotes             *string `json:"close_notes,omitempty"`
	CmdbCi                 *string `json:"cmdb_ci,omitempty"`
	Comments               *string `json:"comments,omitempty"`
	CommentsAndWorkNotes   *string `json:"comments_and_work_notes,omitempty"`
	Company                *string `json:"company,omitempty"`
	ContactType            *string `json:"contact_type,omitempty"`
	CorrelationDisplay     *string `json:"correlation_display,omitempty"`
	CorrelationID          *string `json:"correlation_id,omitempty"`
	DeliveryPlan           *string `json:"delivery_plan,omitempty"`
	DeliveryTask           *string `json:"delivery_task,omitempty"`
	Description            *string `json:"description,omitempty"`
	DueDate                *string `json:"due_date,omitempty"`
	Escalation             *string `json:"escalation,omitempty"`
	ExpectedStart          *string `json:"expected_start,omitempty"`
	FollowUp               *string `json:"follow_up,omitempty"`
	GroupList              *string `json:"group_list,omitempty"`
	Impact                 *string `json:"impact,omitempty"`
	IncidentState          *string `json:"incident_state,omitempty"`
	Knowledge              *string `json:"knowledge,omitempty"`
	Location               *string `json:"location,omitempty"`
	MadeSLA                *string `json:"made_sla,omitempty"`
	Notify                 *string `json:"notify,omitempty"`
	Number                 *string `json:"number,omitempty"`
	OpenedAt               *string `json:"opened_at,omitempty"`
	OpenedBy               *string `json:"opened_by,omitempty"`
	Order                  *string `json:"order,omitempty"`
	Parent                 *string `json:"parent,omitempty"`
	ParentIncident         *string `json:"parent_incident,omitempty"`
	Priority               *string `json:"priority,omitempty"`
	ProblemID              *string `json:"problem_id,omitempty"`
	ReassignmentCount      *string `json:"reassignment_count,omitempty"`
	RejectionGoto          *string `json:"rejection_goto,omitempty"`
	ReopenCount            *string `json:"reopen_count,omitempty"`
	ResolvedAt             *string `json:"resolved_at,omitempty"`
	ResolvedBy             *string `json:"resolved_by,omitempty"`
	Rfc                    *string `json:"rfc,omitempty"`
	Severity               *string `json:"severity,omitempty"`
	ShortDescription       *string `json:"short_description,omitempty"`
	SLADue                 *string `json:"sla_due,omitempty"`
	State                  *string `json:"state,omitempty"`
	Subcategory            *string `json:"subcategory,omitempty"`
	SysClassName           *string `json:"sys_class_name,omitempty"`
	SysCreatedBy           *string `json:"sys_created_by,omitempty"`
	SysCreatedOn           *string `json:"sys_created_on,omitempty"`
	SysDomain              *string `json:"sys_domain,omitempty"`
	SysDomainPath          *string `json:"sys_domain_path,omitempty"`
	SysID                  *string `json:"sys_id,omitempty"`
	SysModCount            *string `json:"sys_mod_count,omitempty"`
	SysTags                *string `json:"sys_tags,omitempty"`
	SysUpdatedBy           *string `json:"sys_updated_by,omitempty"`
	SysUpdatedOn           *string `json:"sys_updated_on,omitempty"`
	TimeWorked             *string `json:"time_worked,omitempty"`
	UponApproval           *string `json:"upon_approval,omitempty"`
	UponReject             *string `json:"upon_reject,omitempty"`
	Urgency                *string `json:"urgency,omitempty"`
	UserInput              *string `json:"user_input,omitempty"`
	WatchList              *string `json:"watch_list,omitempty"`
	WfActivity             *string `json:"wf_activity,omitempty"`
	WorkEnd                *string `json:"work_end,omitempty"`
	WorkNotes              *string `json:"work_notes,omitempty"`
	WorkNotesList          *string `json:"work_notes_list,omitempty"`
	WorkStart              *string `json:"work_start,omitempty"`
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
