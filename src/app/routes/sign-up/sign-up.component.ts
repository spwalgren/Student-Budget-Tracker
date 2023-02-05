import { Component } from '@angular/core';
import { FormGroup, FormControl, Validators, AbstractControl, FormGroupDirective, NgForm, ValidatorFn, ValidationErrors } from '@angular/forms';
import { ErrorStateMatcher } from '@angular/material/core';

export class PasswordChecker implements ErrorStateMatcher {
  isErrorState(control: AbstractControl<any, any> | null, form: FormGroupDirective | NgForm | null): boolean {
    return control?.parent?.errors?.['passwordMismatch'];
  }
}

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent {

  signUpForm: FormGroup;
  passwordChecker = new PasswordChecker();

  constructor() {
    this.signUpForm = new FormGroup({
      firstName: new FormControl('', [Validators.required]),
      lastName: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required]),
      reenteredPass: new FormControl(''),
    }, { validators: this.checkPasswords });
  }

  checkPasswords(group: AbstractControl): ValidationErrors | null {
    const password = group.get('password')?.value;
    const reenteredPass = group.get('reenteredPass')?.value;
    return password == reenteredPass ? null : { passwordMismatch: true };
  }

  goSubmitForm() {
    if (!this.signUpForm.invalid) {
      console.log({
        firstName: this.signUpForm.get('firstName')?.value,
        lastName: this.signUpForm.get('lastName')?.value,
        email: this.signUpForm.get('email')?.value,
        password: this.signUpForm.get('password')?.value,
        reenteredPass: this.signUpForm.get('reenteredPass')?.value,
      });
    } else {
      console.log('Invalid entry');
    }
  }
}