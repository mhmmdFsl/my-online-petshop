import { Injectable } from '@angular/core';
import {Observable, of} from "rxjs";
import {SignUpdRq} from "./signuprq.interface";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {LoginRq} from "./login-rq.interface";

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  constructor(private http: HttpClient) { }
  
  signUp(signUpRq: SignUpdRq): Observable<any> {
      const url: string = 'http://localhost:3000/api/v1/auth/sign-up';
      const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
      return this.http.post(url, signUpRq, {headers})
  }

  login(loginRq: LoginRq): Observable<any> {
    const url: string = 'http://localhost:3000/api/v1/auth/login'
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    return this.http.post(url, loginRq, {headers})
  }
}
