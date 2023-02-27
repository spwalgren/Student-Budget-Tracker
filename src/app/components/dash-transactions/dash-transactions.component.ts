import { AfterViewInit, Component, Input, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort, Sort } from '@angular/material/sort';
import { MatTable, MatTableDataSource } from '@angular/material/table';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { TransactionService } from 'src/app/transaction.service';
import { EditTransactionRequest, Transaction } from 'src/types/transaction-system';
import { MatDialog } from '@angular/material/dialog'
import { TransactionsModalComponent } from '../transactions-modal/transactions-modal.component';


@Component({
  selector: 'app-dash-transactions',
  templateUrl: './dash-transactions.component.html',
  styleUrls: ['./dash-transactions.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0' })),
      state('expanded', style({ height: '*' })),
      transition('expanded <=> collapsed', animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')),
    ]),
  ],
})
export class DashTransactionsComponent {

  transactionData: MatTableDataSource<Transaction>;
  displayedColumns = ['name', 'amount', 'category', 'date', 'expand'];
  expandedRow: Transaction | null = null;
  isChanging: boolean = false;

  constructor(private transactionService: TransactionService, public dialog: MatDialog) {
    this.transactionData = new MatTableDataSource<Transaction>([]);
  }

  @ViewChild(MatTable) table!: MatTable<Transaction>;
  @ViewChild(MatSort) sort!: MatSort;

  ngAfterViewInit() {
    this.transactionData.sort = this.sort;
  }

  ngOnInit() {
    this.transactionService.getTransactions()
      .subscribe((res) => {
        if (!res.err) {
          this.transactionData = new MatTableDataSource(res.data);
          this.transactionData.sort = this.sort;
        }
      })
  }

  openAddDialog(): void {
    // Open the dialog and pass it blank data
    let dialogRef = this.dialog.open(TransactionsModalComponent, {
      data: {
        data: {
          name: '',
          amount: 0,
          date: new Date().toISOString(),
          category: '',
          description: '',
        },
        mode: "Add"
      }
    });

    // When the dialog closes...
    dialogRef.afterClosed().subscribe(dialogRes => {
      // If the dialog returned data...
      if (dialogRes) {
        // Create the data in the database
        this.transactionService.createTransaction(dialogRes)
          .subscribe();
        // Add the data to the table
        this.transactionData = new MatTableDataSource([...this.transactionData.data, dialogRes.data]);
        this.transactionData.sort = this.sort;
        this.table.renderRows();
      }
      console.log('The dialog was closed');
      // console.log(dialogRes.data);
    });
  }

  openEditDialog(transaction: Transaction) {
    if (!this.isChanging) {
      this.isChanging = true;
      // Open the dialog and pass it the current data
      let dialogRef = this.dialog.open(TransactionsModalComponent, {
        data: {
          data: transaction,
          mode: "Edit"
        },
      });
      // When the dialog closes...
      dialogRef.afterClosed().subscribe(dialogRes => {
        // If the dialog has returned data...
        if (dialogRes) {
          const req: EditTransactionRequest = {
            data: dialogRes.data
          }
          this.transactionService.editTransaction(req)
            .subscribe(res => {
              if (!res.err) {
                const targetIndex = this.transactionData.data.findIndex((elem) => elem.transactionId == transaction.transactionId);
                this.transactionData.data.splice(targetIndex, 1, dialogRes.data);
                this.transactionData = new MatTableDataSource([...this.transactionData.data]);
                this.transactionData.sort = this.sort;
                this.table.renderRows();
                this.isChanging = false;
              }
            });
        }
        console.log('The dialog was closed');
        console.log(dialogRes);
      });
    }
  }

  goDeleteTransaction(transaction: Transaction) {

    if (!this.isChanging) {
      this.isChanging = true;
      this.transactionService.deleteTransaction(transaction.transactionId!)
        .subscribe(res => {
          if (!res.err) {
            const targetIndex = this.transactionData.data.findIndex((elem) => elem.transactionId == transaction.transactionId);
            this.transactionData.data.splice(targetIndex, 1);
            this.transactionData = new MatTableDataSource([...this.transactionData.data]);
            this.transactionData.sort = this.sort;
            this.table.renderRows();
            this.isChanging = false;
          }
        });
    }
  }

  parseDate(str: string) {
    return new Date(str).toLocaleDateString();
  }

}
