import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { GenericResponse } from 'src/types/api-system';
import { Period } from 'src/types/budget-system';
import { GetProgressResponse } from 'src/types/progress.system';

@Injectable({
  providedIn: 'root'
})
export class ProgressService {

  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) { }
  

  GetProgress(): Observable<GetProgressResponse> {
    const url = `${this.requestBase}/progress`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetProgressResponse>(url, options).pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not get progress",
            data:[]
          });
        })
      );
  }

  GetPreviousProgress(): Observable<GetProgressResponse> {
    const url = `${this.requestBase}/progress/previous`;
    
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetProgressResponse>(url, options).pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not get previous progress",
            data:[]
          });
        })
      );

  }

}

