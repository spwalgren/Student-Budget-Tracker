import { Budget, GetBudgetsResponse, Period } from "./types/budget-system";

const budgets: Budget[] = [
  {
    userId: 20,
    budgetId: 0,
    data: {
      category: "Groceries",
      amountLimit: 100,
      frequency: Period.weekly,
      duration: 1,
    }
  },
  {
    userId: 20,
    budgetId: 1,
    data: {
      category: "Groceries",
      amountLimit: 500,
      frequency: Period.monthly,
      duration: 1,
    }
  },
  {
    userId: 20,
    budgetId: 2,
    data: {
      category: "Rent",
      amountLimit: 500,
      frequency: Period.monthly,
      duration: 1,
      count: 6
    }
  },
  {
    userId: 20,
    budgetId: 2,
    data: {
      category: "Food",
      amountLimit: 80,
      frequency: Period.weekly,
      duration: 2,
    }
  },

]

async function pause<T>(ms: number): Promise<T> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export async function getBudgets(): Promise<GetBudgetsResponse> {
  let newBudgets: Budget[] = [];
  await pause<void>(1000);
  newBudgets = [...budgets];
  return { budgets: newBudgets };
}