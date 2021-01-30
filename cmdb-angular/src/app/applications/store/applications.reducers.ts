import * as ApplicationsActions from './applications.actions';

import { Application } from '../../shared/application.model';

export interface State {
    applications: Application[];
}

const initialState: State = {
    applications: [],
};

export function applicationsReducer(state = initialState, action: ApplicationsActions.ApplicationsActions) {
    switch (action.type) {
        case (ApplicationsActions.SET_APPLICATIONS):
            return {
                ...state,
                applications: [...action.payload]
            };
        default:
            return state;
    }
}
