describe('transactions', () => {

  it('should be accessible', () => {
    cy.registerUser(101, true);
    cy.logInUser(101, true);

    cy.visit('/dashboard/transactions');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[data-cy="cancel-btn"]').click();

    cy.deleteUser(true);
  })

  it('should be able to store data', () => {
    cy.registerUser(101, true);
    cy.logInUser(101, true);

    cy.visit('/dashboard/transactions');
    cy.get('table').should('not.exist');
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="name"]').type('Basics');
    cy.get('[formControlName="amount"]').type('10');
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table tr').should('have.length', 3);
    cy.get('[data-cy="add-btn"]').click();
    cy.get('[formControlName="name"]').type('Basics');
    cy.get('[formControlName="amount"]').type('10');
    cy.get('[data-cy="submit-btn"]').click();
    cy.get('table tr').should('have.length', 5);
    cy.logOutUser(true);
    cy.logInUser(101, true);
    cy.visit('/dashboard/transactions');
    cy.get('table tr').should('have.length', 5);

    cy.deleteUser(true);
  })
})