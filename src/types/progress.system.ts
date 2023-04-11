import { GenericResponse} from "./api-system";
import { Period } from "./budget-system";
//return the total spent (given category, week, freq)
//return total spent for each category
//pie chart-> total spent / limit given freq


export interface ProgressByPeriod {
    userId: number,
    totalSpent: number,
    transactionIdList: number[],
    budgetIdList: number[],
    category: string,
    budgetGoal: number,
    frequency: Period
}

export interface GetProgressRequest {
    frequency: Period
}

export interface GetProgressResponse extends GenericResponse {
    data: ProgressByPeriod
}
