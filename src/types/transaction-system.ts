import { GenericResponse } from "./api-system"

export interface Transaction {
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
  index: number,
  data: Transaction
}

export interface GetTransactionsResponse extends GenericResponse {
  data: Transaction[]
}
