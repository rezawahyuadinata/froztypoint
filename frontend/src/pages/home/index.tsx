// src/pages/Home/index.tsx
import { useProducts } from '../../features/product/hooks/useProduct';

export default function HomePage() {
    // mengambil data constant pada use product yaitu product dengan cara destrukturisasi
    // karena pada useProduct mengembalikan object dengan properti data, isLoading, dan error
//   const { data: products, isLoading, error } = useProducts();
  const { data: products } = useProducts();

//   if (isLoading) return <div className="p-4">Loading...</div>;
//   if (error) return <div className="p-4 text-red-500">Failed to load products</div>;

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Welcome to Froztypoint</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        {products?.map((product) => (
          <div key={product.id} className="border p-4 rounded shadow">
            <h2 className="font-semibold">{product.name}</h2>
            <p className="text-gray-600">Rp {product.price.toLocaleString()}</p>
            <button className="mt-2 bg-blue-500 text-white px-3 py-1 rounded">
              Add to Cart
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}