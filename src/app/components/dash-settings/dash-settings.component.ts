import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/login.service';

@Component({
  selector: 'app-dash-settings',
  templateUrl: './dash-settings.component.html',
  styleUrls: ['./dash-settings.component.css']
})
export class DashSettingsComponent {

  isDeleting = false;

  constructor(private loginService: LoginService, private router: Router) { }

  deleteUser() {
    if (this.isDeleting) {
      this.loginService.deleteUser().subscribe(() => {
        this.router.navigate(['/login']);
      });
    } else {
      this.isDeleting = true;
      setTimeout(() => { this.isDeleting = false; }, 5000);
    }
  }
}
