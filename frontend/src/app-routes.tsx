// src/app-routes.tsx
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import HomePage from './pages/home/index';

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomePage />,
  },
]);

export function AppRoutes() {
  return <RouterProvider router={router} />;
}