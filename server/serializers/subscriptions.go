package serializers

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/Brightscout/mattermost-plugin-azure-devops/server/constants"
)

type UserID struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	UniqueName  string `json:"uniqueName"`
}

type PublisherInputsGeneric struct {
	ProjectID string `json:"projectId"`
}

type PublisherInputsBoards struct {
	ProjectID string `json:"projectId"`
}

type PublisherInputsRepos struct {
	ProjectID                    string `json:"projectId"`
	Repository                   string `json:"repository"`
	Branch                       string `json:"branch"`
	PullRequestCreatedBy         string `json:"pullrequestCreatedBy"`
	PullRequestReviewersContains string `json:"pullrequestReviewersContains"`
}

type ConsumerInputs struct {
	URL string `json:"url"`
}

type SubscriptionValue struct {
	ID               string         `json:"id"`
	URL              string         `json:"url"`
	EventType        string         `json:"eventType"`
	ServiceType      string         `json:"serviceType"`
	ConsumerID       string         `json:"consumerId"`
	ConsumerActionID string         `json:"consumerActionId"`
	CreatedBy        UserID         `json:"createdBy"`
	ModifiedBy       UserID         `json:"modifiedBy"`
	PublisherInputs  interface{}    `json:"publisherInputs"`
	ConsumerInputs   ConsumerInputs `json:"consumerInputs"`
}

type SubscriptionList struct {
	Count             int                 `json:"count"`
	SubscriptionValue []SubscriptionValue `json:"value"`
}
type CreateSubscriptionRequestPayload struct {
	Organization                 string `json:"organization"`
	Project                      string `json:"project"`
	EventType                    string `json:"eventType"`
	ServiceType                  string `json:"serviceType"`
	ChannelID                    string `json:"channelID"`
	Repository                   string `json:"repository"`
	RepositoryName               string `json:"repositoryName"`
	TargetBranch                 string `json:"targetBranch"`
	PullRequestCreatedBy         string `json:"pullrequestCreatedBy"`
	PullRequestReviewersContains string `json:"pullrequestReviewersContains"`
}

type GetSubscriptionFilterPossibleValuesRequestPayload struct {
	Organization string   `json:"organization"`
	ProjectID    string   `json:"projectId"`
	EventType    string   `json:"eventType"`
	Filters      []string `json:"filters"`
	RepositoryID string   `json:"repositoryId"`
}

type SubscriptionFilter struct {
	InputID string `json:"inputId"`
}

type GetSubscriptionFilterValuesRequestPayloadFromClient struct {
	Subscription *CreateSubscriptionBodyPayload `json:"subscription"`
	InputValues  []SubscriptionFilter           `json:"inputValues"`
	Scope        int                            `json:"scope"`
}

type PossibleValues struct {
	DisplayValue string `json:"displayValue"`
	Value        string `json:"value"`
}

type InputValues struct {
	SubscriptionFilter
	PossibleValues []PossibleValues `json:"possibleValues"`
}

type SubscriptionFilterPossibleValuesResponseFromClient struct {
	InputValues []InputValues `json:"inputValues"`
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
	MattermostUserID             string `json:"mattermostUserID"`
	ProjectName                  string `json:"projectName"`
	ProjectID                    string `json:"projectID"`
	OrganizationName             string `json:"organizationName"`
	EventType                    string `json:"eventType"`
	ServiceType                  string `json:"serviceType"`
	ChannelID                    string `json:"channelID"`
	ChannelName                  string `json:"channelName"`
	ChannelType                  string `json:"channelType"`
	SubscriptionID               string `json:"subscriptionID"`
	CreatedBy                    string `json:"createdBy"`
	TargetBranch                 string `json:"targetBranch"`
	Repository                   string `json:"repository"`
	RepositoryName               string `json:"repositoryName"`
	PullRequestCreatedBy         string `json:"pullrequestCreatedBy"`
	PullRequestReviewersContains string `json:"pullrequestReviewersContains"`
}

type DetailedMessage struct {
	Markdown string `json:"markdown"`
}

type SubscriptionNotification struct {
	DetailedMessage DetailedMessage `json:"detailedMessage"`
	Message         DetailedMessage `json:"Message"`
	EventType       string          `json:"eventType"`
	Resource        Resource        `json:"resource"`
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
	Comment       Comment      `json:"comment"`
	PullRequest   PullRequest  `json:"pullRequest"`
	Commits       []Commit     `json:"commits"`
	RefUpdates    []RefUpdates `json:"refUpdates"`
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
	PullRequestCreatedBy         string `json:"pullrequestCreatedBy"`
	PullRequestReviewersContains string `json:"pullrequestReviewersContains"`
}

func GetSubscriptionFilterPossibleValuesRequestPayloadFromJSON(data io.Reader) (*GetSubscriptionFilterPossibleValuesRequestPayload, error) {
	var body *GetSubscriptionFilterPossibleValuesRequestPayload
	if err := json.NewDecoder(data).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
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
