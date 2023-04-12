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

    this.calendarService.getEvents(0).subscribe(res => {
      console.log(res);
      if (!res.err) {
        this.events.push(...res.events);
        this.events.sort((a, b) => a.data.endDate > b.data.endDate ? 1 : -1);
      }
      this.calendarEvents = this.events.map(event => {
        return {
          start: new Date(event.data.endDate),
          title: event.data.category,
        }
      });
    });
    this.calendarService.getEvents(1).subscribe(res => {
      console.log(res);
      if (!res.err) {
        this.events.push(...res.events);
        this.events.sort((a, b) => a.data.endDate > b.data.endDate ? 1 : -1);
      }
      this.calendarEvents = this.events.map(event => {
        return {
          start: new Date(event.data.endDate),
          title: event.data.category,
        }
      });
    });
    // this.calendarService.getEvents(2).subscribe(res => {
    //   console.log(res);
    //   if (!res.err) {
    //     this.events.push(...res.events);
    //     this.events.sort((a, b) => a.data.endDate > b.data.endDate ? 1 : -1);
    //   }
    //   this.calendarEvents = this.events.map(event => {
    //     return {
    //       start: new Date(event.data.endDate),
    //       title: event.data.category,
    //     }
    //   });
    //   console.log(this.events);

    // });
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

  handleEventClick(eventContent?: EventContent): void {
    if (eventContent) {
      this.viewDate = new Date(eventContent.endDate);
      this.activeDayIsOpen = true;
    }
  }
}
