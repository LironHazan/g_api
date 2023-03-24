import { render } from '@testing-library/react';

import TodoLib from './todo-lib';

describe('TodoLib', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<TodoLib />);
    expect(baseElement).toBeTruthy();
  });
});
