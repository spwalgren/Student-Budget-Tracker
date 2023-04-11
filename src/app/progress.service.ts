import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { GenericResponse } from 'src/types/api-system';
import { Period } from 'src/types/budget-system';
import { GetProgressRequest, GetProgressResponse } from 'src/types/progress.system';

@Injectable({
  providedIn: 'root'
})
export class ProgressService {

  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) { }
  

  GetProgress(requestData: GetProgressRequest): Observable<GetProgressResponse> {
    const url = `${this.requestBase}/progress`;
    const body = { ...requestData };
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.post<GetProgressResponse>(url, body, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not get progress",
            data: {
              userId: 0,
              totalSpent: 0,
              transactionIdList: [0],
              budgetIdList: [0],
              category: "general",
              budgetGoal: 0,
              frequency: Period.weekly
            }
          })
        })
      )
  }

}

