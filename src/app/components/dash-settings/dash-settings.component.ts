import { Component } from '@angular/core';
import { LoginService } from 'src/app/login.service';

@Component({
  selector: 'app-dash-settings',
  templateUrl: './dash-settings.component.html',
  styleUrls: ['./dash-settings.component.css']
})
export class DashSettingsComponent {

  isDeleting = false;

  constructor(private loginService: LoginService) { }

  deleteUser() {
    if (this.isDeleting) {

    } else {
      this.isDeleting = true;
      setTimeout(() => { this.isDeleting = false; }, 5000);
    }
  }
}
