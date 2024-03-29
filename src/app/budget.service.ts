import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { createBudget, deleteBudget, getBudgets, updateBudget } from 'src/sample-budget-data';
import { GenericResponse } from 'src/types/api-system';
import { CreateBudgetRequest, CreateBudgetResponse, GetBudgetCategoriesResponse, GetBudgetsResponse, GetCyclePeriodResponse, UpdateBudgetRequest } from 'src/types/budget-system';

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
    const url = `${this.requestBase}/budget`;
    const body = { ...requestData };
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.post<CreateBudgetResponse>(url, body, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not create budget",
            userId: 0,
            budgetId: 0
          })
        })
      )
  }

  getBudgets(): Observable<GetBudgetsResponse> {
    const url = `${this.requestBase}/budget`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetBudgetsResponse>(url, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not get budgets",
            budgets: []
          })
        })
      )
  }

  getBudgetCategories(): Observable<GetBudgetCategoriesResponse> {
    const url = `${this.requestBase}/budget/categories`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetBudgetCategoriesResponse>(url, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not get budget categories",
            categories: []
          });
        })
      );
  }

  getCyclePeriod(): Observable<GetCyclePeriodResponse> {
    let today = new Date();
    today.setMinutes(today.getMinutes() - today.getTimezoneOffset());
    const todayString = today.toISOString().split('T')[0];
    const url = `${this.requestBase}/budget/cycle/${todayString}`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetCyclePeriodResponse>(url, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not get cycle period",
            data: []
          });
        })
      );
  }

  updateBudget(requestData: UpdateBudgetRequest): Observable<GenericResponse> {
    const url = `${this.requestBase}/budget`;
    const body = { ...requestData };
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.put<GenericResponse>(url, body, options)
      .pipe(
        map((_) => ({})),
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not update budget",
          })
        })
      )
  }

  deleteBudget(budgetId: number): Observable<GenericResponse> {
    const url = `${this.requestBase}/budget/${budgetId}`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.delete<GenericResponse>(url, options)
      .pipe(
        map((_) => ({})),
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not delete budget",
          })
        })
      )
  }
}
