import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatButtonModule,
  MatCheckboxModule,
  MatDatepickerModule,
  MatDialogModule,
  MatExpansionModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule,
  MatListModule,
  MatMenuModule,
  MatRadioModule,
  MatTabsModule,
  MatTableModule,
  MAT_DATE_LOCALE
} from '@angular/material';
import { DateAdapter, MatNativeDateModule, MAT_DATE_FORMATS } from '@angular/material/core';
import { MatSortModule } from '@angular/material/sort';
import { MatMomentDateModule, MomentDateAdapter } from '@angular/material-moment-adapter';

import { Nl2BrPipeModule } from 'nl2br-pipe';
import 'hammerjs';

import { AppComponent } from './app.component';
import {
  OverviewDialogComponent,
  AppComponentsDetailComponent,
  AppComponentsPeriodViewComponent
} from './components';

const materialModules = [
  MatButtonModule,
  MatCheckboxModule,
  MatDialogModule,
  MatDatepickerModule,
  MatExpansionModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule,
  MatListModule,
  MatMenuModule,
  MatMomentDateModule,
  MatNativeDateModule,
  MatRadioModule,
  MatTabsModule,
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
  AppComponentsDetailComponent,
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
