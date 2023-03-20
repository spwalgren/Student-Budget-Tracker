import { Component, Input } from '@angular/core';
import { BudgetService } from 'src/app/budget.service';
import { Budget, BudgetContent, Period } from 'src/types/budget-system';

@Component({
  selector: 'app-budget-group',
  templateUrl: './budget-group.component.html',
  styleUrls: ['./budget-group.component.css']
})
export class BudgetGroupComponent {
  @Input()
  groupName: string = '';
  @Input()
  budgetData: Budget[] = []
  displayedColumns = [
    "amountLimit",
    "period",
    "startDate",
    "currentPeriod",
    "editAndDelete"
  ]

  numberFormatter = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD'
  })

  getPeriodDef(budgetContent: BudgetContent) {
    const frequency = budgetContent.frequency;
    const duration = budgetContent.duration;
    const count = budgetContent.count;
    let result = "Every ";
    let frequencyString =
      frequency === Period.monthly
        ? "month"
        : frequency === Period.weekly
          ? "week"
          : frequency === Period.yearly
            ? "year"
            : "period"
    if (duration > 1) {
      result += `${duration} ${frequencyString}s`;
    } else {
      result += `${frequencyString}`;
    }

    if (count) {
      result += ` for ${count} ${frequencyString}`;
      if (count > 1) {
        result += 's';
      }
    }
    return result;
  }

  parseDate(str: string) {
    return new Date(str).toLocaleDateString();
  }

  getPeriod(budgetContent: BudgetContent): { periodStart: Date, periodEnd: Date, daysLeft: number } | null {
    const startDate = new Date(budgetContent.startDate);
    const today = new Date();

    let addOne: (current: Date) => Date;
    if (budgetContent.frequency === Period.monthly) {
      addOne = (current) => {
        current.setMonth(current.getMonth() + 1);
        return current;
      }
    } else if (budgetContent.frequency === Period.yearly) {
      addOne = (current) => {
        current.setFullYear(current.getFullYear() + 1);
        return current;
      }
    } else if (budgetContent.frequency === Period.weekly) {
      addOne = (current) => {
        current.setDate(current.getDate() + 7);
        return current;
      }
    } else {
      addOne = (current) => {
        current.setDate(current.getDate() + 1);
        return current;
      }
    }

    let periodStart = new Date(startDate);
    let periodEnd = new Date(startDate);
    for (let i = 0; i < budgetContent.duration; i++) {
      periodEnd = addOne(periodEnd);
    }


    let periodsPassed = 0;

    while (periodEnd < today && (!budgetContent.count || periodsPassed < budgetContent.count)) {
      for (let i = 0; i < budgetContent.duration; i++) {
        periodStart = addOne(periodStart);
        periodEnd = addOne(periodEnd);
      }
      periodsPassed++;

    }

    if (budgetContent.count && periodsPassed >= budgetContent.count) {
      return null;
    }
    return ({
      periodStart: periodStart,
      periodEnd: periodEnd,
      daysLeft: Math.floor((periodEnd.getTime() - today.getTime()) / 86400000)
    })
  }

  getPeriodString(budgetContent: BudgetContent) {
    const currentPeriod = this.getPeriod(budgetContent);
    if (currentPeriod === null) {
      return "Expired";
    }
    return `${currentPeriod.periodStart.toLocaleDateString()} 
    to ${currentPeriod.periodEnd.toLocaleDateString()}
    (${currentPeriod.daysLeft} day(s) left)`
  }


}
