import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { createTransaction, getTransactions, updateTransaction, deleteTransaction } from 'src/sample-transaction-data';
import { GenericResponse } from 'src/types/api-system';
import { CreateTransactionRequest, UpdateTransactionRequest, GetTransactionsResponse, CreateTransactionResponse, Transaction } from 'src/types/transaction-system';

@Injectable({
  providedIn: 'root'
})
export class TransactionService {
  private requestBase = 'http://localhost:8080/api';
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  };

  constructor(private http: HttpClient) { }

  createTransaction(transactionRequest: CreateTransactionRequest): Observable<CreateTransactionResponse> {
    const url = `${this.requestBase}/transaction`
    const body = { ...transactionRequest };
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.post<CreateTransactionResponse>(url, body, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not create transaction",
            userId: 0,
            transactionId: 0
          })
        })
      );
  }

  getTransactions(): Observable<GetTransactionsResponse> {
    const url = `${this.requestBase}/transaction`
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.get<GetTransactionsResponse>(url, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not retrieve transactions",
            data: []
          })
        })
      );
  }

  updateTransaction(transactionRequest: UpdateTransactionRequest): Observable<GenericResponse> {
    const url = `${this.requestBase}/transaction`
    const body = { ...transactionRequest };
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };

    return this.http.put<GenericResponse>(url, body, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not update transaction"
          })
        })
      )
  }

  deleteTransaction(transaction: Transaction): Observable<GenericResponse> {
    const url = `${this.requestBase}/transaction/${transaction.userId}/${transaction.transactionId}`;
    const options = {
      headers: this.httpOptions.headers,
      withCredentials: true,
    };
    return this.http.delete<GenericResponse>(url, options)
      .pipe(
        catchError((err) => {
          console.log(err);
          return of({
            err: "Could not delete transaction"
          })
        })
      )
  }
}
