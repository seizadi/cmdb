import { ActionReducerMap } from '@ngrx/store';

import * as fromApplications from '../applications/store/applications.reducers'
import * as fromAuth from '../auth/store/auth.reducers';

export interface AppState {
  applications: fromApplications.State,
  auth: fromAuth.State
}

export const reducers: ActionReducerMap<AppState> = {
  applications: fromApplications.applicationsReducer,
  auth: fromAuth.authReducer
};
