describe('login', () => {
  it('Should not login if the form is invalid', () => {
    cy.visit('/')
    cy.get('button').eq(0).click()
    cy.url().should('includes', 'login');
  })
})
