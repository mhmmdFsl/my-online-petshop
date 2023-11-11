import { Injectable } from '@angular/core';
import {Apollo, QueryRef} from "apollo-angular";
import {gql} from "@apollo/client/core";
import { Observable, Subject, Subscription, of } from 'rxjs';
import { Product } from './product.interface';

const GET_PRODDUCT_QUERY = gql`
  query Query($input: ProductFilter) {
      products(input: $input) {
        id
        name
        imageUrl
        slug
        price
        createdAt
        updatedAt
      }
    }
  `

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  constructor(
    private apollo: Apollo
  ) { }

  queryRef?: QueryRef<any, any>;
  

  setName(s: string) {
    console.log(s)
  }

  getAllProduct() {
    this.queryRef =this.apollo
      .watchQuery({
          query: GET_PRODDUCT_QUERY,
          variables: {
            input: {
              limit: 10,
            }
          },
        pollInterval: 1000
        }
      )
    return this.queryRef
  }
  
  getProductBySlug(slug: string) {
    return this.apollo
    .watchQuery({
      query: GET_PRODDUCT_QUERY,
      variables: {
        input: {
          limit: 1,
          slug: slug
        }
      }
    })
  }
}
