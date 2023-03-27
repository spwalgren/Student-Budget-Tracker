// ***********************************************
// This example namespace declaration will help
// with Intellisense and code completion in your
// IDE or Text Editor.
// ***********************************************
// declare namespace Cypress {
//   interface Chainable<Subject = any> {
//     customCommand(param: any): typeof customCommand;
//   }
// }
//
// function customCommand(param: any): void {
//   console.warn(param);
// }
//
// NOTE: You can use it like so:
// Cypress.Commands.add('customCommand', customCommand);
//
// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })

declare namespace Cypress {
  interface Chainable<Subject = any> {
    logInUser(sampleUserNumber: number, waitForRedirect: boolean): typeof logInUser;
    registerUser(sampleUserNumber: number, waitForRedirect: boolean): typeof registerUser;
    logOutUser(waitForRedirect: boolean): typeof logOutUser;
    deleteUser(waitForRedirect: boolean): typeof deleteUser;
  }
}

function logInUser(sampleUserNumber: number, waitForRedirect = true): void {
  cy.visit('/login');
  cy.get('[formControlName="email"]').type(`sample${sampleUserNumber}@example.com`);
  cy.get('[formControlName="password"]').type('1234');
  cy.get('[data-cy="Submit Login"]').click()
  if (waitForRedirect)
    cy.url({ timeout: 10000 }).should('include', 'dashboard');
}

function registerUser(sampleUserNumber: number, waitForRedirect = true): void {
  cy.visit('/sign-up');
  cy.get('[data-cy="testFirstName"]').type(`firstName`);
  cy.get('[data-cy="testLastName"]').type(`lastName`);
  cy.get('[formControlName="email"]').type(`sample${sampleUserNumber}@example.com`);
  cy.get('[formControlName="password"]').type('1234');
  cy.get('[formControlName="reenteredPass"]').type('1234');
  cy.get('[type="submit"]').click()
  if (waitForRedirect) {
    cy.url({ timeout: 10000 }).should('include', 'login');
  }
}

function logOutUser(waitForRedirect = true): void {
  cy.visit('/dashboard');
  cy.get('[data-cy="log-out-button"]').click();
  if (waitForRedirect)
    cy.url({ timeout: 10000 }).should('include', 'login');
}

function deleteUser(waitForRedirect = true): void {
  cy.visit('/dashboard/settings');
  cy.get('[data-cy="delete-user-btn"]').click();
  cy.get('[data-cy="delete-user-btn"]').click();
  if (waitForRedirect)
    cy.url({ timeout: 10000 }).should('include', 'login');
}

Cypress.Commands.add('logInUser', logInUser);
Cypress.Commands.add('logOutUser', logOutUser);
Cypress.Commands.add('registerUser', registerUser);
Cypress.Commands.add('deleteUser', deleteUser);