import { ProgressService } from 'src/app/progress.service';
import { DashProgressComponent } from './dash-progress.component'
import { Period } from 'src/types/budget-system';
import { Observable, of } from 'rxjs';
import { GetProgressResponse } from 'src/types/progress.system';
import { MatDialog } from '@angular/material/dialog';


import { MatButtonModule } from '@angular/material/button'
import { MatTableModule } from '@angular/material/table'
import { MatIconModule } from '@angular/material/icon'
import { MatInputModule } from '@angular/material/input'
//import { MatDialog, MatDialogModule } from '@angular/material/dialog'
import { BudgetService } from 'src/app/budget.service'
//import { Observable, from, of } from 'rxjs'
import { CreateBudgetRequest, CreateBudgetResponse, GetBudgetsResponse, GetCyclePeriodResponse, UpdateBudgetRequest } from 'src/types/budget-system'
import { createBudget, deleteBudget, getBudgets, updateBudget } from 'src/sample-budget-data'
import { GenericResponse } from 'src/types/api-system'
import { MatFormFieldModule } from '@angular/material/form-field'
import { MatCheckboxModule } from '@angular/material/checkbox'
import { MatSidenavModule } from '@angular/material/sidenav'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatSelectModule } from '@angular/material/select'
import { MatNativeDateModule } from '@angular/material/core'
import { NgModule } from '@angular/core'
import { FormsModule, ReactiveFormsModule } from '@angular/forms'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'

describe('DashProgressComponent', () => {
  
  const mockProgressService: Partial<ProgressService> = {
    GetProgress(): Observable<GetProgressResponse> {
      return of({
        data: [
          {
            userId: 20,
            totalSpent: 10,
            transactionIdList: [5],
            budgetId: 11,
            category: "General",
            budgetGoal: 100,
            frequency: Period.weekly,
          }
        ]
        
      })
    }
  }
  // GetPreviousProgress(): Observable<GetProgressResponse> {
  //   return of({
  //     data: [
  //       {
  //         userId: 20,
  //         totalSpent: 10,
  //         transactionIdList: [5],
  //         budgetId: 11,
  //         category: "General",
  //         budgetGoal: 100,
  //         frequency: Period.weekly,
  //       }
  //     ]
      
  //   })
  //}

  beforeEach(() => {
    cy.mount(DashProgressComponent, {
      imports: [
        MatTableModule,
        MatButtonModule,
        MatIconModule,
        MatInputModule,
        MatFormFieldModule,
        MatSelectModule,
        MatDatepickerModule,
        MatNativeDateModule,
        MatCheckboxModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
      ],
      providers: [
        MatDialog, { provide: ProgressService, useValue: mockProgressService }
      ],
      declarations: [
        DashProgressComponent
      ]
    })
  })
  it('should mount', () => {
    //cy.mount(DashProgressComponent)
  })
})

