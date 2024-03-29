package serializers

import (
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/mattermost/mattermost-plugin-azure-devops/server/constants"
)

type UserID struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	UniqueName  string `json:"uniqueName"`
}

type PublisherInputsGeneric struct {
	ProjectID                    string `json:"projectId,omitempty"`
	AreaPath                     string `json:"areaPath,omitempty"`
	Repository                   string `json:"repository,omitempty"`
	Branch                       string `json:"branch,omitempty"`
	PullRequestCreatedBy         string `json:"pullrequestCreatedBy,omitempty"`
	PullRequestReviewersContains string `json:"pullrequestReviewersContains,omitempty"`
	PushedBy                     string `json:"pushedBy,omitempty"`
	MergeResult                  string `json:"mergeResult,omitempty"`
	NotificationType             string `json:"notificationType,omitempty"`
	DefinitionName               string `json:"definitionName,omitempty"`
	BuildStatus                  string `json:"buildStatus,omitempty"`
	ReleaseDefinitionID          string `json:"releaseDefinitionId,omitempty"`
	ReleaseEnvironmentID         string `json:"releaseEnvironmentId,omitempty"`
	ReleaseApprovalType          string `json:"releaseApprovalType,omitempty"`
	ReleaseApprovalStatus        string `json:"releaseApprovalStatus,omitempty"`
	ReleaseEnvironmentStatus     string `json:"releaseEnvironmentStatus,omitempty"`
	PipelineID                   string `json:"pipelineId,omitempty"`
	StageName                    string `json:"stageName,omitempty"`
	EnvironmentName              string `json:"environmentName,omitempty"`
	StageNameID                  string `json:"stageNameId,omitempty"`
	StageStateID                 string `json:"stageStateId,omitempty"`
	StageResultID                string `json:"stageResultId,omitempty"`
	RunStateID                   string `json:"runStateId,omitempty"`
	RunResultID                  string `json:"runResultId,omitempty"`
}

type ConsumerInputs struct {
	URL string `json:"url"`
}

type SubscriptionValue struct {
	ID               string      `json:"id"`
	URL              string      `json:"url"`
	EventType        string      `json:"eventType"`
	ServiceType      string      `json:"serviceType"`
	ConsumerID       string      `json:"consumerId"`
	ConsumerActionID string      `json:"consumerActionId"`
	CreatedBy        UserID      `json:"createdBy"`
	ModifiedBy       UserID      `json:"modifiedBy"`
	PublisherInputs  interface{} `json:"publisherInputs"`
}

type SubscriptionList struct {
	Count             int                 `json:"count"`
	SubscriptionValue []SubscriptionValue `json:"value"`
}

type CreateSubscriptionRequestPayload struct {
	Organization                     string `json:"organization"`
	Project                          string `json:"project"`
	EventType                        string `json:"eventType"`
	ServiceType                      string `json:"serviceType"`
	ChannelID                        string `json:"channelID"`
	Repository                       string `json:"repository"`
	RepositoryName                   string `json:"repositoryName"`
	TargetBranch                     string `json:"targetBranch"`
	PullRequestCreatedBy             string `json:"pullRequestCreatedBy"`
	PullRequestReviewersContains     string `json:"pullRequestReviewersContains"`
	PullRequestCreatedByName         string `json:"pullRequestCreatedByName"`
	PullRequestReviewersContainsName string `json:"pullRequestReviewersContainsName"`
	PushedBy                         string `json:"pushedBy"`
	PushedByName                     string `json:"pushedByName"`
	MergeResult                      string `json:"mergeResult"`
	MergeResultName                  string `json:"mergeResultName"`
	NotificationType                 string `json:"notificationType"`
	NotificationTypeName             string `json:"notificationTypeName"`
	AreaPath                         string `json:"areaPath"`
	BuildPipeline                    string `json:"buildPipeline"`
	BuildStatus                      string `json:"buildStatus"`
	BuildStatusName                  string `json:"buildStatusName"`
	ReleasePipeline                  string `json:"releasePipeline"`
	ReleasePipelineName              string `json:"releasePipelineName"`
	StageName                        string `json:"stageName"`
	StageNameValue                   string `json:"stageNameValue"`
	ApprovalType                     string `json:"approvalType"`
	ApprovalTypeName                 string `json:"approvalTypeName"`
	ApprovalStatus                   string `json:"approvalStatus"`
	ApprovalStatusName               string `json:"approvalStatusName"`
	ReleaseStatus                    string `json:"releaseStatus"`
	ReleaseStatusName                string `json:"releaseStatusName"`
	RunPipeline                      string `json:"runPipeline"`
	RunPipelineName                  string `json:"runPipelineName"`
	RunStageName                     string `json:"runStage"`
	RunEnvironmentName               string `json:"runEnvironment"`
	RunStageNameID                   string `json:"runStageId"`
	RunStageStateID                  string `json:"runStageStateId"`
	RunStageStateIDName              string `json:"runStageStateIdName"`
	RunStageResultID                 string `json:"runStageResultId"`
	RunStateID                       string `json:"runStateId"`
	RunStateIDName                   string `json:"runStateIdName"`
	RunResultID                      string `json:"runResultId"`
}

type GetSubscriptionFilterPossibleValuesRequestPayload struct {
	Organization      string   `json:"organization"`
	ProjectID         string   `json:"projectId"`
	EventType         string   `json:"eventType"`
	Filters           []string `json:"filters"`
	RepositoryID      string   `json:"repositoryId"`
	ReleasePipelineID string   `json:"releasePipelineId"`
	RunPipeline       string   `json:"runPipeline"`
}

type SubscriptionFilter struct {
	InputID string `json:"inputId"`
}

type GetSubscriptionFilterValuesRequestPayloadFromClient struct {
	Subscription *CreateSubscriptionBodyPayload `json:"subscription"`
	InputValues  []*SubscriptionFilter          `json:"inputValues"`
	Scope        int                            `json:"scope"`
}

type PossibleValues struct {
	DisplayValue string `json:"displayValue"`
	Value        string `json:"value"`
}

type InputValues struct {
	SubscriptionFilter
	PossibleValues []*PossibleValues `json:"possibleValues"`
}

type SubscriptionFilterPossibleValuesResponseFromClient struct {
	InputValues []*InputValues `json:"inputValues"`
}

type CreateSubscriptionBodyPayload struct {
	PublisherID      string         `json:"publisherId"`
	EventType        string         `json:"eventType"`
	ConsumerID       string         `json:"consumerId"`
	ConsumerActionID string         `json:"consumerActionId"`
	PublisherInputs  interface{}    `json:"publisherInputs"`
	ConsumerInputs   ConsumerInputs `json:"consumerInputs"`
}

type SubscriptionDetails struct {
	SubscriptionID   string    `json:"subscriptionID"`
	MattermostUserID string    `json:"mattermostUserID"`
	ProjectName      string    `json:"projectName"`
	ProjectID        string    `json:"projectID"`
	OrganizationName string    `json:"organizationName"`
	EventType        string    `json:"eventType"`
	ServiceType      string    `json:"serviceType"`
	ChannelID        string    `json:"channelID"`
	ChannelName      string    `json:"channelName"`
	ChannelType      string    `json:"channelType"`
	CreatedBy        string    `json:"createdBy"`
	CreatedAt        time.Time `json:"createdAt"`
	// Below all are filters that could be present on different categories of subscriptions from Boards, Repos and Pipelines
	TargetBranch                     string `json:"targetBranch"`
	Repository                       string `json:"repository"`
	RepositoryName                   string `json:"repositoryName"`
	PullRequestCreatedBy             string `json:"pullRequestCreatedBy"`
	PullRequestReviewersContains     string `json:"pullRequestReviewersContains"`
	PullRequestCreatedByName         string `json:"pullRequestCreatedByName"`
	PullRequestReviewersContainsName string `json:"pullRequestReviewersContainsName"`
	PushedBy                         string `json:"pushedBy"`
	PushedByName                     string `json:"pushedByName"`
	MergeResult                      string `json:"mergeResult"`
	MergeResultName                  string `json:"mergeResultName"`
	NotificationType                 string `json:"notificationType"`
	NotificationTypeName             string `json:"notificationTypeName"`
	AreaPath                         string `json:"areaPath"`
	BuildPipeline                    string `json:"buildPipeline"`
	BuildStatus                      string `json:"buildStatus"`
	BuildStatusName                  string `json:"buildStatusName"`
	ReleasePipeline                  string `json:"releasePipeline"`
	ReleasePipelineName              string `json:"releasePipelineName"`
	StageName                        string `json:"stageName"`
	StageNameValue                   string `json:"stageNameValue"`
	ApprovalType                     string `json:"approvalType"`
	ApprovalTypeName                 string `json:"approvalTypeName"`
	ApprovalStatus                   string `json:"approvalStatus"`
	ApprovalStatusName               string `json:"approvalStatusName"`
	ReleaseStatus                    string `json:"releaseStatus"`
	ReleaseStatusName                string `json:"releaseStatusName"`
	RunPipeline                      string `json:"runPipeline"`
	RunPipelineName                  string `json:"runPipelineName"`
	RunStageName                     string `json:"runStage"`
	RunEnvironmentName               string `json:"runEnvironment"`
	RunStageNameID                   string `json:"runStageId"`
	RunStageStateID                  string `json:"runStageStateId"`
	RunStageStateIDName              string `json:"runStageStateIdName"`
	RunStageResultID                 string `json:"runStageResultId"`
	RunStateID                       string `json:"runStateId"`
	RunStateIDName                   string `json:"runStateIdName"`
	RunResultID                      string `json:"runResultId"`
}

type DetailedMessage struct {
	Markdown string `json:"markdown"`
}

type SubscriptionNotification struct {
	SubscriptionID  string          `json:"subscriptionID"`
	DetailedMessage DetailedMessage `json:"detailedMessage"`
	Message         DetailedMessage `json:"message"`
	EventType       string          `json:"eventType"`
	Resource        Resource        `json:"resource"`
}

type Approval struct {
	ID                   interface{}     `json:"id"`
	Approver             Approver        `json:"approver"`
	Steps                []*ApprovalStep `json:"steps"`
	MinRequiredApprovers int             `json:"minRequiredApprovers"`
	ExecutionOrder       string          `json:"executionOrder"`
}

type ApprovalStep struct {
	AssignedApprover Approver `json:"assignedApprover"`
	Status           string   `json:"status"`
	Order            int      `json:"order"`
	ID               int      `json:"id"`
	Approver         Approver `json:"approver"`
}

type Approver struct {
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
}

type Resource struct {
	PullRequestID int          `json:"pullRequestId"`
	Reviewers     []Reviewer   `json:"reviewers"`
	SourceRefName string       `json:"sourceRefName"`
	TargetRefName string       `json:"targetRefName"`
	MergeStatus   string       `json:"mergeStatus"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	Repository    Repository   `json:"repository"`
	Comment       interface{}  `json:"comment"`
	PullRequest   PullRequest  `json:"pullRequest"`
	Commits       []Commit     `json:"commits"`
	RefUpdates    []RefUpdates `json:"refUpdates"`
	Definition    Definition   `json:"definition"`
	SourceBranch  string       `json:"sourceBranch"`
	Project       Project      `json:"project"`
	RequestedFor  RequestedFor `json:"requestedFor"`
	StartTime     string       `json:"startTime"`
	FinishTime    string       `json:"finishTime"`
	Release       Release      `json:"release"`
	StageName     string       `json:"stageName"`
	Environment   Environment  `json:"environment"`
	Stage         Stage        `json:"stage"`
	Pipeline      Definition   `json:"pipeline"`
	Run           Stage        `json:"run"`
	Approval      Approval     `json:"approval"`
	ProjectID     string       `json:"projectId"`
	Fields        Fields       `json:"fields"`
	Revision      Revision     `json:"revision"`
}

type Stage struct {
	Name  string      `json:"name"`
	Links ProjectLink `json:"_links"`
}

type Release struct {
	Name              string      `json:"name"`
	CreatedBy         Reviewer    `json:"createdBy"`
	Artifacts         []*Artifact `json:"artifacts"`
	ReleaseDefinition Definition  `json:"releaseDefinition"`
	Reason            string      `json:"reason"`
	ModifiedOn        string      `json:"modifiedOn"`
	ModifiedBy        Reviewer    `json:"modifiedBy"`
	Links             ProjectLink `json:"_links"`
}

type Artifact struct {
	Name string `json:"alias"`
}

type RequestedFor struct {
	Name string `json:"displayName"`
}

type Definition struct {
	Name  string      `json:"name"`
	URL   string      `json:"url"`
	Links ProjectLink `json:"_links"`
}

type Revision struct {
	Fields Fields `json:"fields"`
}

type Fields struct {
	ProjectName  interface{} `json:"System.TeamProject"`
	AreaPath     interface{} `json:"System.AreaPath"`
	State        interface{} `json:"System.State"`
	WorkItemType interface{} `json:"System.WorkItemType"`
	Title        interface{} `json:"System.Title"`
}

type RefUpdates struct {
	Name string `json:"name"`
}

type Commit struct {
	CommitID string `json:"commitId"`
	Comment  string `json:"comment"`
	URL      string `json:"url"`
}

type Repository struct {
	Name string `json:"name"`
}

type PullRequest struct {
	PullRequestID int        `json:"pullRequestId"`
	Reviewers     []Reviewer `json:"reviewers"`
	SourceRefName string     `json:"sourceRefName"`
	TargetRefName string     `json:"targetRefName"`
	MergeStatus   string     `json:"mergeStatus"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Repository    Repository `json:"repository"`
}

type Comment struct {
	Content string `json:"content"`
}

type Reviewer struct {
	DisplayName string `json:"displayName"`
}

type DeleteSubscriptionRequestPayload struct {
	Organization                 string `json:"organization"`
	Project                      string `json:"project"`
	EventType                    string `json:"eventType"`
	ServiceType                  string `json:"serviceType"`
	ChannelID                    string `json:"channelID"`
	MMUserID                     string `json:"mmUserID"`
	TargetBranch                 string `json:"targetBranch"`
	Repository                   string `json:"repository"`
	PullRequestCreatedBy         string `json:"pullRequestCreatedBy"`
	PullRequestReviewersContains string `json:"pullRequestReviewersContains"`
	PushedBy                     string `json:"pushedBy"`
	MergeResult                  string `json:"mergeResult"`
	NotificationType             string `json:"notificationType"`
	AreaPath                     string `json:"areaPath"`
	BuildPipeline                string `json:"buildPipeline"`
	BuildStatus                  string `json:"buildStatus"`
	ReleasePipeline              string `json:"releasePipeline"`
	StageName                    string `json:"stageName"`
	ApprovalType                 string `json:"approvalType"`
	ApprovalStatus               string `json:"approvalStatus"`
	ReleaseStatus                string `json:"releaseStatus"`
	RunPipeline                  string `json:"runPipeline"`
	RunStageName                 string `json:"runStage"`
	RunEnvironmentName           string `json:"runEnvironment"`
	RunStageNameID               string `json:"runStageId"`
	RunStageStateID              string `json:"runStageStateId"`
	RunStageResultID             string `json:"runStageResultId"`
	RunStateID                   string `json:"runStateId"`
	RunResultID                  string `json:"runResultId"`
}

type PipelineRunApprovalDetails struct {
	ID                   string          `json:"id"`
	Status               string          `json:"status"`
	ApprovalSteps        []*ApprovalStep `json:"steps"`
	MinRequiredApprovers int             `json:"minRequiredApprovers"`
}

type PipelineRunApproveResponse struct {
	Value []*PipelineRunResponseValue `json:"value"`
}

type PipelineRunResponseValue struct {
	Status               string          `json:"status"`
	MinRequiredApprovers int             `json:"minRequiredApprovers"`
	ApprovalSteps        []*ApprovalStep `json:"steps"`
}

func GetSubscriptionFilterPossibleValuesRequestPayloadFromJSON(data io.Reader) (*GetSubscriptionFilterPossibleValuesRequestPayload, error) {
	var body *GetSubscriptionFilterPossibleValuesRequestPayload
	if err := json.NewDecoder(data).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

type PipelineApprovalDetails struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type BuildDetails struct {
	BuildNumber  string      `json:"buildNumber"`
	SourceBranch string      `json:"sourceBranch"`
	Repository   Repository  `json:"repository"`
	Status       string      `json:"status"`
	RequestedBy  RequestedBy `json:"requestedBy"`
	Project      Project     `json:"project"`
	Link         Link        `json:"_links"`
	Definition   Definition  `json:"definition"`
}

type RequestedBy struct {
	DisplayName string `json:"displayName"`
}

type ReleaseDetails struct {
	Name              string            `json:"name"`
	ID                int               `json:"id"`
	Status            string            `json:"status"`
	Environments      []*Environment    `json:"environments"`
	Link              Link              `json:"_links"`
	ReleaseDefinition ReleaseDefinition `json:"releaseDefinition"`
}

type Environment struct {
	Name              string     `json:"name"`
	Release           Release    `json:"release"`
	ReleaseDefinition Definition `json:"releaseDefinition"`
}

type ReleaseDefinition struct {
	Name string `json:"name"`
}

func CreateSubscriptionRequestPayloadFromJSON(data io.Reader) (*CreateSubscriptionRequestPayload, error) {
	var body *CreateSubscriptionRequestPayload
	if err := json.NewDecoder(data).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func SubscriptionNotificationFromJSON(data io.Reader) (*SubscriptionNotification, error) {
	var body *SubscriptionNotification
	if err := json.NewDecoder(data).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func DeleteSubscriptionRequestPayloadFromJSON(data io.Reader) (*DeleteSubscriptionRequestPayload, error) {
	var body *DeleteSubscriptionRequestPayload
	if err := json.NewDecoder(data).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}

func (t *GetSubscriptionFilterPossibleValuesRequestPayload) IsSubscriptionRequestPayloadValid() error {
	if t.Organization == "" {
		return errors.New(constants.OrganizationRequired)
	}
	if t.ProjectID == "" {
		return errors.New(constants.ProjectIDRequired)
	}
	if t.EventType == "" {
		return errors.New(constants.EventTypeRequired)
	}
	if t.Filters == nil {
		return errors.New(constants.FiltersRequired)
	}
	return nil
}

func (t *CreateSubscriptionRequestPayload) IsSubscriptionRequestPayloadValid() error {
	if t.Organization == "" {
		return errors.New(constants.OrganizationRequired)
	}
	if t.Project == "" {
		return errors.New(constants.ProjectRequired)
	}
	if t.EventType == "" {
		return errors.New(constants.EventTypeRequired)
	}
	if t.ServiceType == "" {
		return errors.New(constants.ServiceTypeRequired)
	}
	if t.ChannelID == "" {
		return errors.New(constants.ChannelIDRequired)
	}
	return nil
}

func (t *DeleteSubscriptionRequestPayload) IsSubscriptionRequestPayloadValid() error {
	if t.Organization == "" {
		return errors.New(constants.OrganizationRequired)
	}
	if t.Project == "" {
		return errors.New(constants.ProjectRequired)
	}
	if t.EventType == "" {
		return errors.New(constants.EventTypeRequired)
	}
	if t.ChannelID == "" {
		return errors.New(constants.ChannelIDRequired)
	}
	if t.MMUserID == "" {
		return errors.New(constants.MMUserIDRequired)
	}
	return nil
}
