import { User } from "../domain/user";
import { HttpClient } from "./httpClient";

export class UserApi {
  private httpClient: HttpClient;

  constructor() {
    this.httpClient = new HttpClient();
  }

  async getAll(): Promise<User[]> {
    return await this.httpClient.get("/users");
  }
}
