import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, tap } from 'rxjs';
import {
  GetUserDataResponse,
  LogInRequest,
  LogInResponse,
  SignUpRequest,
  SignUpResponse,
} from 'src/types/login-system';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) {}

  logIn(logInRequest: LogInRequest): Observable<LogInResponse> {
    const url = `${this.requestBase}/login`;
    const body = { ...logInRequest };
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.post<LogInResponse>(url, body, options).pipe(
      catchError((err) => {
        console.log(err);
        return of({ Message: 'Error' });
      })
    );
  }

  signUp(signUpRequest: SignUpRequest): Observable<SignUpResponse> {
    const url = `${this.requestBase}/signup`;
    const body = { ...signUpRequest };
    const options = this.httpOptions;

    return this.http.post<SignUpResponse>(url, body, options).pipe(
      catchError((err) => {
        console.log(err);
        return of({ id: '' });
      })
    );
  }

  getUserData(): Observable<GetUserDataResponse> {
    const url = `${this.requestBase}/user`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetUserDataResponse>(url, options).pipe(
      catchError((err) => {
        console.log(err);
        return of({ Message: 'Error' });
      })
    );
  }

  // getUserData(token: string): Observable<UserData> {
  //   return of({ customString: '' });
  // }

  // getUsers(): Observable<User[]> {
  //   return this.http.get<User[]>(this.requestBase)
  //     .pipe(
  //       tap(_ => console.log("Got Users")),
  //       catchError((err): Observable<User[]> => {
  //         console.error(err);
  //         return of([] as User[]);
  //       })
  //     );
  // }
}
