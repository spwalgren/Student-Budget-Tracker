import { AfterViewInit, Component, Input, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort, Sort } from '@angular/material/sort';
import { MatTable, MatTableDataSource } from '@angular/material/table';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { TransactionService } from 'src/app/transaction.service';
import { UpdateTransactionRequest, Transaction, CreateTransactionRequest } from 'src/types/transaction-system';
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

  constructor(
    public transactionService: TransactionService,
    public dialog: MatDialog) {
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

  rerenderTable() {
    this.transactionData = new MatTableDataSource([...this.transactionData.data]);
    this.transactionData.sort = this.sort;
    this.table.renderRows();
    this.isChanging = false;
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
      const dialogOutput = dialogRes as CreateTransactionRequest;
      // If the dialog returned data...
      if (dialogOutput) {
        // Create the data in the database
        this.transactionService.createTransaction(dialogOutput)
          .subscribe(res => {
            // Add the data to the table
            const newTransaction: Transaction = {
              userId: res.userId,
              transactionId: res.transactionId,
              ...dialogOutput.data
            }
            this.transactionData = new MatTableDataSource([...this.transactionData.data, newTransaction]);
            this.transactionData.sort = this.sort;
            this.table.renderRows();
          });
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
        const dialogOutput = dialogRes as CreateTransactionRequest;
        // If the dialog has returned data...
        if (dialogOutput) {
          const req: UpdateTransactionRequest = {
            data: {
              userId: transaction.userId,
              transactionId: transaction.transactionId,
              ...dialogOutput.data
            }
          }
          this.transactionService.updateTransaction(req)
            .subscribe(res => {
              if (!res.err) {
                const targetIndex = this.transactionData.data.findIndex((elem) => elem.transactionId == transaction.transactionId);
                this.transactionData.data.splice(targetIndex, 1, req.data);
                this.rerenderTable();
              }
            });
        }
        console.log('The dialog was closed');
      });
    }
  }

  goDeleteTransaction(transaction: Transaction) {

    if (!this.isChanging) {
      this.isChanging = true;
      this.transactionService.deleteTransaction(transaction)
        .subscribe(res => {
          if (!res.err) {
            const targetIndex = this.transactionData.data.findIndex((elem) => elem.transactionId == transaction.transactionId);
            this.transactionData.data.splice(targetIndex, 1);
            this.rerenderTable();
          }
        });
    }
  }

  parseDate(str: string) {
    return new Date(str).toLocaleDateString();
  }

}
