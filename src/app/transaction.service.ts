import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, of, map, from } from 'rxjs';
import { createTransaction, getTransactions, updateTransaction, deleteTransaction } from 'src/sample-transaction-data';
import { GenericResponse } from 'src/types/api-system';
import { CreateTransactionRequest, UpdateTransactionRequest, GetTransactionsResponse, CreateTransactionResponse } from 'src/types/transaction-system';

@Injectable({
  providedIn: 'root'
})
export class TransactionService {

  constructor() { }

  createTransaction(transactionRequest: CreateTransactionRequest): Observable<CreateTransactionResponse> {
    return from(createTransaction(transactionRequest));
  }

  getTransactions(): Observable<GetTransactionsResponse> {
    return from(getTransactions());
  }

  updateTransaction(transactionRequest: UpdateTransactionRequest): Observable<GenericResponse> {
    return from(updateTransaction(transactionRequest));
  }

  deleteTransaction(transactionId: number): Observable<GenericResponse> {
    return from(deleteTransaction(transactionId));
  }
}
