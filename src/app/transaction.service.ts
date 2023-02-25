import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { createTransaction, getTransactions } from 'src/sample-transaction-data';
import { GenericResponse } from 'src/types/api-system';
import { CreateTransactionRequest, GetTransactionsResponse } from 'src/types/transaction-system';

@Injectable({
  providedIn: 'root'
})
export class TransactionService {

  constructor() { }

  createTransaction(transactionRequest: CreateTransactionRequest): Observable<GenericResponse> {
    return from(createTransaction(transactionRequest));
  }

  getTransactions(): Observable<GetTransactionsResponse> {
    return from(getTransactions());
  }
}
