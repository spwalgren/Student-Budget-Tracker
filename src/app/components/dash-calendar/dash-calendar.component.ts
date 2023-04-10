import { Component } from '@angular/core';
import { CalendarView, CalendarEvent } from 'angular-calendar';

@Component({
  selector: 'app-dash-calendar',
  templateUrl: './dash-calendar.component.html',
  styleUrls: ['./dash-calendar.component.css']
})
export class DashCalendarComponent {
  viewDate: Date = new Date();
  events: CalendarEvent[] = [
    {
      start: new Date('2023-04-28'),
      title: 'An event'
    }
  ];
  view: CalendarView = CalendarView.Month;
}
