describe('template spec', () => {

  beforeEach(() => {

    cy.logInUser(1);
    cy.visit('/dashboard/transactions');
    cy.get('tbody').then(($tbody) => {
      if ($tbody.children('.transaction__row').length > 0) {
        cy.get('.transaction__row').each(($elem) => {
          cy.wrap($elem).click();
          cy.get('.transaction__detail-delete').eq(0).click();
          cy.wrap($elem).should('not.exist');
        });
      }
    })
    cy.logOutUser();

    cy.logInUser(2);
    cy.visit('/dashboard/transactions');
    cy.get('tbody').then(($tbody) => {
      if ($tbody.children('.transaction__row').length > 0) {
        cy.get('.transaction__row').each(($elem) => {
          cy.wrap($elem).click();
          cy.get('.transaction__detail-delete').eq(0).click();
          cy.wrap($elem).should('not.exist');
        });
      }
    })
    cy.logOutUser();
  })

  it('passes', () => {
    cy.visit('https://example.cypress.io');
  })
})