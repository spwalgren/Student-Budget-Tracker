import { GenericResponse } from "./api-system"

export interface Transaction {
  userId?: number,
  transactionId?: number
  name: string,
  amount: number,
  date: string,
  category?: string,
  description?: string
}

export interface CreateTransactionRequest {
  data: Transaction
}

export interface EditTransactionRequest {
  data: Transaction
}

export interface GetTransactionsResponse extends GenericResponse {
  data: Transaction[]
}
