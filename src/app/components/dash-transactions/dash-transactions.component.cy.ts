import { MatDialog, MatDialogModule } from '@angular/material/dialog'
import { DashTransactionsComponent } from './dash-transactions.component'
import { MatButtonModule } from '@angular/material/button'
import { MatTableModule } from '@angular/material/table'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { MatIconModule } from '@angular/material/icon'
import { MatInputModule } from '@angular/material/input'
import { Observable, of } from 'rxjs'
import { TransactionService } from 'src/app/transaction.service'
import { GetTransactionsResponse } from 'src/types/transaction-system'

describe('DashTransactionsComponent', () => {

  const mockTransactionService: Partial<TransactionService> = {
    getTransactions(): Observable<GetTransactionsResponse> {
      return of({
        data: [{
          userId: 20,
          transactionId: 0,
          name: "Publix",
          amount: 30,
          date: new Date("2023-02-18").toISOString(),
          category: "Groceries"
        },
        {
          userId: 20,
          transactionId: 1,
          name: "Starbucks",
          amount: 8,
          date: new Date("2022-1-19").toISOString(),
          category: "Food",
          description: "Also paid for my friend's drink."
        }]
      });
    }
  };

  beforeEach(() => {

    cy.mount(DashTransactionsComponent, {
      imports: [
        MatDialogModule,
        MatButtonModule,
        MatTableModule,
        BrowserAnimationsModule,
        MatIconModule,
        MatInputModule
      ],
      providers: [MatDialog, { provide: TransactionService, useValue: mockTransactionService }]
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
    cy.get('button.add-button').click();
    cy.get('div.transactions-dialog')
      .should('exist')
      .should('be.visible');

    cy.get('body').click(0, 0);
    cy.get('div.transactions-dialog').should('not.exist');
  });

  it('should have 2 entries', () => {
    cy.get('tr.transaction__row').should('have.length', 2);
  })

  it('should have a detail row that pops out', () => {
    cy.get('.transaction__detail-contents').should('not.be.visible');
    cy.get('td>button').eq(0).click();
    cy.get('.transaction__detail-contents').should('be.visible').should('contain', '[none]');
    cy.get('td>button').eq(1).click();
    cy.get('.transaction__detail-contents').should('be.visible').should('contain', "Also paid for my friend's drink.");
  })

})