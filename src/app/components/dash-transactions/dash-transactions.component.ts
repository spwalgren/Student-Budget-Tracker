import { Component } from '@angular/core';
import { TransactionService } from 'src/app/transaction.service';
import { Transaction } from 'src/types/transaction-system';

@Component({
  selector: 'app-dash-transactions',
  templateUrl: './dash-transactions.component.html',
  styleUrls: ['./dash-transactions.component.css']
})
export class DashTransactionsComponent {

  transactionData: Transaction[] = [];

  constructor(private transactionService: TransactionService) { }

  ngOnInit() {
    this.transactionService.getTransactions()
      .subscribe((res) => {
        if (!res.err) {
          this.transactionData = res.data;
        }
      })
  }
}
