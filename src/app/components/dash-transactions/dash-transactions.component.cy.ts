import { MatDialog, MatDialogModule } from '@angular/material/dialog'
import { DashTransactionsComponent, TransactionsDialogComponent } from './dash-transactions.component'
import { MatButtonModule } from '@angular/material/button'
import { MatTableModule } from '@angular/material/table'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { MatIconModule } from '@angular/material/icon'
import { MatInputModule } from '@angular/material/input'
import { Observable, from, of } from 'rxjs'
import { TransactionService } from 'src/app/transaction.service'
import { CreateTransactionRequest, CreateTransactionResponse, GetTransactionsResponse, Transaction, UpdateTransactionRequest } from 'src/types/transaction-system'
import { MatFormFieldModule } from '@angular/material/form-field'
import { FormsModule, ReactiveFormsModule } from '@angular/forms'
import { MatCheckboxModule } from '@angular/material/checkbox'
import { MatNativeDateModule } from '@angular/material/core'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatSelectModule } from '@angular/material/select'
import { BudgetService } from 'src/app/budget.service'
import { GetBudgetCategoriesResponse } from 'src/types/budget-system'
import { createTransaction, deleteTransaction, getTransactions, updateTransaction } from 'src/sample-transaction-data'
import { GenericResponse } from 'src/types/api-system'

describe('DashTransactionsComponent', () => {

  const mockTransactionService: Partial<TransactionService> = {
    getTransactions(): Observable<GetTransactionsResponse> {
      return from(getTransactions())
    },
    createTransaction(transactionRequest: CreateTransactionRequest): Observable<CreateTransactionResponse> {
      return from(createTransaction(transactionRequest))
    },
    updateTransaction(transactionRequest: UpdateTransactionRequest): Observable<GenericResponse> {
      return from(updateTransaction(transactionRequest))
    },
    deleteTransaction(transaction: Transaction): Observable<GenericResponse> {
      return from(deleteTransaction(transaction.transactionId))
    }
  };

  const mockBudgetService: Partial<BudgetService> = {
    getBudgetCategories(): Observable<GetBudgetCategoriesResponse> {
      return of({
        categories: [
          "General",
          "Groceries",
          "Food",
          "Rent",
          "Supplies",
        ]
      });
    }
  };


  beforeEach(() => {

    cy.mount(DashTransactionsComponent, {
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
        MatDialog,
        { provide: TransactionService, useValue: mockTransactionService },
        { provide: BudgetService, useValue: mockBudgetService }
      ],
      declarations: [
        TransactionsDialogComponent
      ]
    }).then((wrapper) => {
      cy.spy(wrapper.component, 'openAddDialog').as('openAdd');
      cy.spy(wrapper.component, 'openEditDialog').as('openEdit');
      cy.spy(wrapper.component, 'goDeleteTransaction').as('goDelete');
      return cy.wrap(wrapper).as('angular');

    });
  });

  it('should mount', () => { });

  it('should have a table', () => {
    cy.get('table[mat-table]').should('exist');
  });

  it('should have a button that opens a dialog', () => {

    cy.get('div.transactions-dialog').should('not.exist');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('div.transactions-dialog')
      .should('exist')
      .should('be.visible');

    cy.get('body').click(0, 0);
    cy.get('div.transactions-dialog').should('not.exist');
  });

  it('should have 6 entries', () => {
    cy.get('tr.transaction__row').should('have.length', 6);
  })

  it('should have a detail row that pops out', () => {
    cy.get('.transaction__detail-contents').should('not.be.visible');
    cy.get('[data-cy="expand-btn"]').eq(0).click();
    cy.get('.transaction__detail-contents').should('be.visible').should('contain', '[none]');
    cy.get('[data-cy="expand-btn"]').eq(1).click();
    cy.get('.transaction__detail-contents').should('be.visible').should('contain', "Also paid for my friend's drink.");
  })

  it('should have a dialog with selectable options', () => {
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="category"]').should('have.text', '[None]');
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option').should('have.length', 6);
  })

  it('should be able to add a new transaction', () => {
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="name"]').type('Test Transaction');
    cy.get('[formControlName="amount"]').type('123');
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option').eq(1).click();
    cy.get('[formControlName="date"]').clear().type('2020-01-01');
    cy.get('[formControlName="description"]').type('This is a test transaction.');
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('tr.transaction__row').should('have.length', 7);
    cy.get('tr.transaction__row').eq(6).should('contain', 'Test Transaction')
      .should('contain', '123.00')
      .should('contain', 'General');
    cy.get('tr.transaction__row').eq(6).find('[data-cy="expand-btn"]').click();
    cy.get('.transaction__detail-contents').should('be.visible').should('contain', 'This is a test transaction.');
  })

  it('should be able to edit a transaction', () => {
    cy.get('[data-cy="edit-btn"]').eq(0).click();
    cy.get('[formControlName="name"]').clear().type('Edited Transaction');
    cy.get('[formControlName="amount"]').clear().type('321');
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option').eq(1).click();
    cy.get('[formControlName="date"]').clear().type('2020-02-02');
    cy.get('[formControlName="description"]').clear().type('This is an edited transaction.');
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('tr.transaction__row').eq(0).should('contain', 'Edited Transaction')
      .should('contain', '321.00')
      .should('contain', 'General');
    cy.get('tr.transaction__row').eq(0).find('[data-cy="expand-btn"]').click();
    cy.get('.transaction__detail-contents').should('be.visible').should('contain', 'This is an edited transaction.');
  })

  it('should be able to delete a transaction', () => {
    cy.get('tr.transaction__row').should('have.length', 7);
    cy.get('[data-cy="delete-btn"]').eq(0).click();
    cy.get('tr.transaction__row').should('have.length', 6);
  })
})