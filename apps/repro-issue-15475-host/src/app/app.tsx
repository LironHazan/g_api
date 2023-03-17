import ReproIssue15475Host from '../../../../libs/repro-issue-15475-host/src/lib/repro-issue-15475-host';
import { Route, Routes } from 'react-router-dom';


export function App() {
  return (
      <Routes>
        <Route path="/" element={<ReproIssue15475Host/>}>
        </Route>
      </Routes>
  );
}

export default App;
