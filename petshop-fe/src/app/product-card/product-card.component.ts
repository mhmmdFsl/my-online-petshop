import {Component, Input} from '@angular/core';
import {Product} from "../product.interface";

@Component({
  selector: 'app-product-card',
  templateUrl: './product-card.component.html',
  styleUrls: ['./product-card.component.css']
})
export class ProductCardComponent {
  
  @Input() product: Product = {name:'', price: 0, imageUrl:'', slug: ''}
  name: string = 'Sample Product';
  price: Number = 10000;
  imageUrl: string = 'https://singpet.id/media/catalog/product/cache/5c4104efd38f87733e130c561cbe6e7b/R/o/Royal-Canin-Maine-Coon-Adult-Cat-4kg-C103-02-000104.jpg'
  
  order(): void {
    console.log(this.product)
  }
}
