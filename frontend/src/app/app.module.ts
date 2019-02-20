import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatDatepickerModule,
  MatFormFieldModule,
  MatInputModule,
  MAT_DATE_LOCALE,
  MatButtonModule
} from '@angular/material';
import { MatMomentDateModule } from '@angular/material-moment-adapter';
import { MAT_MOMENT_DATE_FORMATS, MomentDateAdapter} from '@angular/material-moment-adapter';
import { DateAdapter, MAT_DATE_FORMATS, MatNativeDateModule } from '@angular/material/core';

import { AppComponent } from './app.component';
import 'hammerjs';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

const materialModules = [
  MatButtonModule,
  MatDatepickerModule,
  MatNativeDateModule,
  MatFormFieldModule,
  MatInputModule,
  MatMomentDateModule
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

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserAnimationsModule,
    BrowserModule,
    FormsModule,
    HttpClientModule,
    ReactiveFormsModule,
    ...materialModules
  ],
  exports: [
    MatDatepickerModule
  ],
  providers: [
    MatDatepickerModule,
    {
      provide: DateAdapter,
      useClass: MomentDateAdapter,
      deps: [MAT_DATE_LOCALE]
    },
    { provide: MAT_DATE_LOCALE, useValue: 'ja-JP' },
    {
        provide: MAT_DATE_FORMATS,
        useValue: MY_DATE_FORMATS
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
