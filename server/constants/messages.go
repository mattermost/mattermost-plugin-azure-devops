package constants

const (
	// Generic
	GenericErrorMessage            = "Something went wrong, please try again later"
	SessionExpiredMessage          = "Session expired. Please connect your Azure DevOps account again"
	ConnectAccount                 = "[Click here to connect your Azure DevOps account](%s%s)"
	ConnectAccountFirst            = "Your Azure DevOps account is not connected \n%s"
	UserConnected                  = "Your Azure DevOps account is successfully connected!"
	MattermostUserAlreadyConnected = "Your Azure DevOps account is already connected"
	UserDisconnected               = "Your Azure DevOps account is now disconnected"
	CreatedTask                    = "Work item [#%d: \"%s\"](%s) of type \"%s\" was successfully created by %s."
	TaskTitle                      = "[%s #%d: %s](%s)"
	PullRequestTitle               = "[#%d: %s](%s)"
	BuildDetailsTitle              = "[#%s](%s): %s"
	PipelineDetailsTitle           = "[%s](%s): %s"
	AlreadyLinkedProject           = "This project is already linked."
	NoProjectLinked                = "No project is linked, please link a project."
	PipelinesRequestBeingProcessed = "Your approval/rejection request is being processed."
	PipelinesRequestProcessed      = "Your approval/rejection request is processed."

	// Validations Errors
	OrganizationRequired            = "organization is required"
	ProjectRequired                 = "project is required"
	TaskTypeRequired                = "task type is required"
	TaskTitleRequired               = "task title is required"
	EventTypeRequired               = "event type is required"
	ServiceTypeRequired             = "service type is required"
	ChannelIDRequired               = "channel ID is required"
	WebhookSecretRequired           = "webhook secret is required"
	MMUserIDRequired                = "mattermsot user ID is required"
	EmptyAzureDevopsAPIBaseURLError = "azure devops API base URL should not be empty"
	EmptyAzureDevopsOAuthAppIDError = "azure devops OAuth app id should not be empty"

	// #nosec G101 -- This is a false positive. The below line is not a hardcoded credential
	EmptyAzureDevopsOAuthClientSecretError = "azure devops OAuth client secret should not be empty"
	EmptyEncryptionSecretError             = "encryption secret should not be empty"
	ProjectIDRequired                      = "project ID is required"
	FiltersRequired                        = "filters required"
)

const (
	// Error messages
	Error                                          = "Error"
	NotAuthorized                                  = "Not authorized"
	UnableToDisconnectUser                         = "Unable to disconnect user"
	UnableToCheckIfAlreadyConnected                = "Unable to check if user account is already connected"
	UnableToStoreOauthState                        = "Unable to store oAuth state for the userID %s"
	UnableToCompleteOAuth                          = "Unable to complete oAuth"
	AuthAttemptExpired                             = "Authentication attempt expired, please try again"
	InvalidAuthState                               = "Invalid oauth state, please try again"
	GetProjectListError                            = "Error in getting project list"
	ErrorFetchProjectList                          = "Error in fetching project list"
	ErrorDecodingBody                              = "Error in decoding body"
	ErrorCreateTask                                = "Error in creating task"
	ErrorCreateSubscription                        = "Error in creating subscription"
	ErrorLinkProject                               = "Error in linking the project"
	FetchSubscriptionListError                     = "Error in fetching subscription list"
	FetchFilteredSubscriptionListError             = "Error in fetching filtered subscription list"
	CreateSubscriptionError                        = "Error in creating subscription"
	ErrorCheckingProjectAdmin                      = "Error in checking if user is an admin on the project %s"
	ProjectNotLinked                               = "Requested project is not linked"
	GetSubscriptionListError                       = "Error getting subscription list"
	SubscriptionAlreadyPresent                     = "Requested subscription already exists"
	SubscriptionNotFound                           = "Requested subscription does not exists"
	ErrorLoadingUserData                           = "Error in loading user data"
	ErrorLoadingDataFromKVStore                    = "Error in loading data from KV store"
	ProjectNotFound                                = "Requested project does not exist"
	ErrorUnlinkProject                             = "Error in unlinking the project"
	InvalidChannelID                               = "Invalid channel ID"
	DeleteSubscriptionError                        = "Error in deleting subscription"
	GetChannelError                                = "Error in getting channels for team and user"
	GetUserError                                   = "Error in getting Mattermost user details"
	InvalidPaginationQueryParam                    = "Invalid value for query param(s) page or per_page"
	ErrorAdminAccess                               = "Cannot delete the subscription, looks like you do not have access to add/delete a subscription for this project. Please make sure you are a project or team administrator for this project"
	ErrorFetchSubscriptionList                     = "Error in fetching subscription list"
	ErrorMessageForAdmin                           = "There is no registered handler for the service hooks event type %s"
	AccessDenied                                   = "Access Denied"
	ErrorOrganizationOrProjectQueryParam           = "Invalid organization or project name"
	ErrorRepositoryPathParam                       = "Invalid organization, project or repository params"
	ErrorInvalidOrganizationOrProject              = "Invalid organization or project name"
	ErrorUpdatingPipelineApprovalRequest           = "Failed to update pipeline approval request"
	ErrorUpdatingNonPendingPipelineRequest         = "Approval(s) %d are not in a pending state. Only pending approval(s) can be updated"
	UnableToDMBot                                  = "Unable to send DM to bot"
	ErrorFetchSubscriptionFilterPossibleValues     = "Error in fetching subscription filter possible values"
	ErrorUnauthorisedSubscriptionsWebhookRequest   = "missing or invalid webhook secret for subscriptions notification"
	ErrorMessageAzureDevopsAccountAlreadyConnected = "azure devops account for %s is already connected"
)
