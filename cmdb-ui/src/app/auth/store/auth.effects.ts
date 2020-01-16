import {Injectable} from '@angular/core';
import {Actions, Effect, ofType} from '@ngrx/effects';
import {Router} from '@angular/router';
import {map, tap, switchMap, mergeMap} from 'rxjs/operators';
import { from } from 'rxjs';
import {HttpClient, HttpRequest} from '@angular/common/http';

import * as AuthActions from './auth.actions';

@Injectable()
export class AuthEffects {
  @Effect({dispatch: false})
  authSignup = this.actions$.pipe(
    ofType(AuthActions.TRY_SIGNUP),
    map((action: AuthActions.TrySignup) => {
        return action.payload;
      })
      , map((authData: { username: string, password: string }) => {
        const user = {user: authData.username, password: authData.password, email: '', remember: false};
        const req = new HttpRequest('POST', 'http://localhost:3000/login', user, {reportProgress: true});
        return this.httpClient.request(req);
      }));

  @Effect({dispatch: false})
  authSignin = this.actions$.pipe(
    ofType(AuthActions.TRY_SIGNIN),
    map((action: AuthActions.TrySignup) => {
        return action.payload;
      })
      , switchMap((authData: { username: string, password: string }) => {
        const user = {user: authData.username, password: authData.password, email: '', remember: false};
        const req = new HttpRequest('POST', 'http://localhost:3000/login', user, {reportProgress: true});
        const resp = this.httpClient.request(req);
        console.log(resp);
        return resp;
      }));

  @Effect({dispatch: false})
  authLogout = this.actions$.pipe(
    ofType(AuthActions.LOGOUT),
    tap(() => {
      this.router.navigate(['/']);
    }));

  constructor(private actions$: Actions, private router: Router, private httpClient: HttpClient) {
  }
}
