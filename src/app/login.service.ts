import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, tap } from 'rxjs';
import User from 'src/types/User';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  private usersUrl = 'api/users';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(private http: HttpClient) { }

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
