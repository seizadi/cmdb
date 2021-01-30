import {Injectable} from '@angular/core';
import {Actions, Effect, ofType} from '@ngrx/effects';
import {switchMap, map} from 'rxjs/operators';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Router} from '@angular/router';

import * as ApplicationsActions from './applications.actions';
import {Application} from "../../shared/application.model";

// TODO - handle API Errrors
// catchError(err => of(new someActions.FetchFailed()))

@Injectable()
export class ApplicationsEffects {
    @Effect()
    applicationsFetch = this.actions$.pipe(
        ofType(ApplicationsActions.FETCH_APPLICATIONS),
        switchMap((action: ApplicationsActions.FetchApplications) => {
            return this.httpClient.get('v1/applications', {
                headers: new HttpHeaders({Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBY2NvdW50SUQiOjF9.GsXyFDDARjXe1t9DPo2LIBKHEal3O7t3vLI3edA7dGU'}),
                observe: 'body',
                responseType: 'json'
            });
        }),
        map((responseApplications: any) => {
                console.log(responseApplications);
                const applications: Application[] = [];
                for (const application of responseApplications.results) {
                    applications.push(application);
                }
                console.log(applications);
                return {
                    type: ApplicationsActions.SET_APPLICATIONS,
                    payload: applications
                };
            }
        ));


    constructor(private actions$: Actions,
                private router: Router,
                private httpClient: HttpClient) { }
}
