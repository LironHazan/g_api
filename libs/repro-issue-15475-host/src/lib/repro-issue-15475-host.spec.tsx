import { render } from '@testing-library/react';

import ReproIssue15475Host from './repro-issue-15475-host';

describe('ReproIssue15475Host', () => {
  it('should render successfully', () => {
    const { baseElement } = render(<ReproIssue15475Host />);
    expect(baseElement).toBeTruthy();
  });
});
