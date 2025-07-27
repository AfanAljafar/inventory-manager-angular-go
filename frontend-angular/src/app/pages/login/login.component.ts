import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { AuthService } from '../../services/auth.service';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, CommonModule],
  templateUrl: './login.component.html',
})
export class LoginComponent {
  constructor(
    private http: HttpClient,
    private router: Router,
    private authService: AuthService
  ) {}

  username = '';
  password = '';
  errorMessage = '';

  onSubmit() {
    if (this.username === '' || this.password === '') {
      this.errorMessage = 'Username and password are required.';
      console.log('Username and password are required.');
      return;
    }

    this.http
      .post<any>('http://localhost:4001/users/login', {
        email: this.username,
        password: this.password,
      })
      .subscribe({
        next: (response) => {
          console.log('Login success:', response);
          console.log(
            'userId di localStorage:',
            localStorage.getItem('userId')
          );

          // ✅ Gunakan AuthService untuk simpan token dan userId
          this.authService.loginSuccess(response.token, response.user_id);
          // Simpan token JWT ke localStorage
          // localStorage.setItem('token', response.token);
          localStorage.setItem('role', response.role); // ⬅️ "admin" atau "regular"
          this.errorMessage = '';
          this.router.navigate(['/dashboard']);
        },
        error: (err) => {
          console.error('Login failed:', err);
          this.errorMessage = err?.error?.error || 'Login failed';
        },
      });

    // if (this.username === 'admin' && this.password === 'admin') {
    //   this.errorMessage = '';

    //   console.log('Login Successful');
    //   this.router.navigate(['/dashboard']);
    // } else {
    //   this.errorMessage = 'Invalid username or password.';
    // }
  }

  ngOnInit() {
    const token = localStorage.getItem('token');
    if (token) {
      this.router.navigate(['/dashboard']);
    }
  }
}
