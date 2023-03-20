import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { createBudget, deleteBudget, getBudgets, updateBudget } from 'src/sample-budget-data';
import { GenericResponse } from 'src/types/api-system';
import { CreateBudgetRequest, CreateBudgetResponse, GetBudgetsResponse, UpdateBudgetRequest } from 'src/types/budget-system';

@Injectable({
  providedIn: 'root'
})
export class BudgetService {

  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) { }

  createBudget(requestData: CreateBudgetRequest): Observable<CreateBudgetResponse> {
    return from(createBudget(requestData));
  }

  getBudgets(): Observable<GetBudgetsResponse> {
    return from(getBudgets());
  }

  updateBudget(requestData: UpdateBudgetRequest): Observable<GenericResponse> {
    return from(updateBudget(requestData));
  }

  deleteBudget(budgetId: number): Observable<GenericResponse> {
    return from(deleteBudget(budgetId));
  }
}
