import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, tap } from 'rxjs';
import User from 'src/types/User';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) {}

  createAuth(email: string, password: string): Observable<string> {
    const url = `${this.requestBase}/login`;
    const body = { email: email, password: password };
    const options = this.httpOptions;

    return this.http.post<any>(url, body, options).pipe(
      tap<any>((res) => {
        console.log(res);
      }),
      catchError((err) => {
        console.log(err);
        return of('');
      })
    );
  }

  createUser(
    firstName: string,
    lastName: string,
    email: string,
    password: string
  ): Observable<any> {
    const url = `${this.requestBase}/signup`;
    const body = {
      firstName: firstName,
      lastName: lastName,
      email: email,
      password: password,
    };
    const options = this.httpOptions;

    return this.http.post<any>(url, body, options).pipe(
      catchError((err) => {
        console.log(err);
        return of({});
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
