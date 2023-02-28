import { LoginService } from 'src/app/login.service'
import { DashboardComponent } from './dashboard.component'
import { Router } from '@angular/router'
import { HttpClientModule } from '@angular/common/http'
import { MatButtonModule } from '@angular/material/button'
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatSidenavModule } from '@angular/material/sidenav'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'
import { of } from 'rxjs'

describe('DashHomeComponent', () => {

  const mockLoginService: Partial<LoginService> = {
    getUserData() {
      return of({
        id: "0",
        email: "sample1@example.com",
        firstName: "firstName1",
        lastName: "lastName1"
      });
    },
    logOut() {
      return of({});
    }
  }

  beforeEach(() => {
    cy.spy(mockLoginService, 'getUserData').as('getUser');
    cy.mount(DashboardComponent, {
      imports: [HttpClientModule, MatButtonModule, MatToolbarModule, MatSidenavModule, BrowserAnimationsModule],
      providers: [Router, { provide: LoginService, useValue: mockLoginService }]
    }).then((wrapper) => {
      cy.stub(wrapper.component, 'goLogOut').as('goLogOut');
    })
  });

  it('should mount', () => { });

  it('should have visible buttons', () => {
    cy.get('mat-sidenav button').eq(0)
      .should('exist')
      .should('be.visible')
      .should('contain.text', 'Home');
    cy.get('mat-toolbar button')
      .should('exist')
      .should('be.visible')
      .should('contain.text', 'Log Out');
  });

  it('should immediately call a function', () => {
    cy.get('@getUser').should('have.been.called');
  })

  it('should call a function when logging out', () => {
    cy.get('@goLogOut').should('not.have.been.called');
    cy.get('mat-toolbar button').click();
    cy.get('@goLogOut').should('have.been.called');
  })
})