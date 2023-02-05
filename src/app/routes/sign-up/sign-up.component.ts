import { Component } from '@angular/core';
import {
  FormGroup,
  FormControl,
  Validators,
  AbstractControl,
  FormGroupDirective,
  NgForm,
  ValidatorFn,
  ValidationErrors,
} from '@angular/forms';
import { ErrorStateMatcher } from '@angular/material/core';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/login.service';

export class PasswordChecker implements ErrorStateMatcher {
  isErrorState(
    control: AbstractControl<any, any> | null,
    form: FormGroupDirective | NgForm | null
  ): boolean {
    return control?.parent?.errors?.['passwordMismatch'];
  }
}

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css'],
})
export class SignUpComponent {
  signUpForm: FormGroup;
  passwordChecker = new PasswordChecker();
  showAlert = false;
  alertMessage = '';
  awaitingRes = false;
  alertType?: 'success' | 'info' | 'error' | undefined;

  constructor(private router: Router, private loginService: LoginService) {
    this.signUpForm = new FormGroup(
      {
        firstName: new FormControl('', [Validators.required]),
        lastName: new FormControl('', [Validators.required]),
        email: new FormControl('', [Validators.required, Validators.email]),
        password: new FormControl('', [Validators.required]),
        reenteredPass: new FormControl(''),
      },
      { validators: this.checkPasswords }
    );
  }

  checkPasswords(group: AbstractControl): ValidationErrors | null {
    const password = group.get('password')?.value;
    const reenteredPass = group.get('reenteredPass')?.value;
    return password == reenteredPass ? null : { passwordMismatch: true };
  }

  goSubmitForm() {
    if (!this.signUpForm.invalid) {
      const firstName = this.signUpForm.get('firstName')?.value;
      const lastName = this.signUpForm.get('lastName')?.value;
      const email = this.signUpForm.get('email')?.value;
      const password = this.signUpForm.get('password')?.value;

      console.log({
        firstName: firstName,
        lastName: lastName,
        email: email,
        password: password,
      });

      this.awaitingRes = true;
      this.loginService
        .createUser(firstName, lastName, email, password)
        .subscribe((res) => {
          this.awaitingRes = false;
          console.log(res);
          if (res.id === '') {
            this.showAlert = true;
            this.alertType = 'error';
            this.alertMessage = 'User with this email already exists';
          } else if (res.id) {
            this.showAlert = true;
            this.alertType = 'success';
            this.alertMessage = 'Registration successful. Redirecting...';
            setTimeout(() => {
              this.router.navigate(['/login']);
            }, 3000);
          } else {
            this.showAlert = true;
            this.alertType = 'error';
            this.alertMessage = 'Something went wrong';
          }
        });
    } else {
      console.log('Invalid entry');
    }
  }
}
