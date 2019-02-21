import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatDatepickerModule,
  MatFormFieldModule,
  MatInputModule,
  MAT_DATE_LOCALE,
  MatButtonModule,
  MatTableModule,
  MatMenuModule,
  MatIconModule,
  MatDialogModule
} from '@angular/material';
import { DateAdapter, MAT_DATE_FORMATS, MatNativeDateModule } from '@angular/material/core';
import { MatSortModule } from '@angular/material/sort';
import { MatMomentDateModule, MomentDateAdapter } from '@angular/material-moment-adapter';

import { Nl2BrPipeModule } from 'nl2br-pipe';
import 'hammerjs';

import { AppComponent } from './app.component';
import {
  OverviewDialogComponent,
  AppComponentsPeriodViewComponent
} from './components';

const materialModules = [
  MatButtonModule,
  MatDialogModule,
  MatDatepickerModule,
  MatNativeDateModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule,
  MatMenuModule,
  MatMomentDateModule,
  MatTableModule,
  MatSortModule
];

const MY_DATE_FORMATS = {
  parse: {
    dateInput: 'LL',
  },
  display: {
    dateInput: 'LL',
    monthYearLabel: 'MMM YYYY',
    dateA11yLabel: 'LL',
    monthYearA11yLabel: 'MMMM YYYY',
  },
};

const COMPONETS = [
  OverviewDialogComponent,
  AppComponentsPeriodViewComponent
];

@NgModule({
  declarations: [
    AppComponent,
    ...COMPONETS
  ],
  imports: [
    BrowserAnimationsModule,
    BrowserModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    Nl2BrPipeModule,
    ...materialModules
  ],
  exports: [
    MatDatepickerModule,
    ...COMPONETS
  ],
  providers: [
    MatDatepickerModule,
    {
      provide: DateAdapter,
      useClass: MomentDateAdapter,
      deps: [MAT_DATE_LOCALE]
    },
    {
      provide: MAT_DATE_LOCALE,
      useValue: 'ja-JP' },
    {
      provide: MAT_DATE_FORMATS,
      useValue: MY_DATE_FORMATS
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
