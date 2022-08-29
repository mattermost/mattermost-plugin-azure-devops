import {combineReducers} from 'redux';

import services from 'services';

import globalModalSlice from './globalModal';
import apiRequestCompletionSlice from './apiRequest';
import linkProjectModalSlice from './linkModal';
import subscriptionModalSlice from './subscribeModal';
import openTaskModalReducer from './taskModal';
import projectDetailsSlice from './projectDetails';
import userConnectedSlice from './userConnected';

const reducers = combineReducers({
    apiRequestCompletionSlice,
    globalModalSlice,
    linkProjectModalSlice,
    openTaskModalReducer,
    subscriptionModalSlice,
    projectDetailsSlice,
    userConnectedSlice,
    [services.reducerPath]: services.reducer,
});

export default reducers;
