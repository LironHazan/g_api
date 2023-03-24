import styled from '@emotion/styled';

import NxWelcome from './nx-welcome';
import TodoLib from '../../../../libs/todo-lib/src/lib/todo-lib';

const StyledApp = styled.div`
  // Your style here
`;

export function App() {
  return (
    <StyledApp>
      <TodoLib/>
      {/*<NxWelcome title="todo" />*/}
    </StyledApp>
  );
}

export default App;
