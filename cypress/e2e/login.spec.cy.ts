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

  it('Should take the user to Sign up page upon clicking \"Sign Up\" Button', () => {
    cy.visit('/')
    cy.get('[data-cy="Log In"]').click()
    cy.get('[formControlName="email"]').type('bob@gmail.com');
    cy.get('[formControlName="password"]').type('bobbyIsTheBest');
    cy.url().should('includes', 'login');
    cy.get('[data-cy="Sign-Up"]').click()
    cy.url().should('include', 'sign-up');
  })

  it('Should let user sign up... clicking \"Sign Up\" from login page', () => {
    cy.visit('/')
    cy.get('[data-cy="Log In"]').click()
    cy.url().should('includes', 'login');
    cy.get('[data-cy="Sign-Up"]').click({waitForAnimations: false})
    cy.url().should('include', 'sign-up');
    cy.get('[data-cy="testFirstName"]').type('Bobby');
    cy.get('[data-cy="testLastName"]').type('Fergison');
    cy.get('[formControlName="email"]').type('bobby@gmail.com');
    cy.get('[formControlName="password"]').type('bobbyIsTheBest');
    cy.get('[formControlName="reenteredPass"]').type('bobbyIsTheBest');
    cy.url().should('includes', 'sign-up');
    cy.get('[data-cy="Submit Sign-Up"]', { timeout: 16000 }).click();
    cy.url().should('include', 'login');

  })
  
  // it('Should let user sign up... clicking \"Sign Up\" from dashboard', () => {
  //   cy.visit('/')
  //   cy.get('[data-cy="Sign Up"]').click({force:true})
  //   cy.url().should('include', 'sign-up');
  //   cy.get('[formControlName="firstName"]').type('Bobby', timeout: 20);
  //   cy.get('[formControlName="lastName"]').type('Fergison');
  //   cy.get('[formControlName="email"]').type('bobby@gmail.com');
  //   cy.get('[formControlName="password"]').type('bobbyIsTheBest');
  //   cy.get('[formControlName="reenteredPass"]').type('bobbyIsTheBest');
  //   cy.url().should('includes', 'sign-up');
  //   cy.get('[data-cy="Submit Sign-Up"]').click({force:true});
  //   cy.url().should('include', 'dashboard');
  // })

  //test uncompleted sign up


})
