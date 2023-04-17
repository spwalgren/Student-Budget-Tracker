import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { CalendarModule, DateAdapter } from 'angular-calendar'
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns'
import { DashCalendarComponent } from './dash-calendar.component'
import { CalendarService } from 'src/app/calendar.service'
import { Observable, of } from 'rxjs'
import { GetEventsResponse } from 'src/types/calendar-system'
import { MatButtonModule } from '@angular/material/button'
import { MatIconModule } from '@angular/material/icon'
import { Period } from 'src/types/budget-system'

describe('DashCalendarComponent', () => {

  const mockCalendarService: Partial<CalendarService> = {
    getEvents(month: number): Observable<GetEventsResponse> {
      const startDate = new Date();
      startDate.setDate(startDate.getDate() - 2);
      const endDate = new Date();
      endDate.setDate(endDate.getDate() + 2);
      return of({
        events: [
          {
            userId: 20,
            eventId: 0,
            data: {
              category: 'Test',
              amountLimit: 100,
              totalSpent: 20,
              startDate: startDate.toISOString(),
              endDate: endDate.toISOString(),
              frequency: Period.weekly,
            }
          }
        ]
      })
    }
  }


  beforeEach(() => {
    cy.mount(DashCalendarComponent, {
      imports: [
        CalendarModule.forRoot({
          provide: DateAdapter,
          useFactory: adapterFactory
        }),
        BrowserAnimationsModule,
        MatButtonModule,
        MatIconModule,
      ],
      providers: [
        { provide: CalendarService, useValue: mockCalendarService }
      ]
    })
  })

  it('should mount', () => { })

  it('should display the current month', () => {
    cy.get('.dash-calendar__header').should('contain', new Date().toLocaleString('default', { month: 'long' }))
  })

  it('should have buttons to navigate to the previous and next month', () => {
    const now = new Date()
    const nextMonth = new Date()
    nextMonth.setMonth(now.getMonth() + 1)
    const prevMonth = new Date()
    prevMonth.setMonth(now.getMonth() - 1)

    cy.get('[mwlCalendarPreviousView]').click();
    cy.get('.dash-calendar__header').should('contain', prevMonth.toLocaleString('default', { month: 'long' }))
    cy.get('[mwlCalendarNextView]').click();
    cy.get('.dash-calendar__header').should('contain', now.toLocaleString('default', { month: 'long' }))
    cy.get('[mwlCalendarNextView]').click();
    cy.get('.dash-calendar__header').should('contain', nextMonth.toLocaleString('default', { month: 'long' }))
    cy.get('[mwlCalendarToday]').click();
    cy.get('.dash-calendar__header').should('contain', now.toLocaleString('default', { month: 'long' }))
  })
})