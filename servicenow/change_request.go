package servicenow

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// ChangeRequestsService handles the communication with the ChangeRequest related
// methods of the ServiceNow API.
type ChangeRequestsService service

// ChangeRequest represents a ServiceNow change.
type ChangeRequest struct {
	Status                         *string `json:"__status,omitempty"`
	Active                         *string `json:"active,omitempty"`
	ActivityDue                    *string `json:"activity_due,omitempty"`
	AdditionalAssigneeList         *string `json:"additional_assignee_list,omitempty"`
	Approval                       *string `json:"approval,omitempty"`
	ApprovalHistory                *string `json:"approval_history,omitempty"`
	ApprovalSet                    *string `json:"approval_set,omitempty"`
	AssignedTo                     *string `json:"assigned_to,omitempty"`
	AssignmentGroup                *string `json:"assignment_group,omitempty"`
	BackoutPlan                    *string `json:"backout_plan,omitempty"`
	BusinessDuration               *string `json:"business_duration,omitempty"`
	BusinessService                *string `json:"business_service,omitempty"`
	CabDate                        *string `json:"cab_date,omitempty"`
	CabDelegate                    *string `json:"cab_delegate,omitempty"`
	CabRecommendation              *string `json:"cab_recommendation,omitempty"`
	CabRequired                    *string `json:"cab_required,omitempty"`
	CalendarDuration               *string `json:"calendar_duration,omitempty"`
	Category                       *string `json:"category,omitempty"`
	ChangePlan                     *string `json:"change_plan,omitempty"`
	ChgModel                       *string `json:"chg_model,omitempty"`
	CloseCode                      *string `json:"close_code,omitempty"`
	CloseNotes                     *string `json:"close_notes,omitempty"`
	ClosedAt                       *string `json:"closed_at,omitempty"`
	ClosedBy                       *string `json:"closed_by,omitempty"`
	CmdbCi                         *string `json:"cmdb_ci,omitempty"`
	Comments                       *string `json:"comments,omitempty"`
	CommentsAndWorkNotes           *string `json:"comments_and_work_notes,omitempty"`
	Company                        *string `json:"company,omitempty"`
	ConflictLastRun                *string `json:"conflict_last_run,omitempty"`
	ConflictStatus                 *string `json:"conflict_status,omitempty"`
	ContactType                    *string `json:"contact_type,omitempty"`
	CorrelationDisplay             *string `json:"correlation_display,omitempty"`
	CorrelationID                  *string `json:"correlation_id,omitempty"`
	Description                    *string `json:"description,omitempty"`
	DueDate                        *string `json:"due_date,omitempty"`
	EndDate                        *string `json:"end_date,omitempty"`
	Escalation                     *string `json:"escalation,omitempty"`
	ExpectedStart                  *string `json:"expected_start,omitempty"`
	FollowUp                       *string `json:"follow_up,omitempty"`
	GroupList                      *string `json:"group_list,omitempty"`
	Impact                         *string `json:"impact,omitempty"`
	ImplementationPlan             *string `json:"implementation_plan,omitempty"`
	Justification                  *string `json:"justification,omitempty"`
	Knowledge                      *string `json:"knowledge,omitempty"`
	Location                       *string `json:"location,omitempty"`
	MadeSLA                        *string `json:"made_sla,omitempty"`
	Number                         *string `json:"number,omitempty"`
	OnHold                         *string `json:"on_hold,omitempty"`
	OnHoldReason                   *string `json:"on_hold_reason,omitempty"`
	OnHoldTask                     *string `json:"on_hold_task,omitempty"`
	OpenedAt                       *string `json:"opened_at,omitempty"`
	OpenedBy                       *string `json:"opened_by,omitempty"`
	Order                          *string `json:"order,omitempty"`
	OutsideMaintenanceSchedule     *string `json:"outside_maintenance_schedule,omitempty"`
	Parent                         *string `json:"parent,omitempty"`
	Phase                          *string `json:"phase,omitempty"`
	PhaseState                     *string `json:"phase_state,omitempty"`
	Priority                       *string `json:"priority,omitempty"`
	ProductionSystem               *string `json:"production_system,omitempty"`
	Reason                         *string `json:"reason,omitempty"`
	ReassignmentCount              *string `json:"reassignment_count,omitempty"`
	RequestedBy                    *string `json:"requested_by,omitempty"`
	RequestedByDate                *string `json:"requested_by_date,omitempty"`
	ReviewComments                 *string `json:"review_comments,omitempty"`
	ReviewDate                     *string `json:"review_date,omitempty"`
	ReviewStatus                   *string `json:"review_status,omitempty"`
	Risk                           *string `json:"risk,omitempty"`
	RiskImpactAnalysis             *string `json:"risk_impact_analysis,omitempty"`
	RiskValue                      *string `json:"risk_value,omitempty"`
	RouteReason                    *string `json:"route_reason,omitempty"`
	Scope                          *string `json:"scope,omitempty"`
	ServiceOffering                *string `json:"service_offering,omitempty"`
	ShortDescription               *string `json:"short_description,omitempty"`
	Skills                         *string `json:"skills,omitempty"`
	SLADue                         *string `json:"sla_due,omitempty"`
	SnEsignDocument                *string `json:"sn_esign_document,omitempty"`
	SnEsignEsignatureConfiguration *string `json:"sn_esign_esignature_configuration,omitempty"`
	StartDate                      *string `json:"start_date,omitempty"`
	State                          *string `json:"state,omitempty"`
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
	TestPlan                       *string `json:"test_plan,omitempty"`
	TimeWorked                     *string `json:"time_worked,omitempty"`
	Type                           *string `json:"type,omitempty"`
	Unauthorized                   *string `json:"unauthorized,omitempty"`
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

	Extra *map[string]interface{} `json:"-"`
}

func (c ChangeRequest) String() string {
	return Stringify(c)
}

// List change requests
func (s *ChangeRequestsService) List(ctx context.Context, opts ListOptions) ([]*ChangeRequest, *Response, error) {
	u := fmt.Sprintf("/change_request.do")
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
		ChangeRequests []*ChangeRequest `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	return res.ChangeRequests, resp, nil
}

// Get a single Change Request.
func (s *ChangeRequestsService) Get(ctx context.Context, number string, opts GetOptions) (*ChangeRequest, *Response, error) {
	u := fmt.Sprintf("/change_request.do")
	if number == "" {
		return nil, nil, fmt.Errorf("change number cannot be empty")
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
		ChangeRequests []*ChangeRequest `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, err
	}

	chg := &ChangeRequest{}
	if len(res.ChangeRequests) > 0 {
		chg = res.ChangeRequests[0]
	}

	return chg, resp, nil
}

// Create a new change request on the specified CMDB CI.
func (s *ChangeRequestsService) Create(ctx context.Context, chg *ChangeRequest, opts CreateOptions) (*ChangeRequest, *Response, error) {
	u := fmt.Sprintf("/change_request.do")
	opts.internalFields.SysparmAction = SysparmActionInsert
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, chg)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		ChangeRequests []*ChangeRequest `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, nil
	}

	resChg := &ChangeRequest{}
	if len(res.ChangeRequests) > 0 {
		resChg = res.ChangeRequests[0]
	}

	return resChg, resp, nil
}

// Update an existing change request on the specified CMDB CI.
func (s *ChangeRequestsService) Update(ctx context.Context, number string, chg *ChangeRequest, opts UpdateOptions) (*ChangeRequest, *Response, error) {
	u := fmt.Sprint("/change_request.do")
	if number == "" {
		return nil, nil, errors.New("change request number cannot be empty")
	}
	opts.internalFields.SysparmQuery = fmt.Sprintf("%s=%s", "number", number)
	opts.internalFields.SysparmAction = SysparmActionUpdate
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, chg)
	if err != nil {
		return nil, nil, err
	}

	var res struct {
		ChangeRequests []*ChangeRequest `json:"records,omitempty"`
	}
	resp, err := s.client.Do(ctx, req, &res)
	if err != nil {
		return nil, resp, nil
	}

	resChg := &ChangeRequest{}
	if len(res.ChangeRequests) > 0 {
		resChg = res.ChangeRequests[0]
	}

	return resChg, resp, nil
}
