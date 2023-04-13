import { GenericResponse} from "./api-system";
import { Period } from "./budget-system";
//return the total spent (given category, week, freq)
//return total spent for each category
//pie chart-> total spent / limit given freq


export interface Progress {
    userId: number,
    totalSpent: number,
    transactionIdList: number[],
    budgetId: number,
    category: string,
    budgetGoal: number,
    frequency: Period
}

export interface GetProgressResponse extends GenericResponse {
    data: Progress[]
}
