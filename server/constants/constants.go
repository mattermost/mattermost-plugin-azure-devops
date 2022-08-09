package constants

const (
	// Bot configs
	BotUsername    = "azuredevops"
	BotDisplayName = "Azure Devops"
	BotDescription = "A bot account created by the Azure Devops plugin."

	// Plugin configs
	PluginID               = "mattermost-plugin-azure-devops"
	ChannelID              = "channel_id"
	HeaderMattermostUserID = "Mattermost-User-ID"
	// TODO: Change later according to the needs.
	HeaderMattermostUserIDAPI = "User-ID"

	// Command configs
	CommandTriggerName = "azuredevops"
	HelpText           = "###### Mattermost Azure Devops Plugin - Slash Command Help\n"
	InvalidCommand     = "Invalid command parameters. Please use `/azuredevops help` for more information."

	// Azure API Routes
	// TODO: WIP.
	// GetProjects = "/%s/_apis/projects"
	// GetTasksID = "/%s/_apis/wit/wiql"
	// GetTasks   = "/%s/_apis/wit/workitems"
	CreateTask = "/%s/%s/_apis/wit/workitems/$%s?api-version=7.1-preview.3"
	GetTask    = "%s/_apis/wit/workitems/%s?api-version=7.1-preview.3"

	// Azure API versions
	// TODO: WIP.
	// ProjectAPIVersion = "7.1-preview.4"
	// TasksIDAPIVersion    = "5.1"
	// TasksAPIVersion      = "6.0"
	// CreateTaskAPIVersion = "7.1-preview.3"

	// Get task link preview constants
	HTTPS              = "https:"
	HTTP               = "http:"
	AzureDevopsBaseURL = "dev.azure.com"
	Workitems          = "_workitems"
	Edit               = "edit"

	// Azure API Versions
	CreateTaskAPIVersion = "7.1-preview.3"

	// Authorization constants
	Bearer        = "Bearer"
	Authorization = "Authorization"

	GetTasksID = "/%s/_apis/wit/wiql"
	GetTasks   = "/%s/_apis/wit/workitems"

	TasksIDAPIVersion = "5.1"
	TasksAPIVersion   = "6.0"

	PageQueryParam       = "$top"
	APIVersionQueryParam = "api-version"
	IDsQueryParam        = "ids"
)
