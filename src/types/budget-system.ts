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
  duration: number, // If budget repeats every 2 weeks, then duration is 2 and frequency is weekly
  count?: number,
  startDate: string
}

export interface Budget {
  userId: number,
  budgetId: number,
  data: BudgetContent
}

export interface CreateBudgetRequest extends BudgetContent { }

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

export interface GetBudgetCategoriesResponse extends GenericResponse {
  categories: string[]
}