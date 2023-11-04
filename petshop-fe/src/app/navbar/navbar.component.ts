import {Component, OnInit} from '@angular/core';
import {NavigationEnd, Router} from "@angular/router";
import {StorageService} from "../storage.service";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  isLogin: boolean = false;

  constructor(private router: Router, private storageService: StorageService) {
    // Listen for route changes
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        // Check if the current route is the login page
        this.isLogin = this.storageService.isLogin();
      }
    });
  }

  ngOnInit() {
  }

  onLogout() {
    this.storageService.clean();
    this.router.navigate(['login'])
  }
}
