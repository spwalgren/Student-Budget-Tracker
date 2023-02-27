import { MAT_DIALOG_DATA, MatDialog, MatDialogModule, MatDialogRef } from '@angular/material/dialog'
import { TransactionsModalComponent } from './transactions-modal.component'
import { TransactionService } from 'src/app/transaction.service'

describe('TransactionsModalComponent', () => {
  it('should mount', () => {
    cy.mount(TransactionsModalComponent, {
      imports: [MatDialogModule],
    })
  })
})