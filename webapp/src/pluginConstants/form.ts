import {subscriptionFilters, filterLabelValuePairAll} from './common';

// Create subscription modal
export const boardEventTypeOptions: LabelValuePair[] = [
    {
        value: 'workitem.created',
        label: 'Work item created',
    },
    {
        value: 'workitem.updated',
        label: 'Work item updated',
    },
    {
        value: 'workitem.deleted',
        label: 'Work item deleted',
    },
    {
        value: 'workitem.commented',
        label: 'Work item commented',
    },
];

export const repoEventTypeOptions: LabelValuePair[] = [
    {
        value: 'git.pullrequest.created',
        label: 'Pull request created',
    },
    {
        value: 'git.pullrequest.updated',
        label: 'Pull request updated',
    },
    {
        value: 'ms.vss-code.git-pullrequest-comment-event',
        label: 'Pull request commented',
    },
    {
        value: 'git.push',
        label: 'Code pushed',
    },
    {
        value: 'git.pullrequest.merged',
        label: 'Pull request merge attempted',
    },
];

export const pipelineEventTypeOptions: LabelValuePair[] = [
    {
        value: 'build.complete',
        label: 'Build completed',
    },
    {
        value: 'ms.vss-release.release-abandoned-event',
        label: 'Release abandoned',
    },
    {
        value: 'ms.vss-release.release-created-event',
        label: 'Release created',
    },
    {
        value: 'ms.vss-release.deployment-approval-completed-event',
        label: 'Release deployment approval completed',
    },
    {
        value: 'ms.vss-release.deployment-approval-pending-event',
        label: 'Release deployment approval pending',
    },
    {
        value: 'ms.vss-release.deployment-completed-event',
        label: 'Release deployment completed',
    },
    {
        value: 'ms.vss-release.deployment-started-event',
        label: 'Release deployment started',
    },
    {
        value: 'ms.vss-pipelinechecks-events.approval-completed',
        label: 'Run stage approval completed',
    },
    {
        value: 'ms.vss-pipelines.stage-state-changed-event',
        label: 'Run stage state changed',
    },
    {
        value: 'ms.vss-pipelinechecks-events.approval-pending',
        label: 'Run stage waiting for approval',
    },
    {
        value: 'ms.vss-pipelines.run-state-changed-event',
        label: 'Run state changed',
    },
];

const serviceTypeOptions: LabelValuePair[] = [
    {
        value: 'boards',
        label: 'Boards',
    },
    {
        value: 'repos',
        label: 'Repos',
    },
    {
        value: 'pipelines',
        label: 'Pipelines',
    },
];

export const buildStatusOptions: LabelValuePair[] = [
    {
        value: 'Succeeded',
        label: 'Succeeded',
    },
    {
        value: 'PartiallySucceeded',
        label: 'Partially Succeeded',
    },
    {
        value: 'Failed',
        label: 'Failed',
    },
    {
        value: 'Stopped',
        label: 'Stopped',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const releaseApprovalTypeOptions: LabelValuePair[] = [
    {
        value: '1',
        label: 'Pre-deployment',
    },
    {
        value: '2',
        label: 'Post-deployment',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const releaseApprovalStatusOptions: LabelValuePair[] = [
    {
        value: '1',
        label: 'Approved',
    },
    {
        value: '2',
        label: 'Rejected',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const releaseStatusOptions: LabelValuePair[] = [
    {
        value: '8',
        label: 'Cancelled',
    },
    {
        value: '4',
        label: 'Succeeded',
    },
    {
        value: '128',
        label: 'Partially Succeeded',
    },
    {
        value: '16',
        label: 'Failed',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const runStageStateIdOptions: LabelValuePair[] = [
    {
        value: 'NotStarted',
        label: 'Not Started',
    },
    {
        value: 'Waiting',
        label: 'Waiting',
    },
    {
        value: 'Running',
        label: 'Running',
    },
    {
        value: 'Completed',
        label: 'Completed',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const runStageResultIdOptions: LabelValuePair[] = [
    {
        value: 'Cancelled',
        label: 'Cancelled',
    },
    {
        value: 'Failed',
        label: 'Failed',
    },
    {
        value: 'Rejected',
        label: 'Rejected',
    },
    {
        value: 'Skipped',
        label: 'Skipped',
    },
    {
        value: 'Succeeded',
        label: 'Succeeded',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const runStateIdOptions: LabelValuePair[] = [
    {
        value: 'InProgress',
        label: 'In Progress',
    },
    {
        value: 'Cancelling',
        label: 'Cancelling',
    },
    {
        value: 'Completed',
        label: 'Completed',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const runResultIdOptions: LabelValuePair[] = [
    {
        value: 'Cancelled',
        label: 'Cancelled',
    },
    {
        value: 'Failed',
        label: 'Failed',
    },
    {
        value: 'Succeeded',
        label: 'Succeeded',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const subscriptionModal: Record<SubscriptionModalFields, ModalFormFieldConfig> = {
    organization: {
        label: 'Organization name',
        type: 'dropdown',
        value: '',
        validations: {
            isRequired: true,
        },
    },
    project: {
        label: 'Project name',
        value: '',
        type: 'dropdown',
        validations: {
            isRequired: true,
        },
    },
    serviceType: {
        label: 'Service type',
        value: 'boards',
        type: 'dropdown',
        optionsList: serviceTypeOptions,
        validations: {
            isRequired: true,
        },
    },
    eventType: {
        label: 'Event type',
        value: boardEventTypeOptions[0].value,
        type: 'dropdown',
        optionsList: boardEventTypeOptions,
        validations: {
            isRequired: true,
        },
    },
    channelID: {
        label: 'Channel name',
        value: '',
        type: 'dropdown',
        validations: {
            isRequired: true,
        },
    },
    repository: {
        label: 'Repository',
        type: 'hidden',
        value: '',
    },
    repositoryName: {
        label: 'Repository Name',
        type: 'hidden',
        value: '',
    },
    targetBranch: {
        label: 'Target branch',
        type: 'hidden',
        value: '',
    },
    pullRequestCreatedBy: {
        label: 'Requested by a member of group',
        type: 'hidden',
        value: '',
    },
    pullRequestReviewersContains: {
        label: 'Reviewer includes group',
        type: 'hidden',
        value: '',
    },
    pullRequestCreatedByName: {
        label: 'Requested by a member of group',
        type: 'hidden',
        value: '',
    },
    pullRequestReviewersContainsName: {
        label: 'Reviewer includes group',
        type: 'hidden',
        value: '',
    },
    pushedBy: {
        label: 'Pushed by a member of group',
        type: 'hidden',
        value: '',
    },
    mergeResult: {
        label: 'Merge Result',
        type: 'hidden',
        value: '',
    },
    notificationType: {
        label: 'Change',
        type: 'hidden',
        value: '',
    },
    pushedByName: {
        label: 'Pushed by a member of group',
        type: 'hidden',
        value: '',
    },
    mergeResultName: {
        label: 'Merge Result',
        type: 'hidden',
        value: '',
    },
    notificationTypeName: {
        label: 'Change',
        type: 'hidden',
        value: '',
    },
    areaPath: {
        label: 'Area Path',
        type: 'hidden',
        value: '',
    },
    buildPipeline: {
        label: 'Build Pipeline',
        type: 'hidden',
        value: '',
    },
    buildStatus: {
        label: 'Build Status',
        type: 'hidden',
        value: '',
        optionsList: buildStatusOptions,
    },
    buildStatusName: {
        label: 'Build Status',
        type: 'hidden',
        value: '',
    },
    releasePipeline: {
        label: 'Release Pipeline Name',
        type: 'hidden',
        value: '',
    },
    releasePipelineName: {
        label: 'Release Pipeline Name',
        type: 'hidden',
        value: '',
    },
    stageName: {
        label: 'Stage Name',
        type: 'hidden',
        value: '',
    },
    stageNameValue: {
        label: 'Stage Name',
        type: 'hidden',
        value: '',
    },
    approvalType: {
        label: 'Approval Type',
        type: 'hidden',
        value: '',
        optionsList: releaseApprovalTypeOptions,
    },
    approvalTypeName: {
        label: 'Approval Type',
        type: 'hidden',
        value: '',
    },
    approvalStatus: {
        label: 'Approval Status',
        type: 'hidden',
        value: '',
        optionsList: releaseApprovalStatusOptions,
    },
    approvalStatusName: {
        label: 'Approval Status',
        type: 'hidden',
        value: '',
    },
    releaseStatus: {
        label: 'Status',
        type: 'hidden',
        value: '',
        optionsList: releaseStatusOptions,
    },
    releaseStatusName: {
        label: 'Status',
        type: 'hidden',
        value: '',
    },
    runPipeline: {
        label: 'Pipeline',
        type: 'hidden',
        value: '',
    },
    runPipelineName: {
        label: 'Pipeline',
        type: 'hidden',
        value: '',
    },
    runStage: {
        label: 'Stage',
        type: 'hidden',
        value: '',
    },
    runEnvironment: {
        label: 'Environment',
        type: 'hidden',
        value: '',
    },
    runStageId: {
        label: 'Stage',
        type: 'hidden',
        value: '',
    },
    runStageStateId: {
        label: 'State',
        type: 'hidden',
        value: '',
        optionsList: runStageStateIdOptions,
    },
    runStageStateIdName: {
        label: 'State',
        type: 'hidden',
        value: '',
        optionsList: runStageStateIdOptions,
    },
    runStageResultId: {
        label: 'Result',
        type: 'hidden',
        value: '',
        optionsList: runStageResultIdOptions,
    },
    runStateId: {
        label: 'State',
        type: 'hidden',
        value: '',
        optionsList: runStateIdOptions,
    },
    runStateIdName: {
        label: 'State',
        type: 'hidden',
        value: '',
        optionsList: runStateIdOptions,
    },
    runResultId: {
        label: 'Result',
        type: 'hidden',
        value: '',
        optionsList: runResultIdOptions,
    },

    // add 'timestamp' field only if you don't want to use cached RTK API query
    timestamp: {
        label: 'time',
        type: 'timestamp',
        value: '',
    },
};

// Create task modal
const taskTypeOptions = [
    {
        value: 'Task',
        label: 'Task',
    },
    {
        value: 'Epic',
        label: 'Epic',
    },
    {
        value: 'Issue',
        label: 'Issue',
    },
];

export const createTaskModal: Record<CreateTaskModalFields, ModalFormFieldConfig> = {
    organization: {
        label: 'Organization name',
        type: 'dropdown',
        value: '',
        validations: {
            isRequired: true,
        },
    },
    project: {
        label: 'Project name',
        value: '',
        type: 'dropdown',
        validations: {
            isRequired: true,
        },
    },
    type: {
        label: 'Work item type',
        value: '',
        type: 'dropdown',
        optionsList: taskTypeOptions,
        validations: {
            isRequired: true,
        },
    },
    title: {
        label: 'Title',
        value: '',
        type: 'text',
        validations: {
            isRequired: true,
        },
    },
    description: {
        label: 'Description',
        value: '',
        type: 'text',
    },
    areaPath: {
        label: 'Area Path',
        value: '',
        type: 'hidden',
    },

    // add 'timestamp' field only if you don't want to use cached RTK API query
    timestamp: {
        label: 'time',
        type: 'timestamp',
        value: '',
    },
};

// Link project modal
export const linkProjectModal: Record<LinkProjectModalFields, ModalFormFieldConfig> = {
    organization: {
        label: 'Organization name',
        type: 'text',
        value: '',
        validations: {
            isRequired: true,
        },
    },
    project: {
        label: 'Project name',
        value: '',
        type: 'text',
        validations: {
            isRequired: true,
        },
    },

    // add 'timestamp' field only if you don't want to use cached RTK API query
    timestamp: {
        label: 'time',
        type: 'timestamp',
        value: '',
    },
};

export const subscriptionFilterCreatedByOptions = [
    {
        value: subscriptionFilters.createdBy.me,
        label: 'Me',
    },
    {
        value: subscriptionFilters.createdBy.anyone,
        label: 'Anyone',
    },
];

export const subscriptionFilterServiceTypeOptions = [
    {
        value: subscriptionFilters.serviceType.boards,
        label: 'Boards',
    },
    {
        value: subscriptionFilters.serviceType.repos,
        label: 'Repos',
    },
    {
        value: subscriptionFilters.serviceType.pipelines,
        label: 'Pipelines',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const subscriptionFilterEventTypeBoardsOptions = () => {
    const options: LabelValuePair[] = [];
    Object.keys(subscriptionFilters.eventType.boards).forEach((eventType) => options.push({
        value: eventType,
        label: subscriptionFilters.eventType.boards[eventType as EventTypeBoards],
    }));

    options.push(filterLabelValuePairAll);
    return options;
};

export const subscriptionFilterEventTypeReposOptions = () => {
    const options: LabelValuePair[] = [];
    Object.keys(subscriptionFilters.eventType.repos).forEach((eventType) => options.push({
        value: eventType,
        label: subscriptionFilters.eventType.repos[eventType as EventTypeRepos],
    }));

    options.push(filterLabelValuePairAll);
    return options;
};

export const subscriptionFilterEventTypePipelinesOptions = () => {
    const options: LabelValuePair[] = [];
    Object.keys(subscriptionFilters.eventType.pipelines).forEach((eventType) => options.push({
        value: eventType,
        label: subscriptionFilters.eventType.pipelines[eventType as EventTypePipelines],
    }));

    options.push(filterLabelValuePairAll);
    return options;
};

// Repos subscription filters
export const mergeResultOptons: LabelValuePair[] = [
    {
        value: 'Succeeded',
        label: 'Merge successful',
    },
    {
        value: 'Unsuccessful',
        label: 'Merge Unsuccessful - Reason: Any',
    },
    {
        value: 'Conflicts',
        label: 'Merge Unsuccessful - Reason: Conflicts',
    },
    {
        value: 'Failure',
        label: 'Merge Unsuccessful - Reason: Failure',
    },
    {
        value: 'RejectedByPolicy',
        label: 'Merge Unsuccessful - Reason: Rejected By Policy',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const pullRequestChangeOptons: LabelValuePair[] = [
    {
        value: 'PushNotification',
        label: 'Source branch updated',
    },
    {
        value: 'ReviewersUpdateNotification',
        label: 'Reviewers changed',
    },
    {
        value: 'StatusUpdateNotification',
        label: 'Status changed',
    },
    {
        value: 'ReviewerVoteNotification',
        label: 'Votes score changed',
    },
    {
        ...filterLabelValuePairAll,
    },
];

export const subscriptionFiltersNameForBoards = {
    areaPath: 'areaPath',
};

export const subscriptionFiltersForBoards = [
    subscriptionFiltersNameForBoards.areaPath,
];

export const subscriptionFiltersNameForRepos = {
    repository: 'repository',
    branch: 'branch',
    pullrequestCreatedBy: 'pullrequestCreatedBy',
    pullrequestReviewersContains: 'pullrequestReviewersContains',
    pushedBy: 'pushedBy',
};

export const subscriptionFiltersForRepos = [
    subscriptionFiltersNameForRepos.repository,
    subscriptionFiltersNameForRepos.branch,
    subscriptionFiltersNameForRepos.pullrequestCreatedBy,
    subscriptionFiltersNameForRepos.pullrequestReviewersContains,
    subscriptionFiltersNameForRepos.pushedBy,
];

export const subscriptionFiltersNameForPipelines = {
    buildPipeline: 'definitionName',
    releasePipelineName: 'releaseDefinitionId',
    stageName: 'releaseEnvironmentId',
    runPipeline: 'pipelineId',
    runStage: 'stageName',
    runEnvironment: 'environmentName',
    runStageId: 'stageNameId',
};

export const subscriptionFiltersForPipelines = [
    subscriptionFiltersNameForPipelines.buildPipeline,
    subscriptionFiltersNameForPipelines.releasePipelineName,
    subscriptionFiltersNameForPipelines.stageName,
    subscriptionFiltersNameForPipelines.runPipeline,
    subscriptionFiltersNameForPipelines.runStage,
    subscriptionFiltersNameForPipelines.runEnvironment,
    subscriptionFiltersNameForPipelines.runStageId,
];
