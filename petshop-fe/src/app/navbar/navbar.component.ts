import {Component, OnInit} from '@angular/core';
import {NavigationEnd, Router} from "@angular/router";
import {StorageService} from "../storage.service";
import { ProductService } from '../product.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {
  isLogin: boolean = false;
  isDropdown: boolean = false;
  searchValue: string = ''

  constructor(
    private router: Router, 
    private storageService: StorageService,
    private productService: ProductService
  ) {
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
    this.router.navigateByUrl('')
  }

  dropDown() {
    this.isDropdown = !this.isDropdown
  }

  search() {
    this.productService.setName(this.searchValue)
  }
}
