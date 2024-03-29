// Angular modules
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  FormsModule,
  ReactiveFormsModule
} from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import { AppRoutingModule } from './app-routing.module';

// Material modules
import { MatDialogModule } from '@angular/material/dialog';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { MatSelectModule } from '@angular/material/select';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatTabsModule } from '@angular/material/tabs';

// Route components
import { LandingComponent } from './routes/landing/landing.component';
import { LoginComponent } from './routes/login/login.component';
import { SignUpComponent } from './routes/sign-up/sign-up.component';

// Other components
import { AppComponent } from './app.component';
import { AlertComponent } from './components/alert/alert.component';
import { DashboardComponent } from './routes/dashboard/dashboard.component';
import { DashHomeComponent } from './components/dash-home/dash-home.component';
import { DashTransactionsComponent, TransactionsDialogComponent } from './components/dash-transactions/dash-transactions.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { BudgetsDialogComponent, DashBudgetsComponent } from './components/dash-budgets/dash-budgets.component';
import { DashSettingsComponent } from './components/dash-settings/dash-settings.component';
import { DashCalendarComponent } from './components/dash-calendar/dash-calendar.component';
import { CalendarModule, DateAdapter } from 'angular-calendar';
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns';
import { DashProgressComponent } from './components/dash-progress/dash-progress.component';
import { EventCardComponent } from './components/event-card/event-card.component';
import { MatFormFieldModule } from '@angular/material/form-field';

// Services

@NgModule({
  declarations: [
    AppComponent,
    LandingComponent,
    LoginComponent,
    SignUpComponent,
    AlertComponent,
    DashboardComponent,
    DashHomeComponent,
    DashTransactionsComponent,
    TransactionsDialogComponent,
    PageNotFoundComponent,
    DashBudgetsComponent,
    BudgetsDialogComponent,
    DashSettingsComponent,
    DashCalendarComponent,
    DashProgressComponent,
    EventCardComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    AppRoutingModule,
    MatToolbarModule,
    MatIconModule,
    MatCardModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatSidenavModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatDialogModule,
    MatDatepickerModule,
    MatNativeDateModule,
    MatSelectModule,
    MatCheckboxModule,
    MatTabsModule,
    MatFormFieldModule,
    CalendarModule.forRoot({ provide: DateAdapter, useFactory: adapterFactory })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
