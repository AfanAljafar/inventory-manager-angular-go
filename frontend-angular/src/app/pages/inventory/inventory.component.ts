import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ApiService } from '../../services/api.service'; // Import ApiService
import { AuthService } from '../../services/auth.service'; // Import AuthService

@Component({
  selector: 'app-inventory',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './inventory.component.html',
  styleUrl: './inventory.component.css',
})
export class InventoryComponent implements OnInit {
  constructor(
    private http: HttpClient,
    private api: ApiService,
    private authService: AuthService
  ) {}

  ngOnInit(): void {
    this.loadItems();

    // ✅ Ambil userId dari AuthService
    const userId = this.authService.getUserId();
    if (userId !== null) {
      this.newItem.goods_in_by = userId;
    } else {
      alert('Silakan login terlebih dahulu.');
    }
  }

  items: any[] = [];
  filteredItems: any[] = [];
  selectedCategory: string = '';
  searchTerm: string = '';
  editIndex: number = -1;
  newItem = {
    category_item: '',
    item_name: '',
    item_picture: '',
    quantity: 0,
    unit: '',
    price: 0,
    goods_in_by: null as number | null, // Pastikan userId diisi (bisa ambil dari auth service nanti)
  };
  isAddItemModalOpen = false;

  // pagination
  currentPage: number = 1;
  itemsPerPage: number = 10;

  get paginatedItems(): any[] {
    const startIndex = (this.currentPage - 1) * this.itemsPerPage;
    const endIndex = startIndex + this.itemsPerPage;
    return this.filteredItems.slice(startIndex, endIndex);
  }

  get totalPages(): number {
    return Math.ceil(this.filteredItems.length / this.itemsPerPage);
  }

  loadItems() {
    this.api.getAllItems().subscribe({
      next: (res) => {
        this.items = res.data;
        this.filterItems();
      },
      error: (err) => {
        console.error('Error fetching items:', err);
      },
    });
  }

  openAddItemModal() {
    this.isAddItemModalOpen = true;
    this.resetNewItem();
    this.editIndex = -1;
  }
  openEditItemModal(item: any, itemId: number) {
    this.newItem = { ...item };
    this.editIndex = itemId; // Simpan ID item yang akan diupdate
    this.isAddItemModalOpen = true;
  }
  closeEditItemModal() {
    this.isAddItemModalOpen = false;
    this.editIndex = -1;
    this.resetNewItem();
  }

  closeAddItemModal() {
    this.isAddItemModalOpen = false;
    this.resetNewItem();
    this.editIndex = -1;
  }

  addItem() {
    const userId = this.authService.getUserId();

    if (userId === null) {
      alert('Anda harus login terlebih dahulu untuk menambahkan item.');
      return;
    }

    this.newItem.goods_in_by = userId;

    if (
      !this.newItem.category_item ||
      !this.newItem.item_name ||
      !this.newItem.item_picture ||
      this.newItem.quantity <= 0 ||
      !this.newItem.unit ||
      this.newItem.price <= 0
    ) {
      alert('Semua field harus diisi dengan benar.');
      return;
    }

    if (this.editIndex !== -1) {
      // Update existing item
      this.api.updateItem(this.editIndex, this.newItem).subscribe({
        next: () => {
          this.closeEditItemModal();
          this.loadItems(); // refresh item list
          this.resetNewItem(); // reset input
        },
        error: (err) => {
          console.error('Error updating item:', err);
          alert('Gagal update item');
        },
      });
      return;
    } else {
      this.api.addItem(this.newItem).subscribe({
        next: () => {
          this.closeAddItemModal();
          this.loadItems(); // refresh item list
          this.resetNewItem(); // reset input
        },
        error: (err) => {
          console.error('Error adding item:', err);
          alert('Gagal tambah item');
        },
      });
    }
  }

  resetNewItem() {
    this.newItem = {
      category_item: '',
      item_name: '',
      item_picture: '',
      quantity: 0,
      unit: '',
      price: 0,
      goods_in_by: this.authService.getUserId() || null,
    };
  }

  filterItems(): void {
    if (this.selectedCategory) {
      this.api.filterItemByCategory(this.selectedCategory).subscribe({
        next: (res) => {
          // Opsional: Apply searchTerm di sisi frontend
          const filtered = res.data;

          this.filteredItems = filtered.filter((item: any) =>
            item.item_name.toLowerCase().includes(this.searchTerm.toLowerCase())
          );
        },
        error: (err) => {
          console.error('Error filtering items:', err);
          this.filteredItems = []; // kosongkan jika error
        },
      });
    } else {
      // fallback: tampilkan semua jika tidak pilih kategori
      this.filteredItems = this.items.filter((item) =>
        item.item_name.toLowerCase().includes(this.searchTerm.toLowerCase())
      );
    }
  }

  deleteItem(itemId: number) {
    const userId = this.authService.getUserId();

    if (userId === null) {
      alert('Anda harus login terlebih dahulu untuk menambahkan item.');
      return;
    }

    this.api.deleteItem(itemId).subscribe({
      next: () => {
        console.log('Item deleted successfully');
        this.loadItems(); // refresh item list
      },
      error: (err) => {
        console.error('Error deleting item:', err);
        alert('Gagal menghapus item');
      },
    });
    // Hapus item dari daftar lokal
    this.items = this.items.filter((item) => item.id !== itemId);
    this.filterItems();
    this.loadItems(); // refresh item list
  }

  exportToCSV(): void {
    const header = [
      'ID',
      'Name',
      'Category',
      'Picture',
      'QTY',
      'Unit',
      'Price',
      'In By',
      'Out By',
      'Time In',
      'Time Out',
    ];
    const rows = this.items.map((item) => [
      item.id,
      item.item_name,
      item.category_item,
      item.item_picture,
      item.quantity,
      item.unit,
      item.price,
      item.goods_in_by,
      item.goods_out_by || '', // Jika tidak ada, bisa kosongkan
      item.goods_in_by ? new Date(item.goods_in_by).toLocaleString() : '',
      item.goods_out_by ? new Date(item.goods_out_by).toLocaleString() : '',
    ]);

    const csvContent = [header, ...rows]
      .map((e) =>
        e.map((field) => `"${String(field).replace(/"/g, '""')}"`).join(',')
      )
      .join('\n');

    const blob = new Blob([csvContent], { type: 'text/csv;charset=utf-8;' });
    const link = document.createElement('a');
    const url = URL.createObjectURL(blob);

    link.setAttribute('href', url);
    link.setAttribute('download', 'inventory.csv');
    link.style.visibility = 'hidden';

    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }

  prevPage(): void {
    if (this.currentPage > 1) {
      this.currentPage--;
    }
  }

  nextPage(): void {
    if (this.currentPage < this.totalPages) {
      this.currentPage++;
    }
  }
}
