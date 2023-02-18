describe('login', () => {
  it('Should not login if the form is invalid', () => {
    cy.visit('/')
    cy.get('[data-cy="Log In"]').click()
    cy.get('[formControlName="email"]').type('bob@gmail.com');
    cy.url().should('includes', 'login');
    cy.get('[data-cy="Submit Login"]').click()
    cy.url().should('not.include', 'dashboard');
  })
})
