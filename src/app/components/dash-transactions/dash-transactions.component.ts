import { AfterViewInit, Component, Input, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort, Sort } from '@angular/material/sort';
import { MatTable, MatTableDataSource } from '@angular/material/table';
import { animate, state, style, transition, trigger } from '@angular/animations';
import { TransactionService } from 'src/app/transaction.service';
import { Transaction } from 'src/types/transaction-system';
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

  openDialog(): void {
    let dialogRef = this.dialog.open(TransactionsModalComponent, {
      data: {}
    });

    dialogRef.afterClosed().subscribe(res => {
      if (res) {
        this.transactionData = new MatTableDataSource([...this.transactionData.data, res.data]);
        this.transactionData.sort = this.sort;
        this.table.renderRows();
      }
      console.log('The dialog was closed');
      console.log(res.data);
    });
  }

  parseDate(str: string) {
    return new Date(str).toLocaleDateString();
  }

}
