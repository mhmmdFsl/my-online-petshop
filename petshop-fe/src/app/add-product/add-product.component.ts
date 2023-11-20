import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Editor } from 'ngx-editor';
import { createProductRq } from '../create-product-rq.interface';
import { ProductService } from '../product.service';
import { StorageService } from '../storage.service';


@Component({
  selector: 'app-add-product',
  templateUrl: './add-product.component.html',
  styleUrls: ['./add-product.component.css']
})
export class AddProductComponent implements OnInit {
    
  editor: Editor;
  html = '';
  form: FormGroup = new FormGroup({
    name: new FormControl(''),
    price: new FormControl(0),
    file: new FormControl(''),
    description: new FormControl('')
  });
  isSuccess: boolean = false;
  shopId = '';

  constructor(
    private productService: ProductService,
    private storageService: StorageService
  ) {
    this.editor =  new Editor();
  }

  ngOnInit(): void {
    this.shopId = this.storageService.getShopId()
    console.log(`get shopId: ${this.shopId}`)
  }


  onDestroy() {
    this.editor?.destroy();
  }

  onSubmit() {
    const createProductRq: createProductRq = {
      name: this.form.get('name')?.value,
      description: this.form.get('description')?.value,
      price: this.form.get('price')?.value,
      imageUrl: this.form.get('file')?.value,
      shopID: this.shopId
    };

    this.productService.createProduct(createProductRq)
      .subscribe((rs: any) => {
        if(!rs['errors']) {
          this.isSuccess = true
          this.form.reset()
        }
      });
  }

  getContentAsPlainText(content: string) {
    const parser = new DOMParser();
    const parsedContent = parser.parseFromString(content, 'text/html');
    const plainText = parsedContent.body.textContent;

    return plainText
  }

}
