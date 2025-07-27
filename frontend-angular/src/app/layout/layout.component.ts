import { Component, OnInit } from '@angular/core';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { AuthService } from '../services/auth.service';
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-layout',
  standalone: true,
  imports: [RouterOutlet, RouterLink, NgIf],
  templateUrl: './layout.component.html',
  styleUrl: './layout.component.css',
})
export class LayoutComponent implements OnInit {
  constructor(private auth: AuthService, private router: Router) {}

  role = '';

  ngOnInit() {
    this.role = localStorage.getItem('role') || '';
  }
  // get role(): string {
  //   const role = localStorage.getItem('role') || '';
  //   // Optional: log setiap kali getter dipanggil
  //   console.log('Sidebar getter role:', role);
  //   return role;
  // }

  isAdmin(): boolean {
    return this.role === 'admin';
  }

  logout() {
    // localStorage.removeItem('token'); // Hapus token
    this.auth.logout(); // Gunakan method logout dari AuthService
    console.log('Logged out successfully');
    this.router.navigate(['/login']); // Redirect ke halaman login
  }
}
