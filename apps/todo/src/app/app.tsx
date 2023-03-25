import styled from '@emotion/styled';
import { TodoList } from '@g_api/todo-lib';

const StyledApp = styled.div`
  // Your style here
`;

export function App() {
  return (
    <StyledApp>
      <TodoList/>
    </StyledApp>
  );
}

export default App;
