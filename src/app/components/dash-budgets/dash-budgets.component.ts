import { Component } from '@angular/core';
import { Budget } from 'src/types/budget-system';
import { BudgetService } from 'src/app/budget.service';

@Component({
  selector: 'app-dash-budgets',
  templateUrl: './dash-budgets.component.html',
  styleUrls: ['./dash-budgets.component.css']
})
export class DashBudgetsComponent {

  budgetData: Budget[] = []
  existingCategories: string[] = []

  constructor(private budgetService: BudgetService) { }

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
}
