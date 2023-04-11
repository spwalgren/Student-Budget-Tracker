import { Component, Input } from '@angular/core';
import { Period } from 'src/types/budget-system';
import { EventContent } from 'src/types/calendar-system';

@Component({
  selector: 'app-event-card',
  templateUrl: './event-card.component.html',
  styleUrls: ['./event-card.component.css']
})
export class EventCardComponent {

  @Input()
  eventData: EventContent = {
    "frequency": Period.weekly,
    "startDate": "2023-04-11T04:00:00Z",
    "endDate": "2023-04-18T03:59:59Z",
    "category": "General",
    "totalSpent": 120,
    "amountLimit": 100
  };
  numberFormatter = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  });
  statusText: string;

  constructor() {
    if (this.eventData.startDate > new Date().toISOString()) {
      this.statusText = "Upcoming";
    } else if (this.eventData.totalSpent > this.eventData.amountLimit) {
      this.statusText = "Overspent";
    } else {
      this.statusText = "On Track";
    }
  }

  getDueDateString(): string {
    return new Date(this.eventData.endDate).toLocaleDateString();
  }

  getDaysLeft(): number {
    const today = new Date();
    const endDate = new Date(this.eventData.endDate);
    const timeDiff = endDate.getTime() - today.getTime();
    return Math.floor(timeDiff / (1000 * 3600 * 24));
  }

}
