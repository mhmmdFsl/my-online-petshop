import {Component, OnInit} from '@angular/core';
import {Apollo, QueryRef} from "apollo-angular";
import {Product} from "../product.interface";
import {gql} from "@apollo/client/core";
import {Observable, Subscription, of} from "rxjs";
import { ProductService } from '../product.service';

@Component({
  selector: 'app-product-container',
  templateUrl: './product-container.component.html',
  styleUrls: ['./product-container.component.css']
})
export class ProductContainerComponent implements OnInit {
  constructor(
    private apollo: Apollo,
    private productService: ProductService
  ) {}

  products?: [Product]
  privateProductQuerySub?: Subscription;
  queryRef?: QueryRef<any, any>;


  ngOnInit() {
    this.queryRef = this.productService.getAllProduct()
    this.queryRef.valueChanges.subscribe(
      (rs: any) => {
        this.products = rs['data']['products'] as [Product];
      }
    );
  }

  
}
