import { LoginService } from 'src/app/login.service'
import { DashboardComponent } from './dashboard.component'
import { Router } from '@angular/router'
import { HttpClientModule } from '@angular/common/http'
import { MatButtonModule } from '@angular/material/button'
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatSidenavModule } from '@angular/material/sidenav'
import { BrowserAnimationsModule } from '@angular/platform-browser/animations'

describe('DashHomeComponent', () => {
  it('should mount', () => {
    cy.mount(DashboardComponent, {
      imports: [HttpClientModule, MatButtonModule, MatToolbarModule, MatSidenavModule, BrowserAnimationsModule],
      providers: [LoginService, Router]
    })
  })
})