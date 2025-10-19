// src/App.tsx
import { QueryProvider } from './app/providers/QueryProvider';
import { AppRoutes } from './app-routes';

export default function App() {
  return (
    <QueryProvider>
      <AppRoutes />
    </QueryProvider>
  );
}