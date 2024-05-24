import axios from "axios";

export class HttpClient {
  private baseUrl: string;


  constructor() {
    this.baseUrl = "https://jsonplaceholder.typicode.com";
  }
  async get(path: string) {
    try {
      const response = await axios(`${this.baseUrl}${path}`);
      return response.data;
    }catch(e) {
      console.log(e);
      return e
    }
  }
}