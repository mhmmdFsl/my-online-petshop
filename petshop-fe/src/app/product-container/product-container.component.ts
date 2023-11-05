import {Component, OnInit} from '@angular/core';
import {Apollo, QueryRef} from "apollo-angular";
import {Product} from "../product.interface";
import {gql} from "@apollo/client/core";
import {Subscription} from "rxjs";

@Component({
  selector: 'app-product-container',
  templateUrl: './product-container.component.html',
  styleUrls: ['./product-container.component.css']
})
export class ProductContainerComponent implements OnInit {
  constructor(private apollo: Apollo) {}

  products?: [Product]
  productQuery?: QueryRef<any>;
  privateProductQuerySub?: Subscription;

  ngOnInit() {
    this.productQuery = this.apollo
      .watchQuery({
          query: gql`
          {
            products {
              id
              name
              imageUrl
              price
            }
          }
        `,
          pollInterval: 500
        }
      );
    
    this.privateProductQuerySub = this.productQuery.valueChanges.subscribe(
      (rs: any) => {
        this.products = rs['data']['products'] as [Product];
      }
    );
  }
  
  refresh(): void {
    this.productQuery?.refetch()
  }
}
