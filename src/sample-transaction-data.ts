import { GenericResponse } from "./types/api-system";
import { CreateTransactionRequest, GetTransactionsResponse, Transaction } from "./types/transaction-system";

const transactions: Transaction[] = [
  {
    name: "Publix",
    amount: 30,
    date: "2023-02-18",
    category: "Groceries"
  },
  {
    name: "Starbucks",
    amount: 8,
    date: "2023-01-19",
    category: "Food",
    description: "Also paid for my friend's drink."
  },
  {
    name: "Bookstore",
    amount: 25,
    date: "2023-01-12",
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    name: "Bookstore",
    amount: 25,
    date: "2023-01-12",
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    name: "Bookstore",
    amount: 25,
    date: "2023-01-12",
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    name: "Bookstore",
    amount: 25,
    date: "2023-01-12",
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

export async function createTransaction(createTransactionRequest: CreateTransactionRequest): Promise<GenericResponse> {
  transactions.push(createTransactionRequest.data);
  await pause<void>(1000);
  return {};
}