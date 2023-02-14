describe('My First Test', () => {
  it('Visits the initial landing page from login', () => {
    cy.visit('/')
    cy.contains('app is running!')
  })
})
