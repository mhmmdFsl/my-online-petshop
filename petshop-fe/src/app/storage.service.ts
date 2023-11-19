import { Injectable } from '@angular/core';

const REFRESH = 'refresh_token'
const ACCESS = 'access_token'

@Injectable({
  providedIn: 'root'
})
export class StorageService {

  constructor() { }

  clean(): void {
    window.sessionStorage.clear();
  }

  saveRefreshToken(value: any): void {
    window.sessionStorage.removeItem(REFRESH);
    window.sessionStorage.setItem(REFRESH, value);
  }

  saveAccessToken(value: any): void {
    window.sessionStorage.removeItem(ACCESS);
    window.sessionStorage.setItem(ACCESS, value);
  }

  saveUserId(value: any): void {
    window.sessionStorage.setItem('user_id', value);
  }

  getUserId(): any {
    return window.sessionStorage.getItem('user_id');
  }

  getAccessToken(): any {
    return window.sessionStorage.getItem(ACCESS);
  }

  getRefreshToken(): any {
    return window.sessionStorage.getItem(REFRESH);
  }

  isLogin(): boolean {
    const r = this.getRefreshToken();
    return !!r;

  }
}
