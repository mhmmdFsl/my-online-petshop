import {Component, OnInit} from '@angular/core';
import {SignUpdRq} from "../signuprq.interface";
import {AbstractControl, FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {AuthenticationService} from "../authentication.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  form: FormGroup = new FormGroup({
    name: new FormControl(''),
    password: new FormControl(''),
    confirmPassword: new FormControl(''),
    principal: new FormControl('')
  })
  error?: string;
  submited = false;

  constructor(
      private formBuilder: FormBuilder,
      private authenticationService: AuthenticationService,
      private router: Router
  ) {}

  ngOnInit() {
    this.form = this.formBuilder.group({
      name: ['', Validators.required],
      password: ['', Validators.required],
      confirmPassword: ['', Validators.required],
      principal: ['', [Validators.required, Validators.minLength(6)]]
    });
  }

  get f(): { [key: string]: AbstractControl } {
    return this.form.controls;
  }

  signUp() {
    this.submited = true;
    const signUpRq: SignUpdRq = {
      name: this.form.value['name'],
      password: this.form.value['password'],
      confirmPassword: this.form.value['confirmPassword'],
      principal: this.form.value['principal']
    };

    this.authenticationService.signUp(signUpRq)
        .subscribe({
          next: (n) => console.log(n),
          error: (e) => this.error = e['error']['message'],
          complete: () => this.router.navigate(['/login'])
        })
  }

  onReset() {
    this.submited = false;
    this.form.reset();
  }
}
