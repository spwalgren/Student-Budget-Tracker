import { Period } from 'src/types/budget-system'
import { EventCardComponent } from './event-card.component'

describe('EventCardComponent', () => {

  it('should mount', () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 2);
    const endDate = new Date();
    endDate.setDate(endDate.getDate() + 2);
    cy.mount(EventCardComponent, {
      componentProperties: {
        eventData: {
          category: 'Test',
          amountLimit: 100,
          totalSpent: 0,
          startDate: startDate.toISOString(),
          endDate: endDate.toISOString(),
          frequency: Period.weekly,
        },
      }
    })
  })

  it('should be on track', () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 2);
    const endDate = new Date();
    endDate.setDate(endDate.getDate() + 2);
    cy.mount(EventCardComponent, {
      componentProperties: {
        eventData: {
          category: 'Test',
          amountLimit: 100,
          totalSpent: 20,
          startDate: startDate.toISOString(),
          endDate: endDate.toISOString(),
          frequency: Period.weekly,
        },
      }
    })
    cy.get('h4').should('contain', 'Test')
    cy.get('.event-card').should('contain.text', 'On Track')
    cy.get('.event-card').should('contain.text', `Due ${endDate.toLocaleDateString()}`)
    cy.get('.event-card').should('contain.text', '1 day(s) left')
    cy.get('.event-card').should('contain.text', 'Spending limit: $100.00')
    cy.get('.event-card').should('contain.text', 'Left to spend: $80.00')
  })

  it('should be upcoming', () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() + 2);
    const endDate = new Date();
    endDate.setDate(endDate.getDate() + 4);
    cy.mount(EventCardComponent, {
      componentProperties: {
        eventData: {
          category: 'Test',
          amountLimit: 100,
          totalSpent: 0,
          startDate: startDate.toISOString(),
          endDate: endDate.toISOString(),
          frequency: Period.weekly,
        },
      }
    })
    cy.get('h4').should('contain', 'Test')
    cy.get('.event-card').should('contain.text', 'Upcoming')
    cy.get('.event-card').should('contain.text', `Due ${endDate.toLocaleDateString()}`)
    cy.get('.event-card').should('contain.text', '3 day(s) left')
    cy.get('.event-card').should('contain.text', 'Spending limit: $100.00')
    cy.get('.event-card').should('not.contain.text', 'Left to spend: $100.00')
  })

  it('should be completed', () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 4);
    const endDate = new Date();
    endDate.setDate(endDate.getDate() - 2);
    cy.mount(EventCardComponent, {
      componentProperties: {
        eventData: {
          category: 'Test',
          amountLimit: 100,
          totalSpent: 30,
          startDate: startDate.toISOString(),
          endDate: endDate.toISOString(),
          frequency: Period.weekly,
        },
      }
    })
    cy.get('h4').should('contain', 'Test')
    cy.get('.event-card').should('contain.text', 'Completed')
    cy.get('.event-card').should('contain.text', `Due ${endDate.toLocaleDateString()}`)
    cy.get('.event-card').should('not.contain.text', 'day(s) left')
    cy.get('.event-card').should('contain.text', 'Spending limit: $100.00')
    cy.get('.event-card').should('contain.text', 'Amount spent: $30.00')
  })

  it('should be overspent', () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 2);
    const endDate = new Date();
    endDate.setDate(endDate.getDate() + 2);
    cy.mount(EventCardComponent, {
      componentProperties: {
        eventData: {
          category: 'Test',
          amountLimit: 100,
          totalSpent: 120,
          startDate: startDate.toISOString(),
          endDate: endDate.toISOString(),
          frequency: Period.weekly,
        },
      }
    })
    cy.get('h4').should('contain', 'Test')
    cy.get('.event-card').should('contain.text', 'Overspent')
    cy.get('.event-card').should('contain.text', `Due ${endDate.toLocaleDateString()}`)
    cy.get('.event-card').should('contain.text', '1 day(s) left')
    cy.get('.event-card').should('contain.text', 'Spending limit: $100.00')
    cy.get('.event-card').should('contain.text', 'Amount spent: $120.00')
  })

  it('should emit event when clicked', () => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - 2);
    const endDate = new Date();
    endDate.setDate(endDate.getDate() + 2);
    cy.mount(EventCardComponent, {
      componentProperties: {
        eventData: {
          category: 'Test',
          amountLimit: 100,
          totalSpent: 0,
          startDate: startDate.toISOString(),
          endDate: endDate.toISOString(),
          frequency: Period.weekly,
        },
      }
    }).then((component) => {
      cy.spy(component.component, 'emitEventContent').as('event')
    })
    cy.get('.event-card').click()
    cy.get('@event').should('have.been.calledOnce')
  })
})