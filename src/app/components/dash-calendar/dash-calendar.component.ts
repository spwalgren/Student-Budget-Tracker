import { Component } from '@angular/core';
import { CalendarView, CalendarEvent } from 'angular-calendar';
import { CalendarService } from 'src/app/calendar.service';
import { Event } from 'src/types/calendar-system';

@Component({
  selector: 'app-dash-calendar',
  templateUrl: './dash-calendar.component.html',
  styleUrls: ['./dash-calendar.component.css']
})
export class DashCalendarComponent {
  viewDate: Date = new Date();
  activeDayIsOpen: boolean = false;
  events: Event[] = [];
  calendarEvents: CalendarEvent[] = [
    {
      start: new Date('2023-04-28'),
      title: 'An event',
    }
  ];
  view: CalendarView = CalendarView.Month;

  constructor(private calendarService: CalendarService) { }

  ngOnInit() {
    this.calendarService.getEvents(this.viewDate.getMonth()).subscribe(res => {
      console.log(res);

      if (!res.err) {
        this.events = res.events;
        this.calendarEvents = this.events.map(event => {
          return {
            start: new Date(event.data.endDate),
            title: event.data.category,
          }
        });
      }
    });
  }

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

  closeActiveDay(): void {
    this.activeDayIsOpen = false;
  }
}
