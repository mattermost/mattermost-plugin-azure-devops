import React, {useEffect, useState} from 'react';
import {useDispatch, useSelector} from 'react-redux';

import {GlobalState} from 'mattermost-redux/types/store';

import Modal from 'components/modal';
import CircularLoader from 'components/loader/circular';
import Form from 'components/form';
import EmptyState from 'components/emptyState';
import ResultPanel from 'components/resultPanel';

import plugin_constants from 'plugin_constants';

import useApiRequestCompletionState from 'hooks/useApiRequestCompletionState';
import usePluginApi from 'hooks/usePluginApi';
import useForm from 'hooks/useForm';

import {toggleShowSubscribeModal} from 'reducers/subscribeModal';
import {toggleShowLinkModal} from 'reducers/linkModal';
import {getSubscribeModalState} from 'selectors';

import Utils from 'utils';

import './styles.scss';

const SubscribeModal = () => {
    const {subscriptionModal} = plugin_constants.form;

    // Hooks
    const {
        formFields,
        errorState,
        onChangeFormField,
        setSpecificFieldValue,
        resetFormFields,
        isErrorInFormValidation,
    } = useForm(subscriptionModal);
    const {
        getApiState,
        makeApiRequest,
        makeApiRequestWithCompletionStatus,
        state,
    } = usePluginApi();
    const {visibility} = getSubscribeModalState(state);
    const {currentTeamId} = useSelector((reduxState: GlobalState) => reduxState.entities.teams);
    const dispatch = useDispatch();

    // State variables
    const [channelOptions, setChannelOptions] = useState<LabelValuePair[]>([]);
    const [organizationOptions, setOrganizationOptions] = useState<LabelValuePair[]>([]);
    const [projectOptions, setProjectOptions] = useState<LabelValuePair[]>([]);
    const [showResultPanel, setShowResultPanel] = useState(false);

    // Function to hide the modal and reset all the states.
    const resetModalState = (isActionDone?: boolean) => {
        dispatch(toggleShowSubscribeModal({isVisible: false, commandArgs: [], isActionDone}));
        resetFormFields();
        setOrganizationOptions([]);
        setProjectOptions([]);
        setChannelOptions([]);
        setShowResultPanel(false);
    };

    // Get organization and project state
    const getOrganizationAndProjectState = () => {
        const {isLoading, isSuccess, isError, data} = getApiState(
            plugin_constants.pluginApiServiceConfigs.getAllLinkedProjectsList.apiServiceName,
        );

        return {
            isLoading,
            isError,
            isSuccess,
            organizationList: isSuccess ? Utils.getOrganizationList(data as ProjectDetails[]) : [],
            projectList: isSuccess ? Utils.getProjectList(data as ProjectDetails[]) : [],
        };
    };

    // Get channel state
    const getChannelState = () => {
        const {isLoading, isSuccess, isError, data} = getApiState(
            plugin_constants.pluginApiServiceConfigs.getChannels.apiServiceName,
            {teamId: currentTeamId},
        );
        return {isLoading, isSuccess, isError, data: data as ChannelList[]};
    };

    // Get option list for each types of dropdown fields
    const getDropDownOptions = (fieldName: SubscriptionModalFields) => {
        switch (fieldName) {
        case 'organization':
            return organizationOptions;
        case 'project':
            return projectOptions;
        case 'eventType':
            return subscriptionModal.eventType.optionsList;
        case 'channelID':
            return channelOptions;
        default:
            return [];
        }
    };

    // Opens link project modal
    const handleOpenLinkProjectModal = () => {
        dispatch(toggleShowLinkModal({isVisible: true, commandArgs: []}));
        resetModalState();
    };

    // Opens subscription modal
    const handleSubscriptionModal = () => {
        dispatch(toggleShowSubscribeModal({isVisible: true, commandArgs: []}));
        resetModalState();
    };

    // Return different types of error messages occurred on API call
    const showApiErrorMessages = (isCreateSubscriptionError: boolean, error: ApiErrorResponse) => {
        if (getChannelState().isError) {
            return plugin_constants.messages.error.errorFetchingChannelsList;
        }
        if (getOrganizationAndProjectState().isError) {
            return plugin_constants.messages.error.errorFetchingOrganizationAndProjectsList;
        }
        return Utils.getErrorMessage(isCreateSubscriptionError, 'SubscribeModal', error);
    };

    // Handles creating subscription on confirmation
    const onConfirm = () => {
        if (!isErrorInFormValidation()) {
            // Make POST api request to create subscription
            makeApiRequestWithCompletionStatus(
                plugin_constants.pluginApiServiceConfigs.createSubscription.apiServiceName,
                formFields as APIRequestPayload,
            );
        }
    };

    // Observe for the change in redux state after the API call to create a subscription and do the required actions
    useApiRequestCompletionState({
        serviceName: plugin_constants.pluginApiServiceConfigs.createSubscription.apiServiceName,
        handleSuccess: () => {
            setShowResultPanel(true);
            dispatch(toggleShowSubscribeModal({isVisible: true, commandArgs: [], isActionDone: true}));
        },
        payload: formFields as APIRequestPayload,
    });

    // Make API request to fetch channel list
    useEffect(() => {
        makeApiRequest(
            plugin_constants.pluginApiServiceConfigs.getChannels.apiServiceName,
            {teamId: currentTeamId},
        );
    }, [visibility]);

    // Pre-select the dropdown value in case of single option
    useEffect(() => {
        const autoSelectedValues: Pick<Record<FormFieldNames, string>, 'organization' | 'project' | 'channelID'> = {
            organization: '',
            project: '',
            channelID: '',
        };

        if (organizationOptions.length === 1) {
            autoSelectedValues.organization = organizationOptions[0].value;
        }
        if (projectOptions.length === 1) {
            autoSelectedValues.project = projectOptions[0].value;
        }
        if (channelOptions.length === 1) {
            autoSelectedValues.channelID = channelOptions[0].value;
        }

        if (autoSelectedValues.organization || autoSelectedValues.project || autoSelectedValues.channelID) {
            setSpecificFieldValue({
                ...formFields,
                ...autoSelectedValues,
            });
        }
    }, [projectOptions, organizationOptions, channelOptions]);

    // Set organization, project and channel list values
    useEffect(() => {
        if (getChannelState().isSuccess) {
            setChannelOptions(getChannelState().data?.map((channel) => ({
                label: <span><i className='fa fa-globe dropdown-option-icon'/>{channel.display_name}</span>,
                value: channel.id,
            })));
        }

        if (getOrganizationAndProjectState().isSuccess && !showResultPanel) {
            setOrganizationOptions(getOrganizationAndProjectState().organizationList);
            setProjectOptions(getOrganizationAndProjectState().projectList);
        }
    }, [
        getChannelState().isLoading,
        getOrganizationAndProjectState().isLoading,
        showResultPanel,
    ]);

    const {isLoading, isError, error} = getApiState(plugin_constants.pluginApiServiceConfigs.createSubscription.apiServiceName, formFields as APIRequestPayload);
    const isAnyProjectLinked = Boolean(getOrganizationAndProjectState().organizationList.length && getOrganizationAndProjectState().projectList.length);

    return (
        <Modal
            show={visibility}
            title='Add New Subscription'
            onHide={resetModalState}
            onConfirm={isAnyProjectLinked ? onConfirm : null}
            confirmBtnText='Add new subscription'
            confirmDisabled={isLoading}
            cancelDisabled={isLoading}
            loading={isLoading}
            showFooter={!showResultPanel}
            error={showApiErrorMessages(isError, error as ApiErrorResponse)}
        >
            <>
                {
                    (getChannelState().isLoading || getOrganizationAndProjectState().isLoading) && <CircularLoader/>
                }
                {
                    !showResultPanel && (
                        isAnyProjectLinked ? (
                            Object.keys(subscriptionModal).map((field) => (
                                <Form
                                    key={subscriptionModal[field as SubscriptionModalFields].label}
                                    fieldConfig={subscriptionModal[field as SubscriptionModalFields]}
                                    value={formFields[field as SubscriptionModalFields] ?? ''}
                                    optionsList={getDropDownOptions(field as SubscriptionModalFields)}
                                    onChange={(newValue) => onChangeFormField(field as SubscriptionModalFields, newValue)}
                                    error={errorState[field as SubscriptionModalFields]}
                                    isDisabled={isLoading}
                                />
                            ))
                        ) : (
                            <EmptyState
                                title='No Project Linked'
                                subTitle={{text: 'You can link a project by clicking the below button.'}}
                                buttonText='Link new project'
                                buttonAction={handleOpenLinkProjectModal}
                            />
                        )
                    )
                }
                {
                    showResultPanel && (
                        <ResultPanel
                            header='Subscription created successfully.'
                            primaryBtnText='Add new subscription'
                            secondaryBtnText='Close'
                            onPrimaryBtnClick={handleSubscriptionModal}
                            onSecondaryBtnClick={() => resetModalState(true)}
                        />
                    )
                }
            </>
        </Modal>
    );
};

export default SubscribeModal;
