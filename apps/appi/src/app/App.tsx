import { Routes, Route } from 'react-router-dom';
import { AppiComponents, Repos, AppMainLayout } from '@g_api/appi-components';

export function App() {
  return (
    <Routes>
        <Route path="/" element={<AppMainLayout />}>
          <Route path="repos" element={<Repos />} />
          <Route path="appi" element={<AppiComponents />} />
        </Route>
    </Routes>
  );
}

export default App;
