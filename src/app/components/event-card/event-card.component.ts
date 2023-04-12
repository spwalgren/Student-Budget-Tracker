import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Period } from 'src/types/budget-system';
import { EventContent } from 'src/types/calendar-system';

@Component({
  selector: 'app-event-card',
  templateUrl: './event-card.component.html',
  styleUrls: ['./event-card.component.css']
})
export class EventCardComponent {

  @Input()
  eventData?: EventContent;
  numberFormatter = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  });
  statusText: string = "";

  @Output()
  clickEvent = new EventEmitter<EventContent | undefined>();
  emitEventContent(eventContent?: EventContent) {
    this.clickEvent.emit(eventContent);
  }

  constructor() {
    console.log(this.eventData);
    this.getStatusText();
  }

  getStatusText(): string {
    if (!this.eventData) return "Loading";
    if (this.eventData.startDate > new Date().toISOString()) {
      this.statusText = "Upcoming";
    } else if (this.eventData.totalSpent > this.eventData.amountLimit) {
      this.statusText = "Overspent";
    } else {
      this.statusText = "On Track";
    }
    return this.statusText;
  }

  getDueDateString(): string {
    if (!this.eventData) return "";
    return new Date(this.eventData.endDate).toLocaleDateString();
  }

  getDaysLeft(): number {
    if (!this.eventData) return 0;
    const today = new Date();
    const endDate = new Date(this.eventData.endDate);
    const timeDiff = endDate.getTime() - today.getTime();
    return Math.floor(timeDiff / (1000 * 3600 * 24));
  }

}
