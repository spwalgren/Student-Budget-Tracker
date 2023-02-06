import { Component } from '@angular/core';
import { LoginService } from 'src/app/login.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {

  yourName: string = "";

  constructor(
    private loginService: LoginService
  ) { }

  ngOnInit() {
    this.loginService.getUserData()
      .subscribe(res => {
        if (res.firstName) {
          this.yourName = res.firstName;
        } else {
          this.yourName = 'ERROR';
        }
      })
  }
}
