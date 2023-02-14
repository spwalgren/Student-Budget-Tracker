describe('login', () => {
  it('Should not login if the form is invalid', () => {
    cy.visit('/')
    cy.get('button').eq(0).click()
    cy.get('[formControlName="email"]').type('profanis@gmail.com');
    cy.url().should('includes', 'login');
    cy.get('button').eq(0).click()
    cy.url().should('not.include', 'dashboard');
  })
})
