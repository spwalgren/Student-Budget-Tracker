import { MatButtonModule } from '@angular/material/button'
import { BudgetsDialogComponent, DashBudgetsComponent } from './dash-budgets.component'
import { MatTableModule } from '@angular/material/table'
import { MatIconModule } from '@angular/material/icon'
import { MatInputModule } from '@angular/material/input'
import { MatDialog, MatDialogModule } from '@angular/material/dialog'
import { BudgetService } from 'src/app/budget.service'
import { Observable, from, of } from 'rxjs'
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

describe('DashBudgetsComponent', () => {

  const mockBudgetService: Partial<BudgetService> = {
    createBudget(requestData: CreateBudgetRequest): Observable<CreateBudgetResponse> {
      return from(createBudget(requestData));
    },
    updateBudget(requestData: UpdateBudgetRequest): Observable<GenericResponse> {
      return from(updateBudget(requestData));
    },
    getBudgets(): Observable<GetBudgetsResponse> {
      return from(getBudgets());
    },
    deleteBudget(budgetId: number): Observable<GenericResponse> {
      return from(deleteBudget(budgetId));
    },
    getCyclePeriod(): Observable<GetCyclePeriodResponse> {
      return of({
        data: []
      })
    }
  };

  beforeEach(() => {
    cy.mount(DashBudgetsComponent, {
      imports: [
        MatTableModule,
        MatButtonModule,
        MatIconModule,
        MatInputModule,
        MatDialogModule,
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
        MatDialog, { provide: BudgetService, useValue: mockBudgetService }
      ],
      declarations: [
        BudgetsDialogComponent
      ]
    })
  })

  it('should mount', () => {

  })

  it('should have 4 different tables', () => {
    cy.get('table[mat-table]').should('have.length', 4)
    cy.get('table[mat-table]').each(($elem) => {
      cy.wrap($elem).should('be.visible');
    })
  })

  it('should be able to delete tables', () => {
    cy.get('table[mat-table].Books').should('exist');
    cy.get('table[mat-table].Books button[data-cy="delete-btn"]').click();
    cy.get('table[mat-table].Books').should('not.exist');
  })

  it('should have a dialog box', () => {
    cy.get('button[data-cy="add-btn"]').click();
    cy.get('.budgets-dialog').should('exist').should('be.visible');
    cy.get('button[data-cy="cancel-btn"]').click();
    cy.get('.budgets-dialog').should('not.exist')
  })

  it('should be able to add tables', () => {
    cy.get('table.General').should('not.exist');
    cy.get('button[data-cy="add-btn"]').click();
    cy.get('.budgets-dialog').should('exist').should('be.visible');
    cy.get('button[data-cy="submit-btn"]').click();
    cy.get('.budgets-dialog').should('not.exist');
    cy.get('table.General').should('exist').should('be.visible');
  })

  it('should be able to add to a table', () => {
    cy.get('table.Food tr').should('have.length', 2);
    cy.get('button[data-cy="add-btn"]').click();
    cy.get('[formControlName="category"]').clear().type('Food');
    cy.get('button[data-cy="submit-btn"]').click();
    cy.get('table.Food tr').should('have.length', 3);
  })

  it('should be able to edit table entry', () => {
    cy.get('table.Food tr').should('have.length', 3);
    cy.get('table.Food tr').eq(1).should('contain.text', 'Every 2 weeks');
    cy.get('table.Food tr button[data-cy="edit-btn"]').eq(0).click();
    cy.get('.budgets-dialog').should('exist').should('be.visible').should('contain.text', 'Edit Budget');
    cy.get('[formControlName="amount"]').clear().type('10000');
    cy.get('[formControlName="duration"]').clear().type('1');
    cy.get('[formControlName="frequency"]').click();
    cy.get('mat-option[ng-reflect-value="yearly"]').click();
    cy.get('button[data-cy="submit-btn"]').click();
    cy.get('table.Food tr').eq(1).should('contain.text', '$10,000.00');
    cy.get('table.Food tr').eq(1).should('contain.text', 'Every year');
  })

  it('should be able to delete table entry', () => {
    cy.get('table.Food tr').should('have.length', 3);
    cy.get('table.Food tr button[data-cy="delete-btn"]').eq(0).click();
    cy.get('table.Food tr').should('have.length', 2);
  })
})