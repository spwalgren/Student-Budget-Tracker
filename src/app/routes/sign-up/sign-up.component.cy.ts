import { Router } from '@angular/router';
import { SignUpComponent } from './sign-up.component';
import { LoginService } from 'src/app/login.service';
import { of } from 'rxjs';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatToolbarModule } from '@angular/material/toolbar';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AlertComponent } from 'src/app/components/alert/alert.component';
import { SignUpRequest } from 'src/types/login-system';

describe('SignUpComponent', () => {

  const mockLoginService: Partial<LoginService> = {
    signUp(signUpRequest: SignUpRequest) {
      return of({ err: "Not allowed in test" })
    }
  }

  beforeEach(() => {
    cy.spy(mockLoginService, 'signUp').as('signUp');

    cy.mount(SignUpComponent, {
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
      providers: [
        Router,
        { provide: LoginService, useValue: mockLoginService }
      ]
    }).then((wrapper) => {
      cy.spy(wrapper.component, 'goSubmitForm').as('goSubmit');
      return cy.wrap(wrapper).as('angular');
    })
  });

  it('should mount', () => { });

  it('should let a person sign up', () => {
    cy.get('[data-cy="testFirstName"]').type('firstName1');
    cy.get('[data-cy="testLastName"]').type('lastName1');
    cy.get('.sign-up__email').type('somethingrandom');
    cy.get('.sign-up__password').type('1234');
    cy.get('.sign-up__reenter-pass').type('1234');
    cy.get('[data-cy="Submit Sign-Up"]').click();
    cy.get('@goSubmit').should('have.been.called');
    cy.get('@signUp').should('not.have.been.called');

    cy.get('[formControlName="email"]').clear()
    cy.get('.sign-up__email').type('sample1@example.com');
    cy.get('[formControlName="reenteredPass"]').clear()
    cy.get('.sign-up__reenter-pass').type('12345');
    cy.get('[data-cy="Submit Sign-Up"]').click();
    cy.get('@goSubmit').should('have.been.called');
    cy.get('@signUp').should('not.have.been.called');

    cy.get('[formControlName="reenteredPass"]').clear()
    cy.get('.sign-up__reenter-pass').type('1234');
    cy.get('[data-cy="Submit Sign-Up"]').click();
    cy.get('@goSubmit').should('have.been.called');
    cy.get('@signUp').should('have.been.called');
    cy.get('app-alert').should('be.visible');
  })
});
