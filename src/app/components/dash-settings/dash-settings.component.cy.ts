import { LoginService } from 'src/app/login.service';
import { DashSettingsComponent } from './dash-settings.component'
import { Observable, of } from 'rxjs';
import { GenericResponse } from 'src/types/api-system';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('DashSettingsComponent', () => {

  const mockLoginService: Partial<LoginService> = {
    deleteUser(): Observable<GenericResponse> {
      return of({ err: "Cannot delete user during test" });
    }
  };

  beforeEach(() => {
    cy.mount(DashSettingsComponent, {
      imports: [
        MatButtonModule,
        MatIconModule,
        MatInputModule,
        BrowserAnimationsModule
      ],
      providers: [
        { provide: LoginService, useValue: mockLoginService }
      ]
    }).then((wrapper) => {
      cy.spy(wrapper.component, 'deleteUser').as('deleteUserComp');
      cy.spy(wrapper.component.loginService, 'deleteUser').as('deleteUserServ');

      return cy.wrap(wrapper).as('angular');
    });
  })

  it('should mount', () => {
  })

  it('should have a button to delete the user', () => {
    cy.get('[data-cy="delete-user-btn"]').should('exist').should('be.visible')
    cy.get('[data-cy="delete-warning"]').should('not.exist');
    cy.get('[data-cy="delete-user-btn"]').click();
    cy.get('[data-cy="delete-user-btn"]').click();
    cy.get('@deleteUserComp').should('have.been.called');
    cy.get('@deleteUserServ').should('have.been.called');
  })

  it('should have a warning message that disappears', () => {
    cy.get('[data-cy="delete-user-btn"]').should('exist').should('be.visible')
    cy.get('[data-cy="delete-warning"]').should('not.exist');
    cy.get('[data-cy="delete-user-btn"]').click();
    cy.get('@deleteUserComp').should('have.been.called');
    cy.get('@deleteUserServ').should('not.have.been.called');
    cy.get('[data-cy="delete-warning"]').should('exist').should('be.visible');
    cy.wait(5500);
    cy.get('[data-cy="delete-warning"]').should('not.exist');
  })
})