import {Component, OnInit} from '@angular/core';
import {AbstractControl, FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {AuthenticationService} from "../authentication.service";
import {Router} from "@angular/router";
import {LoginRq} from "../login-rq.interface";
import {StorageService} from "../storage.service";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{

  form: FormGroup = new FormGroup({
    principal: new FormControl(''),
    password: new FormControl(''),
  })
  error?: string;
  submited = false;

  constructor(
    private formBuilder: FormBuilder,
    private authenticationService: AuthenticationService,
    private router: Router,
    private storageService: StorageService
  ) {}

  ngOnInit() {
    this.form = this.formBuilder.group({
      password: ['', Validators.required],
      principal: ['', [Validators.required]]
    });
  }

  get f(): { [key: string]: AbstractControl } {
    return this.form.controls;
  }

  onSubmit(): void {
    this.submited = true;
    const loginRq: LoginRq = {
      principal: this.form.value['principal'],
      password: this.form.value['password']
    }
    this.authenticationService.login(loginRq)
      .subscribe({
        next: (n) => {
          console.log(n)
          this.storageService.saveAccessToken(n['data']['accessToken'])
          this.storageService.saveRefreshToken(n['data']['refreshToken'])
        },
        error: (e) => this.error = e['error']['message'],
        complete: () => this.router.navigate([''])
      })
  }

  onReset(): void {
    this.submited = false;
    this.form.reset();
  }
}
