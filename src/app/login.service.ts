import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, tap } from 'rxjs';
import User from 'src/types/User';
import UserData from 'src/types/UserData';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  private usersUrl = 'http://localhost:8080/users';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private http: HttpClient) { }

  getAuthentication(username: string, password: string): Observable<string> {
    return of('');
  }

  getUserData(token: string): Observable<UserData> {
    return of({ customString: '' });
  }

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.usersUrl)
      .pipe(
        tap(_ => console.log("Got Users")),
        catchError((err): Observable<User[]> => {
          console.error(err);
          return of([] as User[]);
        })
      );
  }
}
