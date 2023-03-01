import { GenericResponse } from "./types/api-system";
import { CreateTransactionRequest, UpdateTransactionRequest, GetTransactionsResponse, Transaction, CreateTransactionResponse } from "./types/transaction-system";

const transactions: Transaction[] = [
  {
    userId: 20,
    transactionId: 0,
    name: "Publix",
    amount: 30,
    date: new Date("2023-02-18").toISOString(),
    category: "Groceries"
  },
  {
    userId: 20,
    transactionId: 1,
    name: "Starbucks",
    amount: 8,
    date: new Date("2022-1-19").toISOString(),
    category: "Food",
    description: "Also paid for my friend's drink."
  },
  {
    userId: 20,
    transactionId: 2,
    name: "Bookstore",
    amount: 25,
    date: new Date("2023-01-15").toISOString(),
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    userId: 20,
    transactionId: 3,
    name: "Bookstore",
    amount: 25,
    date: new Date("2023-01-14").toISOString(),
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    userId: 20,
    transactionId: 4,
    name: "Bookstore",
    amount: 25,
    date: new Date("2023-01-13").toISOString(),
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    userId: 20,
    transactionId: 5,
    name: "Bookstore",
    amount: 25,
    date: new Date("2023-01-12").toISOString(),
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
];

async function pause<T>(ms: number): Promise<T> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export async function getTransactions(): Promise<GetTransactionsResponse> {
  let newTransactions: Transaction[] = [];
  await pause<void>(1000);
  newTransactions = [...transactions];
  return { data: newTransactions };
}

export async function createTransaction(createTransactionRequest: CreateTransactionRequest): Promise<CreateTransactionResponse> {

  let newId = transactions.length;
  transactions.push({
    userId: 20,
    transactionId: newId,
    ...createTransactionRequest.data
  });
  await pause<void>(1000);
  return {
    userId: 20,
    transactionId: newId
  };
}

export async function updateTransaction(updateTransactionRequest: UpdateTransactionRequest): Promise<GenericResponse> {
  const targetIndex = transactions.findIndex((elem) => elem.transactionId == updateTransactionRequest.data.transactionId);
  transactions[targetIndex] = updateTransactionRequest.data;
  console.log(transactions);
  await pause<void>(1000);
  return {};
}

export async function deleteTransaction(transactionId: number): Promise<GenericResponse> {
  const targetIndex = transactions.findIndex((elem) => elem.transactionId == transactionId);
  transactions.splice(targetIndex, 1);
  await pause<void>(1000);
  return {};
}