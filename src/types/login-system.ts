import { GenericResponse } from "./api-system"

export interface SignUpRequest {
  firstName: string,
  lastName: string,
  email: string,
  password: string
}

export interface LogInRequest {
  email: string,
  password: string
}

export interface GetUserDataResponse extends GenericResponse {
  id: string,
  email: string,
  firstName: string,
  lastName: string,
}