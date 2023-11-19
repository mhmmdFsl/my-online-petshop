import { Component, OnInit } from '@angular/core';
import { ProductService } from '../product.service';
import { StorageService } from '../storage.service';
import { Shop } from '../shop.interface';
import { FormControl, FormGroup } from '@angular/forms';
import { CreateShopRq } from '../create-shop-rq.interface';

@Component({
  selector: 'app-shop',
  templateUrl: './shop.component.html',
  styleUrls: ['./shop.component.css']
})
export class ShopComponent implements OnInit{
  userId = '';
  shop?: Shop;
  form: FormGroup = new FormGroup({
    name: new FormControl(''),
    file: new FormControl(''),
  });

  constructor(
    private productService: ProductService,
    private storageService: StorageService
  ) {}
  
  ngOnInit(): void {
    this.userId = this.storageService.getUserId();
    this.productService.getShop(this.userId)
      .valueChanges.subscribe((rs: any) => {
        const shops = rs['data']['shops'] as [Shop]
        this.shop = shops[0]
      });
  }

  onSubmit() {
    const createShopRq: CreateShopRq = {
      name: this.form.value['name'],
      userId: this.userId,
      logoUrl: this.form.value['file']
    };

    this.productService.createShop(createShopRq)
    .subscribe((rs: any) => {
      this.shop = rs['data']['createShop'] as Shop
    })
  }

}
