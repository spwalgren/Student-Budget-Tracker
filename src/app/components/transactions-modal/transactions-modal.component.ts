import { Component, Inject } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { TransactionService } from 'src/app/transaction.service';
import { Transaction, CreateTransactionRequest } from 'src/types/transaction-system';


@Component({
  selector: 'app-transactions-modal',
  templateUrl: './transactions-modal.component.html',
  styleUrls: ['./transactions-modal.component.css']
})
export class TransactionsModalComponent {

  transactionForm: FormGroup;

  constructor(
    private transactionService: TransactionService,
    public dialogRef: MatDialogRef<TransactionsModalComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {
    this.transactionForm = new FormGroup({
      name: new FormControl('', [Validators.required]),
      amount: new FormControl(0, [Validators.required]),
      date: new FormControl<Date>(new Date(), [Validators.required]),
      category: new FormControl(''),
      description: new FormControl('')
    });
  } //data b/w data and source

  ngOnInit() { }

  goSubmitTransaction() {
    if (!this.transactionForm.invalid) {
      console.log(this.transactionForm.get("date"));

      const transactionRequest: CreateTransactionRequest = {
        data: {
          name: this.transactionForm.get("name")?.value,
          amount: this.transactionForm.get("amount")?.value,
          date: (this.transactionForm.get("date")?.value as Date).toISOString(),
          category: this.transactionForm.get("category")?.value,
          description: this.transactionForm.get("description")?.value,
        }
      }
      this.transactionService.createTransaction(transactionRequest)
        .subscribe(res => {
          if (!res.err)
            this.dialogRef.close(transactionRequest);
        })
    }
  }
}
