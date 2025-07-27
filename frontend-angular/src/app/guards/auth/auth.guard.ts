import { Injectable } from '@angular/core';
import { CanActivate, Router, UrlTree } from '@angular/router';
import { AuthService } from '../../services/auth.service';

@Injectable({
  providedIn: 'root',
})
export class AuthGuard implements CanActivate {
  constructor(private auth: AuthService, private router: Router) {}

  canActivate(): boolean | UrlTree {
    console.log('AuthGuard running...');
    const token = localStorage.getItem('token');

    if (this.auth.isLoggedIn()) {
      return true; // izinkan akses
    } else {
      console.warn('Blocked by AuthGuard');
      // Redirect ke login
      return this.router.parseUrl('/login');
      // return this.router.createUrlTree(['/login']);
    }
  }
}
