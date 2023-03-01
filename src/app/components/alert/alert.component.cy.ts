import { MatIconModule } from "@angular/material/icon";
import { AlertComponent } from "./alert.component";

describe('AlertComponent', () => {

  it('should mount', () => {
    cy.mount(AlertComponent, {
      imports: [MatIconModule]
    });
  });

  it('should display text', () => {
    cy.mount(AlertComponent, {
      componentProperties: {
        message: "This is an error"
      },
      imports: [MatIconModule]
    });
    cy.get('div').contains('This is an error');
  });

  it('should display different colors', () => {
    cy.mount(AlertComponent, {
      componentProperties: {
        type: "success"
      },
      imports: [MatIconModule]
    });
    cy.get('div.alert').should('have.css', 'background-color', 'rgb(237, 247, 237)').should('have.css', 'color', 'rgb(28, 68, 30)');
    cy.mount(AlertComponent, {
      componentProperties: {
        type: "info"
      },
      imports: [MatIconModule]
    });
    cy.get('div.alert').should('have.css', 'background-color', 'rgb(229, 246, 253)').should('have.css', 'color', 'rgb(1, 67, 97)');
    cy.mount(AlertComponent, {
      componentProperties: {
        type: "error"
      },
      imports: [MatIconModule]
    });
    cy.get('div.alert').should('have.css', 'background-color', 'rgb(253, 237, 237)').should('have.css', 'color', 'rgb(95, 33, 32)');
    cy.mount(AlertComponent);
    cy.get('div.alert').should('have.css', 'background-color', 'rgb(253, 237, 237)').should('have.css', 'color', 'rgb(95, 33, 32)');
  });

  it('should display an icon', () => {
    cy.mount(AlertComponent, {
      componentProperties: {
        type: "success"
      },
      imports: [MatIconModule]
    });
    cy.get('mat-icon').should('exist');
  })
});