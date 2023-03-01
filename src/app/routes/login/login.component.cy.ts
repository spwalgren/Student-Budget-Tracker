import { Router } from '@angular/router';
import { LoginComponent } from './login.component';
import { LoginService } from 'src/app/login.service';
import { HttpClientModule } from '@angular/common/http';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { of } from 'rxjs';
import { AlertComponent } from 'src/app/components/alert/alert.component';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatIconModule } from '@angular/material/icon';

describe('LoginComponent', () => {

  const mockLoginService: Partial<LoginService> = {
    getUserData() {
      return of({
        err: "Could not login",
        id: "0",
        email: "sample1@example.com",
        firstName: "firstName1",
        lastName: "lastName1"
      });
    },
    logIn() {
      return of({ err: "Not allowed in test" });
    }
  }

  beforeEach(() => {

    cy.spy(mockLoginService, 'getUserData').as('getUser');
    cy.spy(mockLoginService, 'logIn').as('logIn');
    cy.mount(LoginComponent, {
      declarations: [AlertComponent],
      imports: [
        HttpClientModule,
        MatToolbarModule,
        MatButtonModule,
        MatInputModule,
        BrowserAnimationsModule,
        MatCardModule,
        MatFormFieldModule,
        MatProgressSpinnerModule,
        FormsModule,
        ReactiveFormsModule,
        MatIconModule],
      providers: [Router, { provide: LoginService, useValue: mockLoginService }]
    }).then((wrapper) => {
      cy.spy(wrapper.component, 'goSubmitForm').as('goSubmit');

      return cy.wrap(wrapper).as('angular');
    });
  });

  it('should mount', () => { });

  it('should attempt to automatically login', () => {
    cy.get('@getUser').should('have.been.called');
  })

  it('should login with a valid form', () => {
    cy.get('mat-form-field.login__email').type('sample1@example.com');
    cy.get('[data-cy="Submit Login"]').click();
    cy.get('@goSubmit').should('have.been.calledOnce');
    cy.get('@logIn').should('not.have.been.called');
    cy.get('[data-cy="alert-component"]').should('not.exist');

    cy.get('mat-form-field.login__password').type('1234');
    cy.get('.login__email [formControlName="email"]').clear();
    cy.get('mat-form-field.login__email').type('somethingrandom');
    cy.get('[data-cy="Submit Login"]').click();
    cy.get('@goSubmit').should('have.been.calledTwice');
    cy.get('@logIn').should('not.have.been.called');
    cy.get('[data-cy="alert-component"]').should('not.exist');

    cy.get('.login__email [formControlName="email"]').clear();
    cy.get('mat-form-field.login__email').type('sample1@example.com');
    cy.get('[data-cy="Submit Login"]').click();
    cy.get('@goSubmit').should('have.been.calledThrice');
    cy.get('@logIn').should('have.been.called');
    cy.get('[data-cy="alert-component"]')
      .should('exist');
  })
});
