import { Component } from '@angular/core';
import {StorageService} from "./storage.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  constructor(private storageService: StorageService, private router: Router) {
  }
  title = 'petshop-fe';
}
