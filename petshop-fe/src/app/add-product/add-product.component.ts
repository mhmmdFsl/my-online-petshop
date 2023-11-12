import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { Editor } from 'ngx-editor';


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
  })

  constructor() {
    this.editor =  new Editor();
  }

  ngOnInit(): void {
  }


  onDestroy() {
    this.editor?.destroy()
  }

  onSubmit() {
    console.log(this.getContentAsPlainText(this.form.get('description')?.value))
  }

  getContentAsPlainText(content: string) {
    const parser = new DOMParser();
    const parsedContent = parser.parseFromString(content, 'text/html');
    const plainText = parsedContent.body.textContent;

    return plainText
  }

}
