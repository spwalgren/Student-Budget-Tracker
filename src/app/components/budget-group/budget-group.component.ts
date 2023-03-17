import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-budget-group',
  templateUrl: './budget-group.component.html',
  styleUrls: ['./budget-group.component.css']
})
export class BudgetGroupComponent {
  @Input()
  groupName: string = '';

}
