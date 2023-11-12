import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShopDashboardComponent } from './shop-dashboard.component';

describe('ShopDashboardComponent', () => {
  let component: ShopDashboardComponent;
  let fixture: ComponentFixture<ShopDashboardComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ShopDashboardComponent]
    });
    fixture = TestBed.createComponent(ShopDashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
