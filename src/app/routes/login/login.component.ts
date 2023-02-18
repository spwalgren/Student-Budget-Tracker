import { Component, SimpleChanges } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/login.service';
import { LogInRequest } from 'src/types/login-system';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  logInForm: FormGroup;
  showAlert = false;
  alertMessage = "";
  alertType: "success" | "info" | "error" | undefined = "error";
  awaitingRes = false;

  constructor(
    private loginService: LoginService,
    private router: Router
  ) {
    this.logInForm = new FormGroup({
      email: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required])
    });
  }

  ngOnInit() {
    this.loginService.getUserData()
      .subscribe(res => {
        if (res.firstName) {
          this.router.navigate(['/dashboard']);
        }
      })
  }

  goSubmitForm() {
    if (!this.logInForm.invalid) {
      const logInRequest: LogInRequest = {
        email: this.logInForm.get('email')?.value.toLowerCase(),
        password: this.logInForm.get('password')?.value
      }
      console.log(logInRequest);
      this.awaitingRes = true;
      this.showAlert = false;
      this.loginService.logIn(logInRequest)
        .subscribe(res => {
          if (!res.err) {
            this.showAlert = true;
            this.alertType = "success";
            this.alertMessage = "Success! Logging you in..."
            setTimeout(() => {
              this.router.navigate(['/dashboard']);
            }, 2000);
          } else {
            this.showAlert = true;
            this.alertMessage = "Could not verify credentials";
            this.awaitingRes = false;
          }
        });
    } else {
      console.log('Invalid entry');
    }
  }
}
