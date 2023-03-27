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

})