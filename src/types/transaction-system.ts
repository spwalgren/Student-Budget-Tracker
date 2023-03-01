import { GenericResponse } from "./api-system"

export interface TransactionContent {
  name: string,
  amount: number,
  date: string,
  category?: string,
  description?: string
}

export interface Transaction extends TransactionContent {
  userId: number,
  transactionId: number
}

export interface CreateTransactionRequest {
  data: TransactionContent
}

export interface UpdateTransactionRequest {
  data: Transaction
}

export interface CreateTransactionResponse extends GenericResponse {
  userId: number,
  transactionId: number
}

export interface GetTransactionResponse extends GenericResponse {
  data: Transaction
}

export interface GetTransactionsResponse extends GenericResponse {
  data: Transaction[]
}