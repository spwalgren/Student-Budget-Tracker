
export interface SignUpRequest {
  firstName: string,
  lastName: string,
  email: string,
  password: string
}

export interface SignUpResponse {
  id: string
}

export interface LogInRequest {
  email: string,
  password: string
}

export interface LogInResponse {
  Message: string
}

export interface GetUserDataResponse {
  ID?: string,
  email?: string,
  firstName?: string,
  lastName?: string,
  Message?: string
}