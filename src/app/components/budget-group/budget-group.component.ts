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

  budgetData: Budget[] = []
  displayedColumns = [
    "amountLimit",
    "period"
  ]

  numberFormatter = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD'
  })

  constructor(private budgetService: BudgetService) { }

  ngOnInit() {
    this.budgetService.getBudgets()
      .subscribe((res) => {
        if (!res.err) {
          this.budgetData = [...res.budgets];
        }
      })
  }

  getPeriodString(budgetContent: BudgetContent) {
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
}
