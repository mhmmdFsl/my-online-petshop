import {Component, OnInit} from '@angular/core';
import {ActivatedRoute} from "@angular/router";
import {Product} from "../product.interface";
import {ProductService} from "../product.service";

@Component({
  selector: 'app-product-page',
  templateUrl: './product-page.component.html',
  styleUrls: ['./product-page.component.css']
})
export class ProductPageComponent implements OnInit {

  slug?: string;
  product: Product = {
    name: '',
    imageUrl: '',
    price: 0,
    slug: ''
  }
  
  constructor(
      private activatedRouter: ActivatedRoute,
      private productService: ProductService
    ) {
  }
  
  ngOnInit(): void {
    this.activatedRouter.params.subscribe(p => {
      this.slug = p['slug'];
    })
    
    this.productService.getProductBySlug(this.slug!!)
      .valueChanges.subscribe( (rs: any) => {
          const data = rs['data']['products'] as [Product]
          this.product = data[0]
        }
      )
  }

}
