import { Component, SimpleChanges } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { LoginService } from 'src/app/login.service';
import User from 'src/types/User';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  logInForm: FormGroup;
  users: User[] = [];
  showAlert = false;
  alertMessage = "";
  awaitingRes = false;

  constructor(
    private loginService: LoginService
  ) {
    this.logInForm = new FormGroup({
      email: new FormControl('', [Validators.required]),
      password: new FormControl('', [Validators.required])
    });
  }

  // ngOnInit() {
  //   this.loginService.getUsers()
  //     .subscribe(users => this.users = users);
  // }

  // ngDoCheck() {
  //   console.log(this.users);
  // }

  goSubmitForm() {
    if (!this.logInForm.invalid) {
      console.log({
        email: this.logInForm.get('email')?.value,
        password: this.logInForm.get('password')?.value
      });
      this.awaitingRes = true;
      this.showAlert = false;
      this.loginService.createAuth(this.logInForm.get('email')?.value, this.logInForm.get('password')?.value)
        .subscribe(_ => {
          this.showAlert = true;
          this.alertMessage = "Could not verify credentials";
          this.awaitingRes = false;
        });
    } else {
      console.log('Invalid entry');
    }
  }
}
