import { Component, Inject } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { TransactionService } from 'src/app/transaction.service';
import { CreateTransactionRequest, TransactionContent } from 'src/types/transaction-system';

interface TransactionModalData {
  data: TransactionContent,
  mode: "Add" | "Edit"
}

@Component({
  selector: 'app-transactions-modal',
  templateUrl: './transactions-modal.component.html',
  styleUrls: ['./transactions-modal.component.css']
})
export class TransactionsModalComponent {

  transactionForm: FormGroup;
  transactionId?: number;
  mode: "Add" | "Edit";

  constructor(
    public dialogRef: MatDialogRef<TransactionsModalComponent, CreateTransactionRequest>,
    @Inject(MAT_DIALOG_DATA) public data: TransactionModalData,
  ) {
    this.transactionForm = new FormGroup({
      name: new FormControl(data.data.name, [Validators.required]),
      amount: new FormControl(data.data.amount, [Validators.required]),
      date: new FormControl<Date>(new Date(data.data.date), [Validators.required]),
      category: new FormControl(data.data.category),
      description: new FormControl(data.data.description)
    });
    this.mode = data.mode;
  } //data b/w data and source

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
      this.dialogRef.close(transactionRequest);
    }
  }
}
