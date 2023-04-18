describe('budgets', () => {

  it('should be accessible', () => {
    cy.registerUser(101, true);
    cy.logInUser(101, true);

    cy.visit('/dashboard/budgets');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[data-cy="cancel-btn"]').click();

    cy.deleteUser(true);
  });

  it('should be able to store data', () => {
    cy.registerUser(101, true);
    cy.logInUser(101, true);

    cy.visit('/dashboard/budgets');
    cy.get('table').should('not.exist');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table.General').should('exist');
    cy.get('table.General tr').should('have.length', 2);
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table.General tr').should('have.length', 3);
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table.General tr').should('have.length', 4);
    cy.logOutUser(true);
    cy.logInUser(101, true);
    cy.visit('/dashboard/budgets');
    cy.get('table.General tr').should('have.length', 4);

    cy.deleteUser(true);
  });

  it('should change the list of available categories', () => {
    cy.registerUser(101, true);
    cy.logInUser(101, true);

    cy.visit('/dashboard/transactions');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option')
      .should('have.length', 1)
      .should('contain', '[None]');
    cy.get('body').type('{esc}');

    cy.visit('/dashboard/budgets');
    cy.get('table').should('not.exist');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table.General').should('exist');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="category"]').clear().type('Food');
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table.Food').should('exist');

    cy.visit('/dashboard/transactions');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option')
      .should('have.length', 3)
      .should('contain', '[None]')
      .should('contain', 'General')
      .should('contain', 'Food');
    cy.get('mat-option').eq(1).click();
    cy.get('[formControlName="name"]').type('Stuff');
    cy.get('[formControlName="amount"]').type('10');
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table tr').eq(1).should('contain.text', 'Stuff');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="name"]').type('Things');
    cy.get('[formControlName="amount"]').type('20');
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option').eq(2).click();
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table tr').eq(3).should('contain.text', 'Things');

    cy.visit('/dashboard/budgets');
    cy.get('[data-cy="edit-btn"]').eq(1).click();
    cy.get('[formControlName="category"]').clear().type('Rent');
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table.Rent').should('exist');
    cy.get('[data-cy="delete-btn"]').eq(0).click();
    cy.get('table.Food').should('not.exist');

    cy.visit('/dashboard/transactions');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="category"]').click();
    cy.get('mat-option')
      .should('have.length', 2)
      .should('contain', '[None]')
      .should('contain', 'Rent');
    cy.get('body').type('{esc}');
    cy.get('body').type('{esc}');
    cy.get('table tr').eq(1).should('contain.text', 'Rent');
    cy.get('table tr').eq(3).should('contain.text', '[None]');

    cy.deleteUser(true);
  });
})