import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of } from 'rxjs';
import { GetEventsResponse } from 'src/types/calendar-system';

@Injectable({
  providedIn: 'root'
})
export class CalendarService {

  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) { }

  getBudgets(month: number): Observable<GetEventsResponse> {
    const url = `${this.requestBase}/calendar/${month}`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetEventsResponse>(url, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not get events",
            events: []
          })
        })
      );
  }
}
