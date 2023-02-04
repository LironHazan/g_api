import { Routes, Route } from 'react-router-dom';
import { AppiPage, AppMainLayout, ReposPage } from '@g_api/appi-components';

export function App() {
  return (
    <Routes>
        <Route path="/" element={<AppMainLayout />}>
          <Route path="repos" element={<ReposPage/>} />
          <Route path="appi" element={<AppiPage />} />
        </Route>
    </Routes>
  );
}

export default App;
