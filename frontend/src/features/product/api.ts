// src/features/products/api.ts
import { type Product } from './types';

// Ganti dengan URL backend Go-mu nanti
const API_BASE = 'http://localhost:8080';

export const fetchProducts = async (): Promise<Product[]> => {
  const res = await fetch(`${API_BASE}/products`);
  if (!res.ok) throw new Error('Failed to fetch products');
  return res.json();
};