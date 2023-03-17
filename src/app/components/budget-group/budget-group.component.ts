import { Component, Input } from '@angular/core';
import { BudgetService } from 'src/app/budget.service';
import { Budget } from 'src/types/budget-system';

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
    "amountLimit"
  ]

  constructor(private budgetService: BudgetService) { }

  ngOnInit() {
    this.budgetService.getBudgets()
      .subscribe((res) => {
        if (!res.err) {
          this.budgetData = [...res.budgets];
        }
      })
  }
}
