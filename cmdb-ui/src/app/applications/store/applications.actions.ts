import { Action } from '@ngrx/store';

import { Application } from '../../shared/application.model';

export const FETCH_APPLICATIONS = 'FETCH_APPLICATIONS';
export const SET_APPLICATIONS = 'SET_APPLICATIONS';


export class FetchApplications implements Action {
    readonly type = FETCH_APPLICATIONS;
}

export class SetApplications implements Action {
    readonly type = SET_APPLICATIONS;

    constructor(public payload: Application[]) {}
}

export type ApplicationsActions = FetchApplications | SetApplications;
