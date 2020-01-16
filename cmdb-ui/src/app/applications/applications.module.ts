import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import {ApplicationsComponent} from "./applications.component";
import { ApplicationListComponent } from "./application-list/application-list.component";

@NgModule({
    declarations: [
        ApplicationsComponent,
        ApplicationListComponent
    ],
    imports: [
        CommonModule,
        FormsModule
    ]
})
export class ApplicationsModule {}
