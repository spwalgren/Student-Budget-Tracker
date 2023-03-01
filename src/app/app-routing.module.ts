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
import { DashGoalsComponent } from './components/dash-goals/dash-goals.component';


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
        path: 'goals',
        component: DashGoalsComponent
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
