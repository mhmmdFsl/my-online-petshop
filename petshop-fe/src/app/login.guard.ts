import {CanActivate, CanActivateFn, Router} from '@angular/router';
import {Injectable} from "@angular/core";
import {StorageService} from "./storage.service";

@Injectable({
  providedIn: 'root'
})
export class LoginGuard implements CanActivate {
  constructor(private storageService: StorageService, private router: Router) {}

  canActivate(): boolean {
    if (!this.storageService.isLogin()) {
      // User is already logged in, allow access
      return true;
    }

    // User is not logged in, redirect to the login page
    this.router.navigate(['/home']);
    return false;
  }
}
