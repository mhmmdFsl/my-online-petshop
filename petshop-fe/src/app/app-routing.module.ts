import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {SignUpComponent} from "./sign-up/sign-up.component";
import {LoginComponent} from "./login/login.component";
import {HomeComponent} from "./home/home.component";
import {LoginGuard} from "./login.guard";
import { ProductPageComponent } from './product-page/product-page.component';

const routes: Routes = [
  { path: 'sign-up', component: SignUpComponent },
  { path: 'login', component: LoginComponent, canActivate: [LoginGuard] },
  { path: '', component: HomeComponent },
  { path: 'product/detail/:slug', component: ProductPageComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
