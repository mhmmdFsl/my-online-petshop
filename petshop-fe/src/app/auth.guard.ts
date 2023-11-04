import {StorageService} from "./storage.service";
import {CanActivate, Router} from "@angular/router";
import {Injectable} from "@angular/core";

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private storageService: StorageService, private router: Router) {}

  canActivate(): boolean {
    if (this.storageService.isLogin()) {
      // User is already logged in, allow access
      return true;
    }

    // User is not logged in, redirect to the login page
    this.router.navigate(['/login']);
    return false;
  }
}