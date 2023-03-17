import ReproIssue15475Host from '../../../../libs/repro-issue-15475-host/src/lib/repro-issue-15475-host';
import { Route, Routes } from 'react-router-dom';
import { lazy, Suspense } from 'react';
const RemoteApp = lazy(() => import('repro-issue-15475-remote/Module'));

export function App() {
  return (
  <Suspense fallback={null}>
    <Routes>
      <Route path="/" element={<ReproIssue15475Host/>}/>
      <Route path="/remote-app" element={<RemoteApp/>}/>
    </Routes>
  </Suspense>
  );
}

export default App;
