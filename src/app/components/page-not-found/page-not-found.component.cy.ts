import { createOutputSpy } from 'cypress/angular';
import { PageNotFoundComponent } from './page-not-found.component'
import { MatButtonModule } from '@angular/material/button';

describe('PageNotFoundComponent', () => {
  it('should mount', () => {
    cy.mount(PageNotFoundComponent);
  });

  it('should have a button', () => {
    cy.mount(PageNotFoundComponent, {
      imports: [MatButtonModule]
    });
    cy.get('button').should('contain', 'Go Back');
  });

  it('should call a function', () => {
    cy.mount(PageNotFoundComponent, {
      imports: [MatButtonModule]
    }).then((wrapper) => {
      cy.stub(wrapper.component, 'goBackToPreviousPage').as('goBack');
      return cy.wrap(wrapper).as('angular')
    });
    cy.get('button').click();
    cy.get('@goBack').should('have.been.called');
  })
})