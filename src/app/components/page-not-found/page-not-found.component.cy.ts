import { createOutputSpy } from 'cypress/angular';
import { PageNotFoundComponent } from './page-not-found.component'
import { MatButtonModule } from '@angular/material/button';

describe('PageNotFoundComponent', () => {
  it('should mount', () => {
    cy.mount(PageNotFoundComponent);
  })

  it('should have a button', () => {
    cy.mount(PageNotFoundComponent, {
      imports: [MatButtonModule]
    });
    cy.get('button').should('contain', 'Go Back');
  })
})