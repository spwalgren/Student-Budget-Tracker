import { Component, Inject } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import {
  Budget,
  BudgetContent,
  CreateBudgetRequest,
  Period,
  UpdateBudgetRequest,
} from 'src/types/budget-system';
import { BudgetService } from 'src/app/budget.service';
import {
  MAT_DIALOG_DATA,
  MatDialog,
  MatDialogRef,
} from '@angular/material/dialog';

export interface BudgetsDialogData {
  mode: 'Add' | 'Edit';
  data: Budget;
}

@Component({
  selector: 'app-dash-budgets',
  templateUrl: './dash-budgets.component.html',
  styleUrls: ['./dash-budgets.component.css'],
})
export class DashBudgetsComponent {
  budgetData: Budget[] = [];
  existingCategories: string[] = [];
  budgetForm!: FormGroup;
  budgetId?: number;
  mode!: 'Add' | 'Edit';
  numberFormatter = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  });
  displayedColumns = [
    'amountLimit',
    'period',
    'startDate',
    'currentPeriod',
    'editAndDelete',
  ];
  isDeleting: boolean = false;

  constructor(private budgetService: BudgetService, public dialog: MatDialog) { }

  ngOnInit() {
    this.rerenderBudgets();
  }

  rerenderBudgets() {
    this.budgetService.getBudgets().subscribe((res) => {
      if (!res.err) {
        this.budgetData = [...res.budgets];
        this.existingCategories = [];
        this.budgetData.forEach((elem) => {
          if (
            !this.existingCategories.find((str) => str === elem.data.category)
          ) {
            this.existingCategories.push(elem.data.category);
          }
        });
        this.existingCategories.sort();
      }
      this.isDeleting = false;
    });
  }

  getPeriodDef(budgetContent: BudgetContent) {
    const frequency = budgetContent.frequency;
    const duration = budgetContent.duration;
    const count = budgetContent.count;
    let result = 'Every ';
    let frequencyString =
      frequency === Period.monthly
        ? 'month'
        : frequency === Period.weekly
          ? 'week'
          : frequency === Period.yearly
            ? 'year'
            : 'period';
    if (duration > 1) {
      result += `${duration} ${frequencyString}s`;
    } else {
      result += `${frequencyString}`;
    }

    if (count) {
      result += ` for ${count * duration} ${frequencyString}`;
      if (count > 1) {
        result += 's';
      }
    }
    return result;
  }

  parseDate(str: string) {
    return new Date(str).toLocaleDateString();
  }

  getPeriodString(budgetContent: BudgetContent) {
    const currentPeriod = this.getPeriod(budgetContent);
    if (currentPeriod === null) {
      return 'Expired';
    }
    return `${currentPeriod.periodStart.toLocaleDateString()} 
    to ${currentPeriod.periodEnd.toLocaleDateString()}
    (${currentPeriod.daysLeft} day(s) left)`;
  }

  getPeriod(
    budgetContent: BudgetContent
  ): { periodStart: Date; periodEnd: Date; daysLeft: number } | null {
    const startDate = new Date(budgetContent.startDate);
    const today = new Date(this.getToday());

    let addOne: (current: Date) => Date;
    if (budgetContent.frequency === Period.monthly) {
      addOne = (current) => {
        current.setMonth(current.getMonth() + 1);
        return current;
      };
    } else if (budgetContent.frequency === Period.yearly) {
      addOne = (current) => {
        current.setFullYear(current.getFullYear() + 1);
        return current;
      };
    } else if (budgetContent.frequency === Period.weekly) {
      addOne = (current) => {
        current.setDate(current.getDate() + 7);
        return current;
      };
    } else {
      addOne = (current) => {
        current.setDate(current.getDate() + 1);
        return current;
      };
    }

    let periodStart = new Date(startDate);
    let periodEnd = new Date(startDate);
    for (let i = 0; i < budgetContent.duration; i++) {
      periodEnd = addOne(periodEnd);
    }

    let periodsPassed = 0;

    while (
      periodEnd < today &&
      (!budgetContent.count || periodsPassed < budgetContent.count)
    ) {
      for (let i = 0; i < budgetContent.duration; i++) {
        periodStart = addOne(periodStart);
        periodEnd = addOne(periodEnd);
      }
      periodsPassed++;
    }

    if (budgetContent.count && periodsPassed >= budgetContent.count) {
      return null;
    }
    return {
      periodStart: periodStart,
      periodEnd: periodEnd,
      daysLeft: Math.floor((periodEnd.getTime() - today.getTime()) / 86400000),
    };
  }

  getFilteredData(category: string, budgetData: Budget[]): Budget[] {
    return budgetData.filter((elem) => elem.data.category === category);
  }

  openAddDialog(): void {
    const dialogRef = this.dialog.open(BudgetsDialogComponent, {
      data: {
        mode: 'Add',
        data: {
          userId: 0,
          budgetId: 0,
          data: {
            category: 'General',
            amountLimit: 100,
            frequency: Period.weekly,
            duration: 1,
            startDate: this.getToday(),
          },
        },
      } as BudgetsDialogData,
    });

    dialogRef.afterClosed().subscribe((res?: Budget) => {
      if (res) {
        const budgetRequest: CreateBudgetRequest = {
          ...res.data,
        };
        this.budgetService.createBudget(budgetRequest).subscribe((_) => {
          this.rerenderBudgets();
        });
      }
    });
  }

  openEditDialog(budget: Budget): void {
    const dialogRef = this.dialog.open(BudgetsDialogComponent, {
      data: {
        mode: 'Edit',
        data: {
          ...budget,
        },
      } as BudgetsDialogData,
    });

    dialogRef.afterClosed().subscribe((res?: Budget) => {
      if (res) {
        const budgetRequest: UpdateBudgetRequest = {
          newBudget: res,
        };
        this.budgetService.updateBudget(budgetRequest).subscribe((_) => {
          this.rerenderBudgets();
        });
      }
    });
  }

  deleteBudget(budget: Budget) {
    this.isDeleting = true;
    this.budgetService.deleteBudget(budget.budgetId).subscribe((_) => {
      this.rerenderBudgets();
    });
  }

  getToday(): string {
    let today = new Date();
    today.setMinutes(today.getMinutes() - today.getTimezoneOffset());
    let todayString = today.toISOString().split('T')[0] + 'T04:00:00.000Z';
    today = new Date(todayString);
    return today.toISOString();
  }
}

@Component({
  selector: 'budgets-dialog',
  templateUrl: 'budgets-dialog.html',
  styleUrls: ['./dash-budgets.component.css'],
})
export class BudgetsDialogComponent {
  budgetForm: FormGroup;
  frequencyOptions = ['weekly', 'monthly', 'yearly'];
  mode: 'Add' | 'Edit';

  constructor(
    public dialogRef: MatDialogRef<BudgetsDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: BudgetsDialogData
  ) {
    this.budgetForm = new FormGroup({
      category: new FormControl(data.data.data.category),
      amount: new FormControl(data.data.data.amountLimit, [
        Validators.required,
      ]),
      frequency: new FormControl(data.data.data.frequency, [
        Validators.required,
      ]),
      duration: new FormControl(data.data.data.duration),
      repeats: new FormControl(data.data.data.count ? true : false, [
        Validators.required,
      ]),
      count: new FormControl(data.data.data.count),
      startDate: new FormControl(data.data.data.startDate),
    });
    this.mode = data.mode;
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
      const repeats: boolean = this.budgetForm.get('repeats')?.value as boolean;
      const budget: Budget = {
        userId: this.data.data.userId,
        budgetId: this.data.data.budgetId,
        data: {
          category: this.budgetForm.get('category')?.value as string,
          amountLimit: this.budgetForm.get('amount')?.value as number,
          frequency: this.budgetForm.get('frequency')?.value as Period,
          duration: this.budgetForm.get('duration')?.value as number,
          count: repeats
            ? (this.budgetForm.get('count')?.value as number)
            : undefined,
          startDate: this.budgetForm.get('startDate')?.value as string,
        },
      };
      this.dialogRef.close(budget);
    }
  }

  getDurationFrequency() {
    return this.budgetForm.get('frequency')?.value;
  }

  isRepeating() {
    return this.budgetForm.get('repeats')?.value;
  }
}
