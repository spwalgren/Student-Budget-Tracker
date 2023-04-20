describe('template spec', () => {

  it('should be accessible', () => {
    cy.registerUser(101, true);
    cy.logInUser(101, true);

    cy.visit('/dashboard/calendar');
    cy.get('mwl-calendar-month-view').should('exist');
    cy.get('app-event-card').should('have.length', 0);

    cy.deleteUser(true);
  });

  it('should be able to have events', () => {
    cy.registerUser(101, true);
    cy.logInUser(101, true);

    cy.visit('/dashboard/calendar');
    cy.get('mwl-calendar-month-view').should('exist');
    cy.get('app-event-card').should('have.length', 0);

    cy.visit('/dashboard/budgets');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table.General').should('exist');

    cy.visit('/dashboard/calendar');
    cy.get('.cal-open-day-events')
      .should('not.exist');
    cy.get('app-event-card').should('have.length.greaterThan', 0);
    // cy.get('app-event-card .event-card--upcoming').should('have.length.greaterThan', 0);
    cy.get('app-event-card .event-card--ontrack')
      .should('have.length', 1)
      .click();
    cy.get('.cal-open-day-events')
      .should('exist')
      .should('contain.text', 'General; Amount spent: 0/100');

    cy.visit('/dashboard/transactions');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="name"]').clear().type('Test');
    cy.get('[formControlName="amount"]').clear().type('20');
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option').eq(1).click();
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table tr').eq(1).should('contain.text', 'Test');

    cy.visit('/dashboard/calendar');
    cy.get('.cal-open-day-events')
      .should('not.exist');
    cy.get('app-event-card .event-card--ontrack')
      .should('have.length', 1)
      .should('contain.text', 'Left to spend: $80.00')
      .click();
    cy.get('.cal-open-day-events')
      .should('exist')
      .should('contain.text', 'General; Amount spent: 20/100');

    cy.deleteUser(true);
  });

})