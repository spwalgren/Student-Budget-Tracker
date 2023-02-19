import { GetTransactionsResponse, Transaction } from "./types/transaction-system";

const transactions: Transaction[] = [
  {
    userId: 1,
    name: "Publix",
    amount: 30,
    date: "2023-02-18",
    category: "Groceries"
  },
  {
    userId: 2,
    name: "Starbucks",
    amount: 8,
    date: "2023-01-19",
    category: "Food",
    description: "Also paid for my friend's drink."
  },
  {
    userId: 3,
    name: "Bookstore",
    amount: 25,
    date: "2023-01-12",
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    userId: 3,
    name: "Bookstore",
    amount: 25,
    date: "2023-01-12",
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    userId: 3,
    name: "Bookstore",
    amount: 25,
    date: "2023-01-12",
    category: "Supplies",
    description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Non aspernatur voluptatum fugit aut aliquam nam autem impedit facere voluptatem sit!"
  },
  {
    userId: 3,
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