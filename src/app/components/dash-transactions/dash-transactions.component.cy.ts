import { MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog'
import { DashTransactionsComponent } from './dash-transactions.component'
import { MatButtonModule } from '@angular/material/button'
import { MatTableModule } from '@angular/material/table'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { MatIconModule } from '@angular/material/icon'
import { MatInputModule } from '@angular/material/input'

describe('DashTransactionsComponent', () => {

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
      providers: [MatDialog]
    })
  });

  it('should mount', () => {
    cy.mount(DashTransactionsComponent, {
      imports: [
        MatDialogModule,
        MatButtonModule,
        MatTableModule,
        BrowserAnimationsModule,
        MatIconModule,
        MatInputModule
      ],
      providers: [MatDialog]
    })
  });

  it('should have a table', () => {
    cy.get('table[mat-table]').should('exist');
  });

  it('should have a button that opens a dialog', () => {

    cy.get('div.transaction-modal').should('not.exist');
    cy.get('button.add-button').click();
    cy.get('div.transaction-modal')
      .should('exist')
      .should('be.visible');

    cy.get('body').click(0, 0);
    cy.get('div.transaction-modal').should('not.exist');
  });

  it('should have a wrapper', () => {
    cy.mount(DashTransactionsComponent, {
      imports: [
        MatDialogModule,
        MatButtonModule,
        MatTableModule,
        BrowserAnimationsModule,
        MatIconModule,
        MatInputModule
      ],
      providers: [MatDialog]
    }).then((wrapper) => {
      console.log(wrapper);

    })
  })


})