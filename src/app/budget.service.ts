import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { getBudgets } from 'src/sample-budget-data';
import { GetBudgetsResponse } from 'src/types/budget-system';

@Injectable({
  providedIn: 'root'
})
export class BudgetService {

  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) { }

  getBudgets(): Observable<GetBudgetsResponse> {
    return from(getBudgets());
  }
}
