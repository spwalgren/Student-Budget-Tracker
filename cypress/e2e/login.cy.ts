describe('login', () => {
  it('Should not login if only username is inputted', () => {
    cy.visit('/')
    cy.get('[data-cy="Log In"]').click()
    cy.get('[formControlName="email"]').type('bob@gmail.com');
    cy.url().should('includes', 'login');
    cy.get('[data-cy="Submit Login"]').click()
    cy.url().should('not.include', 'dashboard');
  })

  it('Should  not login if username and password are not recognized', () => {
    cy.visit('/')
    cy.get('[data-cy="Log In"]').click()
    cy.get('[formControlName="email"]').type('bob@gmail.com');
    cy.get('[formControlName="password"]').type('b0btheb3st');
    cy.url().should('includes', 'login');
    cy.get('[data-cy="Submit Login"]').click()
    cy.url().should('not.include', 'dashboard');
  })

  it('Should take the user to Sign up page upon clicking "Sign Up" Button', () => {
    cy.visit('/')
    cy.get('[data-cy="Log In"]').click()
    cy.get('[formControlName="email"]').type('bob@gmail.com');
    cy.get('[formControlName="password"]').type('bobbyIsTheBest');
    cy.url().should('includes', 'login');
    cy.get('[data-cy="Sign-Up"]').click()
    cy.url().should('include', 'sign-up');
  })

  it('Should let user sign up and delete their user', () => {
    cy.logInUser(101, false);
    cy.get('[data-cy="alert-component"]').should('exist').should('contain.text', 'Could not verify credentials');
    cy.registerUser(101, true);
    cy.logInUser(101, true);
    cy.deleteUser(true);
    cy.logInUser(101, false);
    cy.get('[data-cy="alert-component"]').should('exist').should('contain.text', 'Could not verify credentials');
  })
})
