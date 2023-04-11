import { Component } from '@angular/core';
import { CalendarView, CalendarEvent } from 'angular-calendar';

@Component({
  selector: 'app-dash-calendar',
  templateUrl: './dash-calendar.component.html',
  styleUrls: ['./dash-calendar.component.css']
})
export class DashCalendarComponent {
  viewDate: Date = new Date();
  activeDayIsOpen: boolean = false;
  events: CalendarEvent[] = [
    {
      start: new Date('2023-04-28'),
      title: 'An event',
    }
  ];
  view: CalendarView = CalendarView.Month;

  handleDayClick(event: { date: Date; events: CalendarEvent[] }): void {
    if (event.events.length === 0) {
      return;
    }

    if (!this.activeDayIsOpen) {
      this.activeDayIsOpen = true;
    } else if (this.viewDate.getTime() === event.date.getTime()) {
      this.activeDayIsOpen = false;
    }
    this.viewDate = event.date;
  }
}
