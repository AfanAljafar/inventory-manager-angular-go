import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ApiService {
  private apiUrl = 'http://localhost:4001'; // Ganti dengan URL API backend kamu

  constructor(private http: HttpClient) {}

  // GET all items
  getAllItems(): Observable<any> {
    return this.http.get(`${this.apiUrl}/items`);
  }

  // POST a new item
  addItem(data: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/items`, data);
  }

  //find item by category
  filterItemByCategory(category: string): Observable<any> {
    const params = new HttpParams().set('category', category);
    return this.http.get(`${this.apiUrl}/items/filter`, { params });
  }

  //DELETE item by ID (optional)
  deleteItem(id: number): Observable<any> {
    return this.http.delete(`${this.apiUrl}/items/${id}`);
  }

  // PATCH update item by ID (optional)
  updateItem(id: number, data: any): Observable<any> {
    return this.http.patch(`${this.apiUrl}/items/${id}`, data);
  }

  // PUT update item (optional)
  // updateItem(id: number, data: any): Observable<any> {
  //   return this.http.put(`${this.apiUrl}/items/${id}`, data);
  // }
}
