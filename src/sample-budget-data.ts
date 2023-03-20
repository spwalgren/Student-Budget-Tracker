import { GenericResponse } from "./types/api-system";
import { Budget, CreateBudgetRequest, CreateBudgetResponse, GetBudgetsResponse, Period, UpdateBudgetRequest } from "./types/budget-system";

const budgets: Budget[] = [
  {
    userId: 20,
    budgetId: 0,
    data: {
      category: "Groceries",
      amountLimit: 100,
      frequency: Period.weekly,
      duration: 1,
      startDate: new Date("2023-02-18").toISOString(),
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
      startDate: new Date("2023-02-18").toISOString(),
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
      count: 6,
      startDate: new Date("2023-02-18").toISOString(),
    }
  },
  {
    userId: 20,
    budgetId: 3,
    data: {
      category: "Food",
      amountLimit: 80,
      frequency: Period.weekly,
      duration: 2,
      startDate: new Date("2023-02-18").toISOString(),
    }
  },
  {
    userId: 20,
    budgetId: 4,
    data: {
      category: "Books",
      amountLimit: 400,
      frequency: Period.weekly,
      duration: 1,
      count: 1,
      startDate: new Date("2023-01-06").toISOString(),
    }
  }
]

async function pause<T>(ms: number): Promise<T> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export async function createBudget(requestData: CreateBudgetRequest): Promise<CreateBudgetResponse> {
  await pause<void>(1000);
  budgets.push(
    {
      userId: 20,
      budgetId: budgets.length,
      data: requestData
    }
  );
  return ({ userId: 20, budgetId: budgets.length - 1 });
}

export async function getBudgets(): Promise<GetBudgetsResponse> {
  let newBudgets: Budget[] = [];
  await pause<void>(1000);
  newBudgets = [...budgets];
  return { budgets: newBudgets };
}

export async function updateBudget(requestData: UpdateBudgetRequest): Promise<GenericResponse> {
  await pause<void>(1000);
  const targetIndex = budgets.findIndex((elem) => elem.budgetId === requestData.newBudget.budgetId);
  budgets[targetIndex] = { ...requestData.newBudget };
  return ({});
}

export async function deleteBudget(budgetId: number): Promise<GenericResponse> {
  await pause<void>(1000);
  const targetIndex = budgets.findIndex((elem) => elem.budgetId === budgetId);
  budgets.splice(targetIndex, 1);
  return ({});
}