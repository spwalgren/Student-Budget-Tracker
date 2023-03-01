import { of } from 'rxjs';
import { LoginService } from 'src/app/login.service';
import { LandingComponent } from './landing.component';
import { MatToolbar, MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';

describe('LandingComponent', () => {

  beforeEach(async () => {
    cy.mount(LandingComponent, {
      imports: [MatToolbarModule, MatButtonModule]
    });
  });

  it('should mount', () => { });


});
