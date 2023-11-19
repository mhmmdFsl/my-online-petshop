import { Injectable } from '@angular/core';
import {Apollo, QueryRef} from "apollo-angular";
import {gql} from "@apollo/client/core";
import { CreateShopRq } from './create-shop-rq.interface';

const GET_PRODDUCT_QUERY = gql`
  query Query($input: ProductFilter) {
      products(input: $input) {
        id
        name
        imageUrl
        slug
        price
        description
        createdAt
        updatedAt
      }
    }
  `

  const GET_SHOP_QUERY = gql`
  query Query($input: ShopFilter) {
      shops(input: $input) {
        id
        userId
        name
        logoUrl
        isVerified
        status
        createdAt
        updatedAt
      }
    }
  `

  const CREATE_SHOP_QUERY = gql`
    mutation Mutation($input: NewShop) {
      createShop(input: $input) {
        id
        name
        logoUrl
        isVerified
        status
        userId
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

  getShop(userId: string) {
    return this.apollo
    .watchQuery({
      query: GET_SHOP_QUERY,
      variables: {
        input: {
          limit: 10,
          userId: userId
        }
      }
    })
  }

  createShop(createShopRq: CreateShopRq) {
    console.log(createShopRq)
    return this.apollo
      .mutate({
        mutation: CREATE_SHOP_QUERY,
        variables: {
          input: {
            name: createShopRq.name,
            logoUrl: createShopRq.logoUrl,
            userId: createShopRq.userId
          }
        }
      })
  }
}
