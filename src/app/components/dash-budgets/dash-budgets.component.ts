import { Component, Inject } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Budget, BudgetContent, CreateBudgetRequest, Period, UpdateBudgetRequest } from 'src/types/budget-system';
import { BudgetService } from 'src/app/budget.service';
import { MAT_DIALOG_DATA, MatDialog, MatDialogRef } from '@angular/material/dialog';

export interface BudgetsDialogData {
  mode: "Add" | "Edit",
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
  budgetForm!: FormGroup;
  budgetId?: number;
  mode!: "Add" | "Edit";

  constructor(
    private budgetService: BudgetService,
    public dialog: MatDialog
  ) { }

  ngOnInit() {
    this.rerenderBudgets()
  }

  rerenderBudgets() {
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

  openAddDialog(): void {
    const dialogRef = this.dialog.open(BudgetsDialogComponent, {
      data: {
        mode: "Add", data: {
          userId: 0,
          budgetId: 0,
          data: {
            category: "General",
            amountLimit: 100,
            frequency: Period.weekly,
            duration: 1,
            startDate: new Date().toISOString()
          }
        }
      } as BudgetsDialogData
    });

    dialogRef.afterClosed().subscribe((res?: BudgetsDialogData) => {
      if (res) {
        const budgetRequest: CreateBudgetRequest = {
          ...res.data.data
        }
        this.budgetService.createBudget(budgetRequest).subscribe(_ => {
          this.rerenderBudgets();
        })
      }
    });
  }

  openEditDialog(budget: Budget): void {
    const dialogRef = this.dialog.open(BudgetsDialogComponent, {
      data: {
        mode: "Edit", data: {
          ...budget
        }
      } as BudgetsDialogData
    });

    dialogRef.afterClosed().subscribe((res?: BudgetsDialogData) => {
      if (res) {
        const budgetRequest: UpdateBudgetRequest = {
          newBudget: res.data,
        }
        this.budgetService.updateBudget(budgetRequest).subscribe(_ => {
          this.rerenderBudgets();
        })
      }
    });
  }
}

@Component({
  selector: 'budgets-dialog',
  templateUrl: 'budgets-dialog.html',
  styleUrls: ['./dash-budgets.component.css']
})
export class BudgetsDialogComponent {

  budgetForm: FormGroup;
  frequencyOptions = ['weekly', 'monthly', 'yearly'];
  mode: "Add" | "Edit";


  constructor(
    public dialogRef: MatDialogRef<BudgetsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: BudgetsDialogData
  ) {
    this.budgetForm = new FormGroup({
      category: new FormControl(data.data.data.category),
      amount: new FormControl(data.data.data.amountLimit, [Validators.required]),
      frequency: new FormControl(data.data.data.frequency, [Validators.required]),
      duration: new FormControl(data.data.data.duration),
      repeats: new FormControl(data.data.data.count ? true : false, [Validators.required]),
      count: new FormControl(data.data.data.count),
      startDate: new FormControl(data.data.data.startDate)
    });
    this.mode = data.mode
  }

  onNoClick(): void {
    // console.log('Budget Name:', this.budgetName.value);
    // console.log('Budget Category:', this.budgetCategory.value);
    // console.log('Budget Amount:', this.budgetAmount.value);
    // console.log('Frequency:', this.frequency.value);
    this.dialogRef.close();
  }

  // onFrequencyChange(event: any) {
  //   console.log(event.target.value);
  // }

  goSubmitBudget() {
    if (!this.budgetForm.invalid) {

      const repeats: boolean = this.budgetForm.get('repeats')?.value as boolean
      const budgetContent: BudgetContent = {
        category: this.budgetForm.get('category')?.value as string,
        amountLimit: this.budgetForm.get('amount')?.value as number,
        frequency: this.budgetForm.get('frequency')?.value as Period,
        duration: this.budgetForm.get('duration')?.value as number,
        count: repeats ? this.budgetForm.get('count')?.value as number : undefined,
        startDate: this.budgetForm.get('startDate')?.value as string,
      }
      this.dialogRef.close(budgetContent);
    }
  }

  getDurationFrequency() {
    return this.budgetForm.get('frequency')?.value;
  }

  isRepeating() {
    return this.budgetForm.get('repeats')?.value;
  }
}
