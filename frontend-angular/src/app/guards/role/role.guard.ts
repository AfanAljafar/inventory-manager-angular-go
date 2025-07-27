// role.guard.ts
import { Injectable } from '@angular/core';
import {
  CanActivate,
  Router,
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
} from '@angular/router';

@Injectable({ providedIn: 'root' })
export class RoleGuard implements CanActivate {
  constructor(private router: Router) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): boolean {
    const expectedRole = route.data['role']; // ambil role dari routing
    const currentRole = localStorage.getItem('role');

    if (currentRole === expectedRole) {
      return true;
    }

    this.router.navigate(['/dashboard']); // redirect ke dashboard jika role tidak sesuai
    return false;
  }
}
