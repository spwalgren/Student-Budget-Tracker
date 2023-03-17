import { GenericResponse } from "./api-system";

export enum Period {
  weekly = "weekly",
  monthly = "monthly",
  yearly = "yearly"
}

export interface BudgetContent {
  category: string,
  amountLimit: number,
  frequency: Period,
  duration: number,
  count: number,
}

export interface Budget {
  userId: number,
  budgetId: number,
  data: BudgetContent
}

export interface UpdateBudgetRequest {
  newBudget: Budget
}

export interface CreateBudgetResponse extends GenericResponse {
  userId: number,
  budgetId: number
}

export interface GetBudgetsResponse extends GenericResponse {
  budgets: Budget[]
}