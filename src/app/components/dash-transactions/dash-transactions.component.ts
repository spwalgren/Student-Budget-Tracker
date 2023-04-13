import { AfterViewInit, Component, Inject, Input, ViewChild } from '@angular/core';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort, Sort } from '@angular/material/sort';
import { MatTable, MatTableDataSource } from '@angular/material/table';
import {
  animate,
  state,
  style,
  transition,
  trigger,
} from '@angular/animations';
import { TransactionService } from 'src/app/transaction.service';
import {
  UpdateTransactionRequest,
  Transaction,
  CreateTransactionRequest,
  TransactionContent,
} from 'src/types/transaction-system';
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from '@angular/material/dialog';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BudgetService } from 'src/app/budget.service';

@Component({
  selector: 'app-dash-transactions',
  templateUrl: './dash-transactions.component.html',
  styleUrls: ['./dash-transactions.component.css'],
  animations: [
    trigger('detailExpand', [
      state('collapsed', style({ height: '0px', minHeight: '0' })),
      state('expanded', style({ height: '*' })),
      transition(
        'expanded <=> collapsed',
        animate('225ms cubic-bezier(0.4, 0.0, 0.2, 1)')
      ),
    ]),
  ],
})
export class DashTransactionsComponent {
  transactionData: Transaction[] = [];
  transactionTableData: MatTableDataSource<Transaction>;
  displayedColumns = ['name', 'amount', 'category', 'date', 'editAndDelete', 'expand'];
  expandedRow: Transaction | null = null;
  isChanging: boolean = false;
  categoryOptions = ['[None]'];

  constructor(
    public transactionService: TransactionService,
    public budgetService: BudgetService,
    public dialog: MatDialog
  ) {
    this.transactionTableData = new MatTableDataSource<Transaction>([]);
  }

  @ViewChild(MatTable) table!: MatTable<Transaction>;
  @ViewChild(MatSort) sort!: MatSort;

  ngAfterViewInit() {
    this.transactionTableData.sort = this.sort;
  }

  ngOnInit() {
    this.transactionService.getTransactions().subscribe((res) => {
      if (!res.err) {
        this.transactionData = [...res.data];
        this.transactionTableData = new MatTableDataSource(this.transactionData);
        this.transactionTableData.sort = this.sort;
      }
    });

    this.budgetService.getBudgetCategories().subscribe((res) => {
      if (!res.err) {
        this.categoryOptions.push(...res.categories);
      }
    });
  }

  rerenderTable() {
    this.transactionTableData = new MatTableDataSource([
      ...this.transactionTableData.data,
    ]);
    this.transactionTableData.sort = this.sort;
    this.table.renderRows();
    this.isChanging = false;
  }

  rerenderTransactions() {
    this.transactionService.getTransactions().subscribe((res) => {
      if (!res.err) {
        this.transactionData = [...res.data];
        this.rerenderTable();
      }
    });
  }

  openAddDialog(): void {
    // Open the dialog and pass it blank data
    let dialogRef = this.dialog.open(TransactionsDialogComponent, {
      data: {
        data: {
          name: '',
          amount: 0,
          date: this.getToday(),
          category: '[None]',
          description: '',
        },
        mode: 'Add',
        categoryOptions: this.categoryOptions,
      },
    });

    // When the dialog closes...
    dialogRef.afterClosed().subscribe((dialogRes) => {
      const dialogOutput = dialogRes as CreateTransactionRequest;
      // If the dialog returned data...
      if (dialogOutput) {
        // Create the data in the database
        this.transactionService
          .createTransaction(dialogOutput)
          .subscribe((_) => {
            this.rerenderTransactions();
          });
      }
      console.log('The dialog was closed');
    });
  }

  openEditDialog(transaction: Transaction) {
    if (!this.isChanging) {
      this.isChanging = true;
      // Open the dialog and pass it the current data
      let dialogRef = this.dialog.open(TransactionsDialogComponent, {
        data: {
          data: transaction,
          mode: 'Edit',
          categoryOptions: this.categoryOptions,
        },
      });
      // When the dialog closes...
      dialogRef.afterClosed().subscribe((dialogRes) => {
        const dialogOutput = dialogRes as CreateTransactionRequest;
        // If the dialog has returned data...
        if (dialogOutput) {
          const req: UpdateTransactionRequest = {
            data: {
              userId: transaction.userId,
              transactionId: transaction.transactionId,
              ...dialogOutput.data,
            },
          };
          this.transactionService.updateTransaction(req).subscribe((_) => {
            this.rerenderTransactions();
          });
        } else {
          this.isChanging = false;
        }
        console.log('The dialog was closed');
      });
    }
  }

  goDeleteTransaction(transaction: Transaction) {
    if (!this.isChanging) {
      this.isChanging = true;
      this.transactionService
        .deleteTransaction(transaction)
        .subscribe((_) => {
          this.rerenderTransactions();
        });
    }
  }

  parseDate(str: string) {
    return new Date(str).toLocaleDateString();
  }

  getToday(): string {
    let today = new Date();
    today.setMinutes(today.getMinutes() - today.getTimezoneOffset());
    let todayString = today.toISOString().split('T')[0] + 'T04:00:00.000Z';
    today = new Date(todayString);
    return today.toISOString();
  }
}

interface TransactionsDialogData {
  data: TransactionContent,
  mode: "Add" | "Edit",
  categoryOptions: string[],
}

@Component({
  selector: 'transactions-dialog',
  templateUrl: 'transactions-dialog.html',
  styleUrls: ['./dash-transactions.component.css']
})
export class TransactionsDialogComponent {

  transactionForm: FormGroup;
  transactionId?: number;
  categoryOptions: string[];
  mode: "Add" | "Edit";

  constructor(
    public dialogRef: MatDialogRef<TransactionsDialogComponent, CreateTransactionRequest>,
    @Inject(MAT_DIALOG_DATA) public data: TransactionsDialogData,
  ) {
    this.transactionForm = new FormGroup({
      name: new FormControl(data.data.name, [Validators.required]),
      amount: new FormControl(data.data.amount, [Validators.required]),
      date: new FormControl<Date>(new Date(data.data.date), [Validators.required]),
      category: new FormControl(data.data.category, [Validators.required]),
      description: new FormControl(data.data.description)
    });
    this.mode = data.mode;
    this.categoryOptions = data.categoryOptions;
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
