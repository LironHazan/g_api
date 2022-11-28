import { Routes, Route } from 'react-router-dom';
import { AppiComponents, Tasks, AppMainLayout } from '@react-stack-2022/appi-components';


export function App() {
  return (
    <Routes>
      <Route path="/" element={<AppMainLayout />}>
        <Route path="tasks" element={<Tasks />} />
        <Route path="appi" element={<AppiComponents />} />
      </Route>
    </Routes>
  );
}

export default App;
