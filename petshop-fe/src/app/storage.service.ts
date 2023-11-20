import { Injectable } from '@angular/core';
import { Subject } from 'rxjs';

const REFRESH = 'refresh_token';
const ACCESS = 'access_token';
const SHOP_ID = 'shop_id';

@Injectable({
  providedIn: 'root'
})
export class StorageService {

  private shopId: string = '';

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

  setShopId(id: string) {
   window.sessionStorage.setItem(SHOP_ID, id);
  }

  getShopId(): string {
    return window.sessionStorage.getItem(SHOP_ID) as string;
  }
}
