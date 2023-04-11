import { GenericResponse } from "./api-system"
import { Period } from "./budget-system"

export interface EventContent {
  frequency: Period,
  startDate: string,
  endDate: string,
  category: string,
  totalSpent: number,
  amountLimit: number,
}

export interface Event {
  userId: number,
  eventId: number,
  data: EventContent
}

export interface GetEventsResponse extends GenericResponse {
  events: Event[]
}