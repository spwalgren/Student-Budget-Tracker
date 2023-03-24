import { Component, Inject } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Budget } from 'src/types/budget-system';
import { BudgetService } from 'src/app/budget.service';
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from '@angular/material/dialog';

export interface BudgetsDialogData {
  mode: "add" | "edit",
  data: Budget
}

@Component({
  selector: 'app-dash-budgets',
  templateUrl: './dash-budgets.component.html',
  styleUrls: ['./dash-budgets.component.css']
})
export class DashBudgetsComponent {

  budgetData: Budget[] = []
  existingCategories: string[] = []

  constructor(
    private budgetService: BudgetService,
    public dialog: MatDialog
  ) { }

  ngOnInit() {
    this.budgetService.getBudgets()
      .subscribe((res) => {
        if (!res.err) {
          this.budgetData = [...res.budgets];
          this.budgetData.forEach((elem) => {
            if (!this.existingCategories.find((str) => str === elem.data.category)) {
              this.existingCategories.push(elem.data.category);
            }
          })
        }
      })
  }

  getFilteredData(category: string, budgetData: Budget[]) {
    return budgetData.filter((elem) => elem.data.category === category);
  }

  openDialog(): void {
    const dialogRef = this.dialog.open(BudgetsDialogComponent, {
      data: { mode: "add", data: null }
    });

    dialogRef.afterClosed().subscribe((res) => {
      console.log(res);
    });
  }
}

@Component({
  selector: 'budgets-dialog',
  templateUrl: 'budgets-dialog.html',
  styleUrls: ['./dash-budgets.component.css']
})
export class BudgetsDialogComponent {
  constructor(
    public dialogRef: MatDialogRef<BudgetsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: BudgetsDialogData
  ) { }

  onNoClick(): void {
    this.dialogRef.close();
  }
}