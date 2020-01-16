import { Component, OnInit } from '@angular/core';
import {Observable} from "rxjs";

import {Application} from "../../shared/application.model";
import * as fromApp from '../../store/app.reducers';
import {Store} from "@ngrx/store";
import * as ApplicationsActions from "../store/applications.actions";

@Component({
  selector: 'app-application-list',
  templateUrl: './application-list.component.html',
  styleUrls: ['./application-list.component.css']
})
export class ApplicationListComponent implements OnInit {
  applicationsState: Observable<{applications: Application[]}>;

  constructor(private store: Store<fromApp.AppState>) { }

  ngOnInit() {
    this.store.dispatch(new ApplicationsActions.FetchApplications());
    this.applicationsState = this.store.select('applications');
  }

}
