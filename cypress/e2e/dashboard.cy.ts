describe('dashboard', () => {
  it('should open when user is logged in', () => {
    cy.visit('/login');
    cy.logInUser(1);
    cy.url({ timeout: 10000 }).should('contain', 'dashboard');
    cy.visit('/login');
    cy.url().should('contain', 'dashboard');
  })

  it('should have links to different parts', () => {
    cy.logInUser(1);
    cy.url({ timeout: 10000 }).should('contain', 'dashboard');
    cy.get('mat-sidenav button').eq(0).click();
    cy.url().should('contain', '/dashboard');
    cy.get('mat-sidenav button').eq(1).click();
    cy.url().should('contain', '/dashboard/transactions');
    cy.get('mat-sidenav button').eq(2).click();
    cy.url().should('contain', '/dashboard/goals');
    cy.get('mat-sidenav button').eq(3).click();
    cy.url().should('contain', '/dashboard/calendar');
    cy.get('mat-sidenav button').eq(4).click();
    cy.url().should('contain', '/dashboard/settings');
  })

  it('should have a 404 page', () => {
    cy.logInUser(1);
    cy.url({ timeout: 10000 }).should('contain', 'dashboard');
    cy.visit('/somethingrandom');
    cy.url().should('contain', '/somethingrandom');
    cy.get('app-page-not-found').should('exist');
    cy.get('button').click();
    cy.url().should('contain', 'dashboard');
  })

  it('should let the user log out', () => {
    cy.logInUser(1);
    cy.url({ timeout: 10000 }).should('contain', 'dashboard');
    cy.get('mat-toolbar button').click();
    cy.url().should('contain', 'login');
    cy.go('back');
    cy.url().should('contain', 'login');
    cy.visit('/dashboard');
    cy.url().should('contain', 'login');
  })
})