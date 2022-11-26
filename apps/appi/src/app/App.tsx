import HostLayout from './RootAppPage';
import { Routes, Route } from 'react-router-dom';
import Tasks from './pages/Tasks';


export function App() {
  return (
    <Routes>
      <Route path="/" element={<HostLayout />}>
        <Route path="tasks" element={<Tasks />} />
      </Route>
    </Routes>
  );
}

export default App;
