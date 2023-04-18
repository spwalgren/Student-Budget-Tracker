import { Component } from '@angular/core';
import { CalendarView, CalendarEvent } from 'angular-calendar';
import { firstValueFrom, from } from 'rxjs';
import { CalendarService } from 'src/app/calendar.service';
import { Period } from 'src/types/budget-system';
import { Event, EventContent } from 'src/types/calendar-system';

@Component({
  selector: 'app-dash-calendar',
  templateUrl: './dash-calendar.component.html',
  styleUrls: ['./dash-calendar.component.css']
})
export class DashCalendarComponent {
  viewDate: Date = new Date();
  activeDayIsOpen: boolean = false;
  events: Event[] = [];
  calendarEvents: CalendarEvent[] = [];
  view: CalendarView = CalendarView.Month;

  constructor(private calendarService: CalendarService) { }

  ngOnInit() {
    this.updateCalendarEvents();
  }

  updateCalendarEvents() {

    const yearOffset = this.viewDate.getFullYear() - new Date().getFullYear();
    const monthOffset = (this.viewDate.getMonth() - new Date().getMonth()) + (yearOffset * 12);
    if (monthOffset < 0) {
      this.events = [];
      this.calendarEvents = [];
      return;
    }

    this.calendarService.getEvents(monthOffset).subscribe(res => {
      if (!res.err) {
        this.events = [...res.events];
        this.events.sort((a, b) => a.data.endDate > b.data.endDate ? 1 : -1);
      }
      this.calendarEvents = this.events.map(event => {
        return {
          start: new Date(event.data.endDate),
          title: `${event.data.category}; Amount spent: ${event.data.totalSpent}/${event.data.amountLimit}`,
        }
      });
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
    this.updateCalendarEvents();
  }

  handleEventClick(eventContent?: EventContent): void {
    if (eventContent) {
      this.viewDate = new Date(eventContent.endDate);
      this.activeDayIsOpen = true;
    }
  }
}
