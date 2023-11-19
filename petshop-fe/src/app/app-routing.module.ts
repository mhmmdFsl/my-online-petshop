import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {SignUpComponent} from "./sign-up/sign-up.component";
import {LoginComponent} from "./login/login.component";
import {HomeComponent} from "./home/home.component";
import {LoginGuard} from "./login.guard";
import { ProductPageComponent } from './product-page/product-page.component';
import { ShopDashboardComponent } from './shop-dashboard/shop-dashboard.component';
import { AddProductComponent } from './add-product/add-product.component';
import { ShopProductComponent } from './shop-product/shop-product.component';
import { ShopComponent } from './shop/shop.component';
import { AuthGuard } from './auth.guard';

const routes: Routes = [
  { 
    path: 'sign-up', 
    component: SignUpComponent 
  },
  { 
    path: 'login', 
    component: LoginComponent, 
    canActivate: [LoginGuard] 
  },
  { 
    path: '', 
    component: HomeComponent 
  },
  { 
    path: 'product/detail/:slug', 
    component: ProductPageComponent
  },
  { 
    path: 'shop', 
    component: ShopDashboardComponent, 
    children: [
      { path: '', component: ShopComponent},
      { path: 'product', component: ShopProductComponent},
      { path: 'add-product', component: AddProductComponent },
    ],
    // canActivate: [AuthGuard]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
