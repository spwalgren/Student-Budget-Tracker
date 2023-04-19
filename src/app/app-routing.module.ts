import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Page components
import { LandingComponent } from './routes/landing/landing.component';
import { LoginComponent } from './routes/login/login.component';
import { SignUpComponent } from './routes/sign-up/sign-up.component';
import { DashboardComponent } from './routes/dashboard/dashboard.component';
import { DashHomeComponent } from './components/dash-home/dash-home.component';
import { DashTransactionsComponent } from './components/dash-transactions/dash-transactions.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { DashBudgetsComponent } from './components/dash-budgets/dash-budgets.component';
import { DashCalendarComponent } from './components/dash-calendar/dash-calendar.component';
import { DashSettingsComponent } from './components/dash-settings/dash-settings.component';
import { DashProgressComponent } from './components/dash-progress/dash-progress.component'



const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: '/landing'
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'sign-up',
    component: SignUpComponent
  },
  {
    path: 'landing',
    component: LandingComponent
  },
  {
    path: 'dashboard',
    component: DashboardComponent,
    children: [
      {
        path: '',
        component: DashHomeComponent
      },
      {
        path: 'transactions',
        component: DashTransactionsComponent
      },
      {
        path: 'budgets',
        component: DashBudgetsComponent
      },
      {
        path: 'progress',
        component: DashProgressComponent
      },
      {
        path: 'calendar',
        component: DashCalendarComponent
      },
      {
        path: 'settings',
        component: DashSettingsComponent
      },
      {
        path: '**',
        component: PageNotFoundComponent
      }
    ]
  },
  {
    path: '**',
    component: PageNotFoundComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
