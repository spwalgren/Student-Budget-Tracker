import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/login.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {

  yourName: string = "";

  constructor(
    private loginService: LoginService,
    private router: Router
  ) { }

  ngOnInit() {
    this.loginService.getUserData()
      .subscribe(res => {
        if (!res.err) {
          this.yourName = res.firstName;
        } else {
          this.yourName = 'ERROR';
          this.router.navigate(['/login']);
        }
      })
  }

  goLogOut() {
    this.loginService.logOut()
      .subscribe(res => {
        if (!res.err) {
          this.router.navigate(['/login']);
        }
      })
  }
}
