import { Injectable } from '@angular/core';
import { InMemoryDbService } from 'angular-in-memory-web-api';
import User from 'src/types/User';

@Injectable({
  providedIn: 'root'
})
export class InMemoryDataService implements InMemoryDbService {

  createDb() {
    const users: User[] = [
      {
        id: "001",
        firstName: "John",
        lastName: "Doe",
        email: "johndoe@example.com",
        password: "12345"
      },
      {
        id: "002",
        firstName: "Jane",
        lastName: "Doe",
        email: "janedoe@example.com",
        password: "54321"
      },
      {
        id: "003",
        firstName: "Jack",
        lastName: "Doe",
        email: "jackdoe@example.com",
        password: "11111"
      }
    ];
    return { users };
  }
}
